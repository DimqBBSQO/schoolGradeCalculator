package cache

import (
	"errors"
	"sync"
	"time"
)

var ErrItemNotFound = errors.New("cache: item not found")

type item struct {
	value     interface{}
	createdAt int64
	ttl       int64
}

type MemoryCache struct {
	cache map[interface{}]*item
	sync.RWMutex
}

func New() *MemoryCache {
	c := &MemoryCache{cache: make(map[interface{}]*item)}
	go c.setTtlTimer()

	return c
}

func (c *MemoryCache) setTtlTimer() {
	for {
		c.Lock()
		for k, v := range c.cache {
			if v.ttl == 0 {
				continue
			}

			if time.Now().Unix()-v.createdAt > v.ttl {
				delete(c.cache, k)
			}
		}
		c.Unlock()

		<-time.After(time.Second)
	}
}

func (c *MemoryCache) Delete(key interface{}) error {
	c.Lock()
	defer c.Unlock()

	if _, ok := c.cache[key]; !ok {
		return ErrItemNotFound
	}
	delete(c.cache, key)

	return nil
}

func (c *MemoryCache) Get(key interface{}) (interface{}, error) {
	c.RLock()
	defer c.RUnlock()

	item, ex := c.cache[key]

	if !ex {
		return nil, ErrItemNotFound
	}

	return item.value, nil
}

func (c *MemoryCache) Set(key, value interface{}, ttl int64) error {
	c.Lock()
	defer c.Unlock()

	c.cache[key] = &item{
		value:     value,
		createdAt: time.Now().Unix(),
		ttl:       ttl,
	}

	return nil
}
