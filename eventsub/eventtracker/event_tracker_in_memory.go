package eventtracker

import (
	"context"

	"github.com/kvizyx/twitchy/internal/shardedmap"
)

// InMemoryEventTracker is a standard in-memory concurrent safe implementation of EventTracker based on
// shardedmap.ShardedMap, which is suitable for cases where it is not necessary to track events synchronously in
// multiple instances of your application, so events can be stored in the process memory.
type InMemoryEventTracker struct {
	events shardedmap.ShardedMap[string, struct{}]
}

var _ EventTracker = (*InMemoryEventTracker)(nil)

func NewInMemoryEventTracker(ctx context.Context, options ...Option) *InMemoryEventTracker {
	opt := option{
		eventTTL: SafeEventTTL,
	}

	for _, userOpt := range options {
		userOpt(&opt)
	}

	return &InMemoryEventTracker{
		events: shardedmap.NewString[struct{}](ctx, opt.eventTTL),
	}
}

func (iet *InMemoryEventTracker) Track(_ context.Context, eventID string) (bool, error) {
	_, isDuplicate := iet.events.GetOrSet(eventID, struct{}{})
	return isDuplicate, nil
}
