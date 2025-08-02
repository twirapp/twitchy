package eventsub

import (
	"context"
	"time"

	"github.com/kvizyx/twitchy/internal/concurrentmap"
)

// InMemoryEventTracker is a concurrent-safe in-memory implementation of the EventTracker which is used by default if no
// other implementation is set by WithEventTracker option.
type InMemoryEventTracker struct {
	events concurrentmap.ConcurrentMap[string, struct{}]
}

var _ EventTracker = (*InMemoryEventTracker)(nil)

func NewInMemoryEventTracker(eventTrackingTTL time.Duration) *InMemoryEventTracker {
	return &InMemoryEventTracker{
		events: concurrentmap.NewString[struct{}](eventTrackingTTL),
	}
}

func (iet *InMemoryEventTracker) Track(_ context.Context, eventID string) (bool, error) {
	_, isAbsent := iet.events.GetOrSet(eventID, struct{}{})
	return isAbsent, nil
}
