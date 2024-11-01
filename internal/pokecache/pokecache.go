package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createAt time.Time
	val      []byte
}

type Cache struct {
	data map[string]cacheEntry
	mux  *sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	c.data[key] = cacheEntry{
		createAt: time.Now(),
		val:      val,
	}
	c.mux.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	data, exists := c.data[key]
	if exists {
		return data.val, exists
	}
	return []byte{}, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mux.Lock()
		data := c.data
		for key, val := range data {
			if val.createAt.Before(time.Now().UTC().Add(-interval)) {
				delete(c.data, key)
			}
		}
		c.mux.Unlock()
	}
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		data: map[string]cacheEntry{},
		mux:  &sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return cache

}
