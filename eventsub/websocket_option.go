package eventsub

import (
	"net/http"
	"time"
)

// WebsocketOption is an optional setting for Websocket.
type WebsocketOption func(*Websocket)

func WebsocketWithClient(client *http.Client) WebsocketOption {
	return func(ws *Websocket) {
		ws.client = client
	}
}

func WebsocketWithServerURL(serverURL string) WebsocketOption {
	return func(ws *Websocket) {
		ws.serverURL = serverURL
	}
}

func WebsocketWithKeepalive(seconds uint) WebsocketOption {
	return func(ws *Websocket) {
		ws.keepaliveSeconds = seconds
	}
}

func WebsocketWithNoRetry() WebsocketOption {
	return func(ws *Websocket) {
		ws.retryAttempts = 0
	}
}

func WebsocketWithRetryAttempts(attempts uint) WebsocketOption {
	return func(ws *Websocket) {
		ws.retryAttempts = attempts
	}
}

func WebsocketWithRetryDelay(delay time.Duration) WebsocketOption {
	return func(ws *Websocket) {
		ws.retryDelay = delay
	}
}
