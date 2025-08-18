package eventtracker

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// RedisKeyBuilder builds and returns key for Redis to store event with provided event identifier.
type RedisKeyBuilder func(eventID string) string

// RedisEventTracker is a standard implementation of EventTracker for Redis, which is suitable for cases when you have
// multiple instances of an application and need to track events in them synchronously.
type RedisEventTracker struct {
	client   RedisClient
	eventTTL time.Duration
	key      RedisKeyBuilder
}

var _ EventTracker = (*RedisEventTracker)(nil)

func NewRedisEventTracker(client RedisClient, key RedisKeyBuilder, options ...Option) (RedisEventTracker, error) {
	if key == nil {
		return RedisEventTracker{}, errors.New("key builder is not provided")
	}

	opt := option{
		eventTTL: EventTTL,
	}

	for _, userOpt := range options {
		userOpt(&opt)
	}

	return RedisEventTracker{
		client:   client,
		eventTTL: opt.eventTTL,
		key:      key,
	}, nil
}

func (ret RedisEventTracker) Track(ctx context.Context, eventID string) (bool, error) {
	key := ret.key(eventID)

	firstTimeSeen, err := ret.client.SetIfNotExists(ctx, key, 1, ret.eventTTL)
	if err != nil {
		return false, fmt.Errorf("set if not exists: %w", err)
	}

	isDuplicate := !firstTimeSeen
	return isDuplicate, nil
}
