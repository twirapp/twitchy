package eventtracker

import (
	"context"

	"github.com/kvizyx/twitchy/internal/concurrentmap"
)

// InMemoryEventTracker is a standard implementation of EventTracker in-memory, safe for concurrent execution, based on
// concurrentmap.ConcurrentMap, which is suitable for cases where it is not necessary to track events synchronously in
// multiple instances of the application, so events can be stored in the process memory.
type InMemoryEventTracker struct {
	events concurrentmap.ConcurrentMap[string, struct{}]
}

var _ EventTracker = (*InMemoryEventTracker)(nil)

func NewInMemoryEventTracker(ctx context.Context, options ...Option) *InMemoryEventTracker {
	opt := option{
		eventTTL: EventTTL,
	}

	for _, userOpt := range options {
		userOpt(&opt)
	}

	return &InMemoryEventTracker{
		events: concurrentmap.NewString[struct{}](ctx, opt.eventTTL),
	}
}

func (iet *InMemoryEventTracker) Track(_ context.Context, eventID string) (bool, error) {
	_, isDuplicate := iet.events.GetOrSet(eventID, struct{}{})
	return isDuplicate, nil
}
