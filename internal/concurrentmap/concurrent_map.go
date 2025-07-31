package concurrentmap

import (
	"hash/fnv"
	"sync"
	"time"
)

const shardCount = 32

type Hasher[K comparable] func(K) uint64

type ConcurrentMap[K comparable, V any] struct {
	shards [shardCount]*shard[K, V]
	hasher Hasher[K]
}

type (
	entry[V any] struct {
		value     V
		timestamp time.Time
	}

	shard[K comparable, V any] struct {
		entries map[K]entry[V]
		sync.RWMutex
	}
)

func (s *shard[K, V]) startTTLEviction(entryTTL time.Duration) {
	ticker := time.NewTicker(entryTTL / 2)
	defer ticker.Stop()

	for range ticker.C {
		cutoff := time.Now().Add(-entryTTL)

		s.Lock()
		for key, ety := range s.entries {
			if ety.timestamp.Before(cutoff) {
				delete(s.entries, key)
			}
		}
		s.Unlock()
	}
}

func newShard[K comparable, V any]() *shard[K, V] {
	return &shard[K, V]{
		entries: make(map[K]entry[V]),
	}
}

func newConcurrentMap[K comparable, V any](hasher Hasher[K], entryTTL time.Duration) ConcurrentMap[K, V] {
	var shards [shardCount]*shard[K, V]

	for index := range shardCount {
		sh := newShard[K, V]()
		shards[index] = sh

		if entryTTL > 0 {
			// TODO: clean etu huetu
			go sh.startTTLEviction(entryTTL)
		}
	}

	return ConcurrentMap[K, V]{
		shards: shards,
		hasher: hasher,
	}
}

func newStringHasher() Hasher[string] {
	return func(key string) uint64 {
		hasher := fnv.New64()
		_, _ = hasher.Write([]byte(key))
		return hasher.Sum64()
	}
}

func NewString[V any](entryTTL time.Duration) ConcurrentMap[string, V] {
	return newConcurrentMap[string, V](newStringHasher(), entryTTL)
}

func (cm *ConcurrentMap[K, V]) GetOrSet(key K, value V) (V, bool) {
	shardEntry := cm.getShard(key)

	shardEntry.Lock()
	defer shardEntry.Unlock()

	if storedValue, exists := shardEntry.entries[key]; exists {
		return storedValue.value, true
	}

	shardEntry.entries[key] = entry[V]{
		value:     value,
		timestamp: time.Now(),
	}

	return value, false
}

func (cm *ConcurrentMap[K, V]) getShard(key K) *shard[K, V] {
	hash := cm.hasher(key)
	return cm.shards[hash%shardCount]
}
