package cache

import (
	"sync"
	"time"
)

type cacheItemValue interface{}

type CacheItem struct {
	value      cacheItemValue
	expiration time.Time
}

type Cache struct {
	storage map[string]CacheItem
	sync    *sync.RWMutex
}

func New() Cache {
	cache := Cache{
		storage: make(map[string]CacheItem),
		sync:    new(sync.RWMutex),
	}
	go cacheCleaner(&cache)
	return cache
}

func (c Cache) Set(key string, value cacheItemValue, ttl time.Duration) {
	c.storage[key] = CacheItem{
		value:      value,
		expiration: time.Now().Add(ttl),
	}
}

func (c Cache) Get(key string) (cacheItemValue, bool) {
	cacheItem, exists := c.storage[key]
	return cacheItem.value, exists
}

func (c Cache) Delete(key string) {
	delete(c.storage, key)
}

func cacheCleaner(cache *Cache) {
	for {
		for key, value := range cache.storage {
			if isExpired(value.expiration) {
				cache.Delete(key)
			}
		}
	}
}

func isExpired(expirationTime time.Time) bool {
	return time.Now().After(expirationTime)
}
