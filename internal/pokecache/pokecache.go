package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu       sync.Mutex
	cache    map[string]cacheEntry
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(t time.Duration) *Cache {
	cache := Cache{
		cache:    make(map[string]cacheEntry),
		interval: t,
	}
	cache.reapLoop()
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.cache[key]
	if !exists {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	go func() {
		for range ticker.C {
			c.mu.Lock()
			for k, v := range c.cache {
				if time.Since(v.createdAt) > c.interval {
					delete(c.cache, k)
				}
			}
			c.mu.Unlock()
		}
	}()
}
