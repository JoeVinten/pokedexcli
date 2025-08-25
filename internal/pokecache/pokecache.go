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
	mu       sync.Mutex
	val      map[string]cacheEntry
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{
		val:      make(map[string]cacheEntry),
		interval: interval,
	}
	go newCache.reapLoop()

	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.val[key] = newEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cachedValue, ok := c.val[key]

	if !ok {
		return nil, false
	}

	return cachedValue.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for {
		<-ticker.C
		now := time.Now()
		c.mu.Lock()
		for key, entry := range c.val {
			expiryTime := entry.createdAt.Add(c.interval)
			if now.After(expiryTime) {
				delete(c.val, key)
			}
		}
		c.mu.Unlock()
	}
}
