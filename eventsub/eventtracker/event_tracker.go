package eventtracker

import (
	"context"
	"time"
)

// SafeEventTTL is a safe (according to the Twitch's documentation) time-to-live for events in storage to count them as duplicated.
// In most cases this is what you want use as TTL for events in your custom EventTracker implementation (if it's not a test
// or something like that). It's used by default in standard EventTracker implementations.
const SafeEventTTL = 10 * time.Minute

// EventTracker tracks events sent from eventsub server to avoid handling the same (duplicate) events multiple times.
type EventTracker interface {
	// Track starts tracking of event with the provided identifier and returns if that event is already being tracked (duplicate or not).
	Track(ctx context.Context, eventID string) (bool, error)
}
