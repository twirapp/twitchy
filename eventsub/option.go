package eventsub

import (
	"time"
)

type Option func(*EventSub)

func WithEventTracker(eventTracker EventTracker) Option {
	return func(eventSub *EventSub) {
		eventSub.eventTracker = eventTracker
	}
}

func WithEventTrackingTTL(ttl time.Duration) Option {
	return func(eventSub *EventSub) {
		eventSub.eventTrackingTTL = ttl
	}
}
