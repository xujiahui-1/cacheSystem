package cache_server

import (
	"cacheSystem/cache"
	"time"
)

type cacheServer struct {
	memcache cache.Cache
}

func NewMemCache() *cacheServer {
	return &cacheServer{
		memcache: cache.NewMemCache(),
	}
}

func (c *cacheServer) SetMaxMemory(size string) bool {
	return c.memcache.SetMaxMemory(size)
}

func (c *cacheServer) Set(key string, val interface{}, expire ...time.Duration) bool {
	timeTls := time.Second * 0
	if len(expire) > 0 {
		timeTls = expire[0]
	}

	return c.memcache.Set(key, val, timeTls)
}

func (c *cacheServer) Get(key string) (interface{}, bool) {
	return c.memcache.Get(key)
}

func (c *cacheServer) Del(key string) bool {
	return c.memcache.Del(key)
}

func (c *cacheServer) Exists(key string) bool {
	return c.memcache.Exists(key)
}

func (c *cacheServer) Flush() bool {
	return c.memcache.Flush()
}

func (c *cacheServer) Keys() int64 {
	return c.memcache.Keys()
}
