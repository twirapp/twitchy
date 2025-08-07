package eventsub

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/coder/websocket"
	"github.com/kvizyx/twitchy/eventsub/eventtracker"
	"github.com/kvizyx/twitchy/internal/json"
)

var (
	ErrConnectionUnused = errors.New("connection is unused as subscription time limit exceeded")
	ErrReconnectTimeout = errors.New("reconnect grace time expired")
)

const websocketURL = "wss://eventsub.wss.twitch.tv/ws"

type Websocket struct {
	client       *http.Client
	eventTracker eventtracker.EventTracker

	serverURL          string
	serverReconnectURL string

	retryAttempts uint
	retryDelay    time.Duration

	conn            *websocket.Conn
	connDialOptions *websocket.DialOptions

	keepaliveSeconds uint
	lastKeepalive    atomic.Value

	isActive       atomic.Bool
	isReconnecting atomic.Bool

	restart     chan struct{}
	welcome     chan struct{}
	reconnected chan struct{}
	disconnect  chan struct{}

	onWelcome        func(WebsocketWelcomeMessage)
	onKeepalive      func(WebsocketKeepaliveMessage)
	onPing           func()
	onReconnect      func(WebsocketReconnectMessage)
	onReconnectError func(error)
	onDisconnect     func()

	callback[WebsocketNotificationMetadata]
}

func newWebsocket(eventTracker eventtracker.EventTracker, options ...WebsocketOption) *Websocket {
	ws := &Websocket{
		client:             http.DefaultClient,
		eventTracker:       eventTracker,
		serverURL:          websocketURL,
		serverReconnectURL: websocketURL,
		retryAttempts:      5,
		retryDelay:         1 * time.Second,
		keepaliveSeconds:   600,
		reconnected:        make(chan struct{}),
		welcome:            make(chan struct{}),
		restart:            make(chan struct{}),
	}

	for _, option := range options {
		option(ws)
	}

	ws.serverURL = fmt.Sprintf("%s?keepalive_timeout_seconds=%d", ws.serverURL, ws.keepaliveSeconds)

	ws.connDialOptions = &websocket.DialOptions{
		HTTPClient: ws.client,
		OnPingReceived: func(_ context.Context, _ []byte) bool {
			if ws.onPing != nil {
				go ws.onPing()
			}

			return true
		},
	}

	return ws
}

func (ws *Websocket) Connect(ctx context.Context) error {
	if !ws.setActive() {
		return nil
	}

	// We should recreate disconnect channel on each connect so user can do connect-disconnect cycle as many times as he
	// wants and don't get panic on double-close already closed channel.
	ws.disconnect = make(chan struct{})

	defer func() {
		ws.setInactivate()

		if ws.onDisconnect != nil {
			go ws.onDisconnect()
		}
	}()

	if err := ws.connect(ctx, ws.serverURL); err != nil {
		return err
	}

	// Context of the whole instance including side workers (e.g. keepalive worker).
	lifecycleCtx, stopLifecycle := context.WithCancel(context.Background())
	defer stopLifecycle()

	// Context of the connection itself that can be canceled to restart for example.
	connectionCtx, stopConnection := context.WithCancel(lifecycleCtx)

	stop := make(chan error)

	go ws.startKeepaliveWorker(lifecycleCtx)
	go func() {
		stop <- ws.startReadWorker(connectionCtx)
	}()

	for {
		select {
		case err := <-stop:
			if ws.isReconnecting.Load() {
				continue
			}
			stopConnection()
			return err
		case <-ws.restart:
			stopConnection()

			// Restore connection context to reuse it on next restart (reconnect).
			connectionCtx, stopConnection = context.WithCancel(lifecycleCtx)

			go func() {
				stop <- ws.startReadWorker(connectionCtx)
			}()
		case <-ws.disconnect:
			stopLifecycle()
		}
	}
}

func (ws *Websocket) Disconnect() error {
	if !ws.setInactivate() {
		return nil
	}

	ws.disconnect <- struct{}{}
	close(ws.disconnect)

	if err := ws.conn.Close(websocket.StatusNormalClosure, "client is shutting down"); err != nil {
		return fmt.Errorf("close connection: %w", err)
	}

	return nil
}

func (ws *Websocket) connect(ctx context.Context, url string) error {
	if ws.retryAttempts > 0 {
		return retry.Do(
			func() error {
				return ws.connectWithoutRetry(ctx, url)
			},
			retry.Attempts(ws.retryAttempts),
			retry.Delay(ws.retryDelay),
			retry.Context(ctx),
		)
	}

	return ws.connectWithoutRetry(ctx, url)
}

func (ws *Websocket) connectWithoutRetry(ctx context.Context, url string) error {
	conn, _, err := websocket.Dial(ctx, url, ws.connDialOptions)
	if err != nil {
		return fmt.Errorf("connect to server: %w", err)
	}

	ws.conn = conn

	return nil
}

// reconnect reconnects to the new eventsub edge server according to the reconnect message flow.
//
// Reference: https://dev.twitch.tv/docs/eventsub/handling-websocket-events/#reconnect-message.
func (ws *Websocket) reconnect(ctx context.Context, reconnectURL string) error {
	if !ws.setReconnecting() {
		return nil
	}

	defer func() {
		ws.isReconnecting.Store(false)

		select {
		case ws.reconnected <- struct{}{}:
		default:
		}
	}()

	// Save old connection to close it later after connecting to the new edge server.
	oldConnection := ws.conn

	// Set new server reconnect URL.
	ws.serverReconnectURL = reconnectURL

	// Connect to the new eventsub edge server.
	if err := ws.connect(ctx, reconnectURL); err != nil {
		return fmt.Errorf("connect: %w", err)
	}

	// Signal to read worker to stop reading from old connection and restart with new one.
	ws.restart <- struct{}{}

	// Wait for a welcome message on new connection.
	<-ws.welcome

	// Close old connection to complete reconnect flow.
	if err := oldConnection.Close(websocket.StatusNormalClosure, "connected to the new edge server"); err != nil {
		return fmt.Errorf("close old connection: %w", err)
	}

	return nil
}

// startReadWorker starts and blocks on reading and processing messages from the connection.
func (ws *Websocket) startReadWorker(ctx context.Context) error {
	for {
		var message websocketRawMessage

		_, payload, err := ws.conn.Read(ctx)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				return nil
			}

			switch websocket.CloseStatus(err) {
			case websocket.StatusNormalClosure:
				if ws.isReconnecting.Load() {
					<-ws.reconnected
					continue
				}

				return nil
			case websocket.StatusCode(4003):
				return ErrConnectionUnused
			case websocket.StatusCode(4004):
				return ErrReconnectTimeout
			}

			return fmt.Errorf("read from connection: %w", err)
		}

		if err = json.Unmarshal(payload, &message); err != nil {
			return fmt.Errorf("unmarshal message: %w", err)
		}

		if err = ws.handleMessage(ctx, message); err != nil {
			return fmt.Errorf("handle message: %w", err)
		}
	}
}

// startKeepaliveWorker starts and blocks on trying to keep the connection healthy and reconnect if needed.
func (ws *Websocket) startKeepaliveWorker(ctx context.Context) {
	const delay = 1 * time.Second

	keepalive := time.NewTicker(time.Second*time.Duration(ws.keepaliveSeconds) + delay)
	defer keepalive.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-keepalive.C:
			lastKeepalive, _ := timestampUTCFromString(time.Now().String())

			// We must reconnect to the eventsub server if keepalive timeout expired.
			if lastKeepalive.Sub(ws.getKeepalive()).Seconds() > float64(ws.keepaliveSeconds) {
				if err := ws.reconnect(ctx, ws.serverReconnectURL); err != nil {
					if ws.onReconnectError != nil {
						go ws.onReconnectError(err)
					}
				}
			}
		}
	}
}

// setReconnecting sets current Websocket instance to reconnecting state meaning that client is reconnecting now.
func (ws *Websocket) setReconnecting() bool {
	return ws.isReconnecting.CompareAndSwap(false, true)
}

// setActive sets current Websocket instance to active state meaning that it's running.
func (ws *Websocket) setActive() bool {
	return ws.isActive.CompareAndSwap(false, true)
}

// setInactivate sets current Websocket instance to inactive state meaning that it's not running anymore.
func (ws *Websocket) setInactivate() bool {
	return ws.isActive.CompareAndSwap(true, false)
}

// setKeepalive sets last keepalive timestamp atomically.
func (ws *Websocket) setKeepalive(timestamp TimestampUTC) {
	ws.lastKeepalive.Store(timestamp)
}

// getKeepalive loads last keepalive timestamp atomically and returns its value.
func (ws *Websocket) getKeepalive() time.Time {
	lastKeepalive := ws.lastKeepalive.Load().(TimestampUTC)
	return lastKeepalive.Time
}
