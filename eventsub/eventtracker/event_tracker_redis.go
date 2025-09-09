package eventtracker

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisKeyBuilder builds and returns key for Redis to store event with provided event identifier.
type RedisKeyBuilder func(eventID string) string

// RedisEventTracker is a standard Redis implementation of EventTracker with official Redis client, which is suitable
// for cases when you have multiple instances of your application, and you need to track events synchronously across them.
type RedisEventTracker struct {
	client   *redis.Client
	eventTTL time.Duration
	key      RedisKeyBuilder
}

var _ EventTracker = (*RedisEventTracker)(nil)

func NewRedisEventTracker(client *redis.Client, key RedisKeyBuilder, options ...Option) (RedisEventTracker, error) {
	if key == nil {
		return RedisEventTracker{}, errors.New("key builder is not provided")
	}

	opt := option{
		eventTTL: SafeEventTTL,
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

	firstTimeSeen, err := ret.client.SetNX(ctx, key, 1, ret.eventTTL).Result()
	if err != nil {
		return false, fmt.Errorf("set nx: %w", err)
	}

	isDuplicate := !firstTimeSeen
	return isDuplicate, nil
}
