package eventsub

import (
	"time"

	"github.com/kvizyx/twitchy/eventsub/eventtracker"
	"github.com/kvizyx/twitchy/internal/json"
)

// EventSub is a module that is responsible for Twitch's eventsub.
//
// Reference: https://dev.twitch.tv/docs/eventsub.
type EventSub struct {
	eventTracker eventtracker.EventTracker
	marshal      json.Marshaller
}

func New(options ...Option) EventSub {
	var es EventSub

	for _, option := range options {
		option(&es)
	}

	if es.marshal != nil {
		json.Marshal = es.marshal
	}

	return es
}

// Webhook returns new Webhook HTTP handler that implements http.Handler.
func (es *EventSub) Webhook(secret []byte, verifySignature bool) (*Webhook, error) {
	return newWebhook(secret, es.eventTracker, verifySignature)
}

// Websocket returns new Websocket client.
func (es *EventSub) Websocket(options ...WebsocketOption) *Websocket {
	return newWebsocket(es.eventTracker, options...)
}

// isExpiredMessage returns does message with provided timestamp is too old (expired) to process or not.
//
// According to the Twitch's documentation we should not process messages that are older than 10 minutes from the moment
// they were create to guard against replay attacks.
func isExpiredMessage(messageTimestamp TimestampUTC) bool {
	const expirationSeconds = 600

	now, _ := timestampUTCFromString(time.Now().String())

	if now.Sub(messageTimestamp.Time).Seconds() > expirationSeconds {
		return true
	}

	return false
}
