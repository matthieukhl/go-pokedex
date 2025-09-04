package pokecache

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, exists := c.Entries[key]
	if !exists {
		return nil, false
	}

	return val.Val, false
}
