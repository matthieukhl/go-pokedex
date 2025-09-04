package pokecache

import "time"

func (c *Cache) ReapLoop(interval time.Duration) {
	ttl := time.Now().Add(-interval)

	for k, v := range c.Entries {
		if v.CreatedAt.Unix() < ttl.Unix() {
			delete(c.Entries, k)
		}
	}
}
