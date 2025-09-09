package eventtracker

import (
	"time"
)

// Option is an option for standard EventTracker implementation.
type Option func(*option)

type option struct {
	eventTTL time.Duration
}

// WithEventTTL sets TTL for events to be tracked. Usually, you don't want to you use this option.
//
// Default value is SafeEventTTL.
func WithEventTTL(ttl time.Duration) Option {
	return func(o *option) {
		o.eventTTL = ttl
	}
}
