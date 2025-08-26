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
	mux *sync.Mutex
	val map[string]cacheEntry
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		val: make(map[string]cacheEntry),
		mux: &sync.Mutex{},
	}
	go newCache.reapLoop(interval)

	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.val[key] = newEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	cachedValue, ok := c.val[key]

	if !ok {
		return nil, false
	}

	return cachedValue.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.val {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.val, k)
		}
	}
}
