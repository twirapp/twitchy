package eventtracker

import (
	"context"
	"time"
)

// EventTTL is a desired (according to the Twitch's documentation) time-to-live for events in storage to count them as deduplicated.
// In most cases this is what you want use as TTL for events in your custom EventTracker implementation (if it's not a test
// or something like that).
//
// This TTL is used by default in default EventTracker implementations.
const EventTTL = 10 * time.Minute

type EventID = string

// EventTracker tracks events from EventSub to avoid handling the same (duplicate) events multiple times.
type EventTracker interface {
	// Track starts tracking an event with the provided ID and returns if that event is already
	// being tracked (is it duplicate or not).
	Track(ctx context.Context, eventID EventID) (bool, error)
}
