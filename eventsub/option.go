package eventsub

import (
	"github.com/kvizyx/twitchy/eventsub/eventtracker"
)

type Option func(*EventSub)

func WithEventTracker(eventTracker eventtracker.EventTracker) Option {
	return func(eventSub *EventSub) {
		eventSub.eventTracker = eventTracker
	}
}
