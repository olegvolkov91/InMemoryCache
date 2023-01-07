package cache

import (
	"fmt"
	"sync"
	"time"
)

type cache map[string]interface{}

type LocalCache struct {
	data cache
	mu   *sync.RWMutex
}

func New() *LocalCache {
	lc := LocalCache{data: make(cache), mu: new(sync.RWMutex)}
	return &lc
}

func cleanCache(key string, lc *LocalCache) func() {
	return func() {
		if _, ok := lc.Get(key); !ok {
			return
		}
		lc.Delete(key)
	}
}

func (lc *LocalCache) Set(key string, value interface{}, ttl time.Duration) {
	lc.data[key] = value
	time.AfterFunc(ttl, cleanCache(key, lc))
}

func (lc *LocalCache) Get(key string) (interface{}, bool) {
	val, ok := lc.data[key]
	if !ok {
		return 0, false
	}
	return val, true
}

func (lc *LocalCache) Delete(key string) {
	_, ok := lc.data[key]
	if !ok {
		fmt.Printf("%s does not exists in cache", key)
		return
	}
	delete(lc.data, key)
}
