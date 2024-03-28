package main

import (
	cache_server "cacheSystem/cache-server"
	"fmt"
	"time"
)

func main() {
	mycache := cache_server.NewMemCache()
	memory := mycache.SetMaxMemory("100MB")
	fmt.Println(memory)
	set := mycache.Set("xxxx", 1234445, time.Second*1)
	fmt.Println(set)
	get, b := mycache.Get("xxxx")
	fmt.Println(get, b)

	del := mycache.Del("xxxx")
	fmt.Println(del)
	exists := mycache.Exists("xxxx")
	fmt.Println(exists)
	mycache.Flush()
	mycache.Keys()
}
