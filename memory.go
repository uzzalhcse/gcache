package cache

import (
	"errors"
	"sync"
	"time"
)

var ErrItemNotFound = errors.New("cache: item not found")

type ItemData struct {
	value     interface{}
	createdAt int64
	ttl       int64
}

type InMemoryCache struct {
	items map[interface{}]*ItemData
	sync.RWMutex
}

func NewInMemoryCache() *InMemoryCache {
	c := &InMemoryCache{items: make(map[interface{}]*ItemData)}
	go c.startTTLCleanup()

	return c
}

func (c *InMemoryCache) startTTLCleanup() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		c.Lock()
		for key, item := range c.items {
			if time.Now().Unix()-item.createdAt > item.ttl {
				delete(c.items, key)
			}
		}
		c.Unlock()
	}
}

func (c *InMemoryCache) Set(key, value interface{}, ttl int64) {
	c.Lock()
	defer c.Unlock()

	c.items[key] = &ItemData{
		value:     value,
		createdAt: time.Now().Unix(),
		ttl:       ttl,
	}
}

func (c *InMemoryCache) Get(key interface{}) (interface{}, error) {
	c.RLock()
	defer c.RUnlock()

	item, exists := c.items[key]
	if !exists {
		return nil, ErrItemNotFound
	}

	return item.value, nil
}
