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
	val       []byte    // Raw data we are caching
	createdAt time.Time // Represents when the entry was created
}

func NewCache(interval time.Duration) Cache {
	return Cache{}
}

func (c *Cache) Add(key string, val []byte) {}

func (c *Cache) Get(key string) ([]byte, bool) {
	return nil, false
}

func (c *Cache) ReapLoop() {}
