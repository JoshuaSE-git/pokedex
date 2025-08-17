package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu    sync.RWMutex
	cache map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{
		mu:    sync.RWMutex{},
		cache: make(map[string]cacheEntry),
	}

	newCache.reapLoop(interval)

	return newCache
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		for {
			t := <-ticker.C
			for key, entry := range c.cache {
				if t.Sub(entry.createdAt) >= interval {
					delete(c.cache, key)
				}
			}
		}
	}()
}

func (c *Cache) Add(key string, data []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{time.Time{}, data}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	data, ok := c.cache[key]
	if !ok {
		return nil, ok
	}
	return data.val, ok
}
