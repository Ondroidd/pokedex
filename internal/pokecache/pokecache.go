package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	var new_cache = Cache{
		entries: make(map[string]cacheEntry),
	}

	go new_cache.reapLoop(interval)
	return &new_cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	new_entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.entries[key] = new_entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	for {
		c.mu.Lock()
		for key, value := range c.entries {
			if time.Since(value.createdAt) > interval {
				delete(c.entries, key)
			}

		}
		c.mu.Unlock()
		time.Sleep(interval)
	}
}
