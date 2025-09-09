package eventsub

import (
	"errors"
	"net/http"

	"github.com/kvizyx/twitchy/internal/json"
)

// webhookRawEvent is a webhook duplicate for RawEvent to close internal JSON wrapper from user.
type webhookRawEvent struct {
	Subscription json.RawMessage `json:"subscription"`
	Event        json.RawMessage `json:"event"`
}

// handleNotification handles webhook notification sent by eventsub server.
func (wh *Webhook) handleNotification(w http.ResponseWriter, metadata WebhookNotificationMetadata, body []byte) {
	messageType := metadata.MessageType

	switch messageType {
	case "notification":
		wh.handleEventNotification(w, metadata, body)
	case "webhook_callback_verification":
		wh.handleCallbackVerificationNotification(w, body)
	case "revocation":
		wh.handleRevocationNotification(w, body)
	default:
		http.Error(w, "undefined message type", http.StatusBadRequest)
	}
}

// handleEventNotification handles event notification request.
//
// Reference: https://dev.twitch.tv/docs/eventsub/handling-webhook-events/#processing-an-event.
func (wh *Webhook) handleEventNotification(w http.ResponseWriter, metadata WebhookNotificationMetadata, body []byte) {
	var rawNotification webhookRawEvent

	if err := json.Unmarshal(body, &rawNotification); err != nil {
		http.Error(w, "failed to unmarshal raw notification", http.StatusInternalServerError)
		return
	}

	rawEvent := RawEvent{
		Subscription: rawNotification.Subscription,
		Event:        rawNotification.Event,
	}

	if err := wh.callback.runEventCallback(metadata.SubscriptionType, metadata.SubscriptionVersion, rawEvent, metadata); err != nil {
		var status int

		if errors.Is(err, ErrUndefinedEventType) {
			status = http.StatusBadRequest
		} else {
			status = http.StatusInternalServerError
		}

		http.Error(w, err.Error(), status)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// handleRevocation handles event subscription revoke notification request.
//
// Reference: https://dev.twitch.tv/docs/eventsub/handling-webhook-events/#revoking-your-subscription.
func (wh *Webhook) handleRevocationNotification(w http.ResponseWriter, body []byte) {
	if _, ok := runEventWebhookHandler(wh.onRevocation, w, body); !ok {
		return
	}

	w.WriteHeader(http.StatusOK)
}

// handleCallbackVerification handles webhook challenge verification notification request.
//
// Reference: https://dev.twitch.tv/docs/eventsub/handling-webhook-events/#responding-to-a-challenge-request.
func (wh *Webhook) handleCallbackVerificationNotification(w http.ResponseWriter, body []byte) {
	notification, ok := runEventWebhookHandler(wh.onVerification, w, body)
	if !ok {
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte(notification.Challenge))
}

// runEventWebhookHandler parses provided request body payload as JSON data to generic payload and runs handler in separate go-routine
// with this payload if handler is defined and returns parsed payload with false, otherwise returns empty payload with true.
func runEventWebhookHandler[Payload any](
	handler func(Payload),
	w http.ResponseWriter,
	bodyPayload []byte,
) (Payload, bool) {
	var payload Payload

	if handler != nil {
		if err := json.Unmarshal(bodyPayload, &payload); err != nil {
			http.Error(w, "failed to unmarshal body payload", http.StatusInternalServerError)
			return payload, false
		}

		go handler(payload)
	}

	return payload, true
}
