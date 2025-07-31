package eventsub

import (
	"context"
	"time"
)

const EventTrackingTTL = 10*time.Minute + 5*time.Second

// EventTracker tracks events from EventSub to avoid handling the same (duplicate) events multiple times.
type EventTracker interface {
	// Track starts tracking an event with the provided ID and returns if that event is already
	// being tracked (is it duplicate or not).
	Track(ctx context.Context, eventID string) (bool, error)
}
