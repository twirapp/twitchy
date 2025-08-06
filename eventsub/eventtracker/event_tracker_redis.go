package eventtracker

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisKeyBuilder interface {
	// Build builds key for redis to store event with provided EventID.
	Build(eventID EventID) string
}

// RedisEventTracker is a standard implementation of EventTracker for Redis, which is suitable for cases where you have
// multiple instances of an application and need to track events in them synchronously.
type RedisEventTracker struct {
	client     *redis.Client
	eventTTL   time.Duration
	keyBuilder RedisKeyBuilder
}

var _ EventTracker = (*RedisEventTracker)(nil)

func NewRedisEventTracker(client *redis.Client, keyBuilder RedisKeyBuilder, options ...Option) (RedisEventTracker, error) {
	opt := option{
		eventTTL: EventTTL,
	}

	for _, userOpt := range options {
		userOpt(&opt)
	}

	return RedisEventTracker{
		client:     client,
		eventTTL:   opt.eventTTL,
		keyBuilder: keyBuilder,
	}, nil
}

func (ret RedisEventTracker) Track(ctx context.Context, eventID string) (bool, error) {
	key := ret.keyBuilder.Build(eventID)

	firstTimeSeen, err := ret.client.SetNX(ctx, key, 1, ret.eventTTL).Result()
	if err != nil {
		return false, fmt.Errorf("set nx: %w", err)
	}

	isDuplicate := !firstTimeSeen
	return isDuplicate, nil
}
