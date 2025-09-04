package pokecache

import "time"

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Entries[key] = cacheEntry{
		Val:       val,
		CreatedAt: time.Now(),
	}
}
