package eventtracker

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisClient is a client for Redis that is used in RedisEventTracker to enable the use of different Redis clients if needed.
type RedisClient interface {
	// SetIfNotExists sets a value for a key if the key does not yet exist (use SetNX as an idiomatic way to do this).
	SetIfNotExists(ctx context.Context, key string, value int, expiration time.Duration) (bool, error)
}

// DefaultRedisClient is a default implementation of RedisClient with official Redis client.
type DefaultRedisClient struct {
	client *redis.Client
}

func NewDefaultRedisClient(client *redis.Client) DefaultRedisClient {
	return DefaultRedisClient{
		client: client,
	}
}

func (drc DefaultRedisClient) SetIfNotExists(
	ctx context.Context,
	key string,
	value int,
	expiration time.Duration,
) (bool, error) {
	command := drc.client.SetNX(ctx, key, value, expiration)
	return command.Result()
}
