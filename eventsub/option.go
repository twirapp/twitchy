package eventsub

import (
	"github.com/kvizyx/twitchy/eventsub/eventtracker"
	"github.com/kvizyx/twitchy/internal/json"
)

// Option is an optional setting for EventSub.
type Option func(*EventSub)

// WithEventTracker sets tracker for events to prevent duplicate events.
// If this option is not set, events will not be tracked as there is no default value.
func WithEventTracker(eventTracker eventtracker.EventTracker) Option {
	return func(es *EventSub) {
		es.eventTracker = eventTracker
	}
}

// WithMarshal sets JSON marshaller that will be used as default marshaller.
//
// Default value is a standard library marshaller.
func WithMarshal(marshal json.MarshalFn) Option {
	return func(es *EventSub) {
		es.marshal = marshal
	}
}
