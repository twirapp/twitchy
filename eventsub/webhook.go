package eventsub

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"net/http"

	"github.com/kvizyx/twitchy/eventsub/eventtracker"
	"github.com/kvizyx/twitchy/internal/json"
)

var ErrInvalidWebhookSecret = errors.New("secret must be a minimum of 10 and maximum of 100 characters long")

type Webhook struct {
	eventTracker              eventtracker.EventTracker
	secret                    []byte
	withSignatureVerification bool

	onRevocation   func(WebhookRevocationNotification)
	onVerification func(WebhookCallbackVerificationNotification)

	callback[WebhookNotificationMetadata]
}

func newWebhook(secret []byte, eventTracker eventtracker.EventTracker, verifySignature bool) (*Webhook, error) {
	if len(secret) < 10 || len(secret) > 100 {
		return nil, ErrInvalidWebhookSecret
	}

	return &Webhook{
		eventTracker:              eventTracker,
		secret:                    secret,
		withSignatureVerification: verifySignature,
	}, nil
}

func (wh *Webhook) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "cannot read body", http.StatusInternalServerError)
		return
	}
	defer func() {
		_ = r.Body.Close()
	}()

	messageId := r.Header.Get("Twitch-Eventsub-Message-Id")

	if wh.withSignatureVerification {
		var (
			messageSignature = r.Header.Get("Twitch-Eventsub-Message-Signature")
			messageTimestamp = r.Header.Get("Twitch-Eventsub-Message-Timestamp")
		)

		if len(messageSignature) == 0 {
			http.Error(w, "missing signature header", http.StatusBadRequest)
			return
		}

		if !wh.isValidSignature(messageSignature, body, messageId, messageTimestamp) {
			http.Error(w, "invalid signature", http.StatusBadRequest)
			return
		}
	}

	if wh.eventTracker != nil {
		isDuplicate, err := wh.eventTracker.Track(r.Context(), messageId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if isDuplicate {
			if wh.onDuplicate != nil {
				go wh.onDuplicate(messageId)
			}
			return
		}
	}

	messageType := r.Header.Get("Twitch-Eventsub-Message-Type")

	switch messageType {
	case "notification":
		wh.handleNotification(w, r.Header, body)
	case "webhook_callback_verification":
		wh.handleCallbackVerification(w, body)
	case "revocation":
		wh.handleRevocation(w, body)
	default:
		http.Error(w, "undefined message type", http.StatusBadRequest)
		return
	}
}

func (wh *Webhook) handleRevocation(w http.ResponseWriter, body []byte) {
	if _, ok := runWebhookHandler(wh.onRevocation, w, body); !ok {
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (wh *Webhook) handleCallbackVerification(w http.ResponseWriter, body []byte) {
	notification, ok := runWebhookHandler(wh.onVerification, w, body)
	if !ok {
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte(notification.Challenge))
}

func (wh *Webhook) isValidSignature(signature string, body []byte, messageId, messageTimestamp string) bool {
	mac := hmac.New(sha256.New, wh.secret)

	mac.Write([]byte(messageId))
	mac.Write([]byte(messageTimestamp))
	mac.Write(body)

	computedSignature := "sha256=" + hex.EncodeToString(mac.Sum(nil))

	return hmac.Equal([]byte(computedSignature), []byte(signature))
}

// runWebhookHandler parses payload from JSON body and runs callback in a separate go-routine with parsed payload if it's
// not nil and returns payload itself. Returns true on success.
func runWebhookHandler[Payload any](callback func(Payload), w http.ResponseWriter, body []byte) (Payload, bool) {
	var payload Payload

	if callback != nil {
		if err := json.Unmarshal(body, &payload); err != nil {
			http.Error(w, "failed to unmarshal payload", http.StatusInternalServerError)
			return payload, false
		}

		go callback(payload)
	}

	return payload, true
}
