package eventsub

import (
	"context"
	"time"
)

// EventTTL is a desired (according to the Twitch's documentation) time-to-live for events in storage to deduplicate them.
// In most cases this is what you want use as TTL for events in your custom EventTracker implementation (if it's not a test
// or something like that).
const EventTTL = 10 * time.Minute

// EventTracker tracks events from EventSub to avoid handling the same (duplicate) events multiple times.
type EventTracker interface {
	// Track starts tracking an event with the provided ID and returns if that event is already
	// being tracked (is it duplicate or not).
	Track(ctx context.Context, eventID string) (bool, error)
}
