package pokecache

import (
	"time"
)

func NewCache(it time.Duration) *Cache {
	cache := &Cache{
		cm:   make(map[string]cacheEntry),
		tick: time.NewTicker(it),
	}
	go cache.reapLoop()
	return cache
}

func (c *Cache) reapLoop() {
	for range c.tick.C {
		for key, item := range c.cm {
			if time.Since(item.createdAt) > 5 {
				c.mu.Lock()
				delete(c.cm, key)
				c.mu.Unlock()
			}
		}
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cm[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	v, ok := c.cm[key]
	if !ok {
		return nil, false
	}

	return v.val, true
}
