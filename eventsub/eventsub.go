package eventsub

import (
	"time"
)

type MessageID = string

type EventSub struct {
	eventTracker     EventTracker
	eventTrackingTTL time.Duration
}

func New(options ...Option) EventSub {
	eventSub := EventSub{
		// TODO: respect the event tracking TTL from options.
		eventTracker: NewInMemoryEventTracker(EventTrackingTTL),
	}

	for _, option := range options {
		option(&eventSub)
	}

	return eventSub
}

func (es *EventSub) Webhook(secret []byte, verifySignature bool) (*Webhook, error) {
	return newWebhook(secret, es.eventTracker, verifySignature)
}

func (es *EventSub) Websocket(options ...WebsocketOption) *Websocket {
	return newWebsocket(es.eventTracker, options...)
}
