package pokecache

import (
	"fmt"
	"sync"
	"time"
)

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]CacheEntry),
		mu:    &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	fmt.Println(fmt.Sprintf(`%s is added to cache`, key))

	c.cache[key] = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	result, isFound := c.cache[key]
	fmt.Println(fmt.Sprintf(`Requested key:%v  is found:%v`, key, isFound))
	return result.val, isFound
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		for k, v := range c.cache {
			if v.createdAt.Before(time.Now().Add(-interval)) {
				delete(c.cache, k)
			}
		}
		c.mu.Unlock()
	}
}
