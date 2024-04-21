package cache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

type Cache struct {
	cache    map[string]cacheEntry
	interval time.Duration
	lock     sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		cache:    make(map[string]cacheEntry),
		interval: interval,
	}

	go c.reapLoop()
	return &c
}

func (c *Cache) Add(key *string, value []byte) {
	c.lock.Lock()
	defer c.lock.Unlock()

	_, ok := c.cache[*key]
	if ok {
		return
	}

	c.cache[*key] = cacheEntry{
		value:     value,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key *string) ([]byte, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	entry, ok := c.cache[*key]
	if !ok {
		fmt.Println("cache miss for key:", key)
		return nil, false
	}

	fmt.Println("cache hit for key:", key)
	return entry.value, true
}

func (c *Cache) reap() {
	c.lock.Lock()
	defer c.lock.Unlock()
	for k, v := range c.cache {
		if time.Since(v.createdAt) >= c.interval {
			delete(c.cache, k)
		}
	}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.reap()
	}
}
