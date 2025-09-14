package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mutex   sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{}
	cache.entries = make(map[string]cacheEntry)
	go cache.reapLoop(interval)
	return &cache
}

func (cache *Cache) Add(key string, value []byte) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	cache.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	if entry, ok := cache.entries[key]; ok {
		return entry.val, true
	} else {
		return nil, false
	}
}

func (cache *Cache) reapLoop(interval time.Duration) {
	var reap = func() {
		cache.mutex.Lock()
		defer cache.mutex.Unlock()

		threshold := time.Now().Add(-interval)

		for key, entry := range cache.entries {
			if entry.createdAt.Before(threshold) {
				delete(cache.entries, key)
			}
		}
	}

	for range time.Tick(interval) {
		reap()
	}
}
