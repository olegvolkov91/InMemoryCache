package cache

import (
	"fmt"
)

type val interface{}

type cache map[string]val

type LocalCache struct {
	data cache
}

func New() *LocalCache {
	lc := LocalCache{data: make(cache)}
	return &lc
}

func (lc *LocalCache) Set(key string, value val) {
	lc.data[key] = value
}

func (lc *LocalCache) Get(key string) (val, bool) {
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
