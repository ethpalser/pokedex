package cache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	data map[string]cacheEntry
	mux  *sync.RWMutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		data: map[string]cacheEntry{},
		mux:  &sync.RWMutex{},
	}
	c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()

	dat, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return dat.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for {
			// Block until tick recieved
			t := <-ticker.C
			c.mux.Lock()
			// Delete all records older than interval
			for index, entry := range c.data {
				if entry.createdAt.Add(interval).Compare(t) < 0 {
					delete(c.data, index)
				}
			}
			c.mux.Unlock()
		}
	}()
}
