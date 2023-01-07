package main

import (
	"fmt"
	"inmemory-cache/cache"
	"time"
)

func main() {
	key := "userId"
	localCache := cache.New()

	localCache.Set(key, 123, 5*time.Second)

	user, ok := localCache.Get(key)

	if !ok {
		fmt.Printf("%s such key does not exist in cache\n", key)
		return
	}

	fmt.Println(user)

	// localCache2 := cache.New()

	fmt.Println("LocalCache1", localCache)
	// fmt.Println("LocalCache2", localCache2)

	time.Sleep(time.Second * 7)
	_, ok = localCache.Get(key)

	if !ok {
		fmt.Printf("%s such key does not exist in cache\n", key)
		return
	}
	// fmt.Println("KEY STILL EXISTS", res)
	// localCache.Delete(key)

	// fmt.Println("LocalCache1 after deletion", localCache)

	// check that both caches has different pointers
	// fmt.Println(&localCache, &localCache2)
}
