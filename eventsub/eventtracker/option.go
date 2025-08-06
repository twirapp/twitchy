package eventtracker

import (
	"time"
)

type Option func(*option)

// option is an options for all internal EventTracker implementations.
type option struct {
	eventTTL time.Duration
}

func WithEventTTL(ttl time.Duration) Option {
	return func(o *option) {
		o.eventTTL = ttl
	}
}
