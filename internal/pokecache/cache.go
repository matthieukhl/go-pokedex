package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entries map[string]cacheEntry // List of cached data
	mu      sync.Mutex
}

type cacheEntry struct {
	Val       []byte    // Raw data we are caching
	CreatedAt time.Time // Represents when the entry was created
}

func NewCache(interval time.Duration) Cache {
	return Cache{}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Entries[key] = cacheEntry{
		Val:       val,
		CreatedAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	val, exists := c.Entries[key]
	if !exists {
		return nil, false
	}

	return val.Val, false
}

func (c *Cache) ReapLoop() {}
