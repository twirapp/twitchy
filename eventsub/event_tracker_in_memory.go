package eventsub

import (
	"context"
	"time"

	"github.com/kvizyx/twitchy/internal/concurrentmap"
)

// InMemoryEventTracker is a concurrent-safe in-memory implementation of the EventTracker which is you can use by default
// if you don't need to track events across multiple instances of your application.
type InMemoryEventTracker struct {
	events concurrentmap.ConcurrentMap[string, struct{}]
}

var _ EventTracker = (*InMemoryEventTracker)(nil)

func NewInMemoryEventTracker(ctx context.Context, eventTTL time.Duration) *InMemoryEventTracker {
	return &InMemoryEventTracker{
		events: concurrentmap.NewString[struct{}](ctx, eventTTL),
	}
}

func (iet *InMemoryEventTracker) Track(_ context.Context, eventID string) (bool, error) {
	_, isAbsent := iet.events.GetOrSet(eventID, struct{}{})
	return isAbsent, nil
}
