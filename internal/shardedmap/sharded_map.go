package shardedmap

import (
	"context"
	"hash/fnv"
	"sync"
	"time"
)

const shardCount = 32

type Hasher[K comparable] func(K) uint64

type ShardedMap[K comparable, V any] struct {
	shards   [shardCount]*shard[K, V]
	hasher   Hasher[K]
	entryTTL time.Duration
}

type (
	entry[V any] struct {
		value     V
		timestamp time.Time
	}

	shard[K comparable, V any] struct {
		entries map[K]entry[V]
		sync.Mutex
	}
)

func (s *shard[K, V]) startTTLEviction(ctx context.Context, entryTTL time.Duration) {
	ticker := time.NewTicker(entryTTL / 2)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			s.Lock()
			for key, value := range s.entries {
				if time.Since(value.timestamp) >= entryTTL {
					delete(s.entries, key)
				}
			}
			s.Unlock()
		}
	}
}

func newShard[K comparable, V any]() *shard[K, V] {
	return &shard[K, V]{
		entries: make(map[K]entry[V]),
	}
}

func newShardedMap[K comparable, V any](ctx context.Context, hasher Hasher[K], entryTTL time.Duration) ShardedMap[K, V] {
	var shards [shardCount]*shard[K, V]

	for i := range shardCount {
		sh := newShard[K, V]()
		shards[i] = sh

		if entryTTL > 0 {
			go sh.startTTLEviction(ctx, entryTTL)
		}
	}

	return ShardedMap[K, V]{
		shards:   shards,
		hasher:   hasher,
		entryTTL: entryTTL,
	}
}

func newStringHasher() Hasher[string] {
	return func(key string) uint64 {
		hasher := fnv.New64()
		_, _ = hasher.Write([]byte(key))
		return hasher.Sum64()
	}
}

func NewString[V any](ctx context.Context, entryTTL time.Duration) ShardedMap[string, V] {
	return newShardedMap[string, V](ctx, newStringHasher(), entryTTL)
}

func (sm *ShardedMap[K, V]) GetOrSet(key K, value V) (V, bool) {
	shardEntry := sm.getShard(key)

	shardEntry.Lock()
	defer shardEntry.Unlock()

	if storedValue, exists := shardEntry.entries[key]; exists {
		if time.Since(storedValue.timestamp) >= sm.entryTTL {
			delete(shardEntry.entries, key)
		}

		return storedValue.value, true
	}

	shardEntry.entries[key] = entry[V]{
		value:     value,
		timestamp: time.Now(),
	}

	return value, false
}

func (sm *ShardedMap[K, V]) getShard(key K) *shard[K, V] {
	hash := sm.hasher(key)
	return sm.shards[hash%shardCount]
}
