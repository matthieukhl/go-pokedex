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

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		Entries: make(map[string]cacheEntry),
		mu:      sync.Mutex{},
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.ReapLoop(interval)

	return c
}
