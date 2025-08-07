package eventsub

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/kvizyx/twitchy/internal/json"
)

type webhookRawNotification struct {
	Subscription json.RawMessage `json:"subscription"`
	Event        json.RawMessage `json:"event"`
}

func (wh *Webhook) handleNotification(w http.ResponseWriter, header http.Header, body []byte) {
	metadata, err := wh.extractNotificationMetadata(header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var notification webhookRawNotification

	if err = json.Unmarshal(body, &notification); err != nil {
		http.Error(w, "failed to unmarshal notification", http.StatusInternalServerError)
		return
	}

	event := RawEvent{
		Subscription: notification.Subscription,
		Event:        notification.Event,
	}

	if err = wh.callback.runEventCallback(metadata.SubscriptionType, metadata.SubscriptionVersion, event, metadata); err != nil {
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

// extractNotificationMetadata extracts Twitch's webhook request headers and returns them as WebhookNotificationMetadata.
func (wh *Webhook) extractNotificationMetadata(header http.Header) (WebhookNotificationMetadata, error) {
	var (
		rawMessageRetry     = header.Get("Twitch-Eventsub-Message-Retry")
		rawMessageTimestamp = header.Get("Twitch-Eventsub-Message-Timestamp")
	)

	messageRetry, err := strconv.Atoi(rawMessageRetry)
	if err != nil {
		return WebhookNotificationMetadata{}, fmt.Errorf("parse retry header: %w", err)
	}

	messageTimestamp, err := timestampUTCFromString(rawMessageTimestamp)
	if err != nil {
		return WebhookNotificationMetadata{}, fmt.Errorf("parse message timestamp header: %w", err)
	}

	subscriptionType := header.Get("Twitch-Eventsub-Subscription-Type")

	return WebhookNotificationMetadata{
		MessageID:           header.Get("Twitch-Eventsub-Message-Id"),
		MessageRetry:        messageRetry,
		MessageType:         header.Get("Twitch-Eventsub-Message-Type"),
		MessageSignature:    header.Get("Twitch-Eventsub-Message-Signature"),
		MessageTimestamp:    messageTimestamp,
		SubscriptionType:    EventType(subscriptionType),
		SubscriptionVersion: header.Get("Twitch-Eventsub-Subscription-Version"),
	}, nil
}
