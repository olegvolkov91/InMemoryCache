package main

import (
	"fmt"
	"inmemory-cache/cache"
	"time"
)

func main() {
	localCache := cache.New()
	fmt.Println(localCache)

	localCache.Set("userName", "koteykin007", time.Second*5)

	fmt.Println(localCache)

	time.Sleep(time.Second * 3)

	fmt.Println(localCache)

	value, _ := localCache.Get("userName")

	fmt.Println("userName value", value)

	time.Sleep(time.Second * 3)

	_, exists := localCache.Get("userName")

	if !exists {
		panic("Such key userName doesn't exists anymore")
	}
}
