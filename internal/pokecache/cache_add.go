package pokecache

import (
	"fmt"
	"time"
)

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Println("===============")
	fmt.Printf("Caching url: %s\n", key)
	fmt.Println("===============")
	c.Entries[key] = cacheEntry{
		Val:       val,
		CreatedAt: time.Now(),
	}
}
