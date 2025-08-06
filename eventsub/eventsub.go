package eventsub

import (
	"github.com/kvizyx/twitchy/eventsub/eventtracker"
)

type EventSub struct {
	eventTracker eventtracker.EventTracker
}

func New(options ...Option) EventSub {
	var es EventSub

	for _, option := range options {
		option(&es)
	}

	return es
}

// Webhook returns new Webhook instance with shared EventSub.
func (es *EventSub) Webhook(secret []byte, verifySignature bool) (*Webhook, error) {
	return newWebhook(secret, es.eventTracker, verifySignature)
}

// Websocket returns new Websocket instance with shared EventSub.
func (es *EventSub) Websocket(options ...WebsocketOption) *Websocket {
	return newWebsocket(es.eventTracker, options...)
}
