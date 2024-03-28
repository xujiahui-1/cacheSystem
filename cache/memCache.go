package cache

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type memCache struct {
	maxMemorySizeStr string

	maxMemorySize int64

	currMemorySize int64

	values map[string]*memCacheValue

	//map的原子性
	locker sync.RWMutex

	//定期清楚缓存的时间间隔
	clearExpiredItemTime time.Duration
}

type memCacheValue struct {
	//value值
	val interface{}
	//过期时间
	expire time.Time
	// 有效时长
	expireTime time.Duration
	//value大小
	size int64
}

func NewMemCache() Cache {
	mc := &memCache{
		values:               make(map[string]*memCacheValue),
		clearExpiredItemTime: time.Second * 10,
	}
	go mc.clearExpiredItem()
	return mc
}

// SetMaxMemory Size 1kb 1GB ....
func (m *memCache) SetMaxMemory(size string) bool {
	m.maxMemorySize, m.maxMemorySizeStr = ParseSize(size)
	fmt.Println(m.maxMemorySize, m.maxMemorySizeStr)
	return true
}

func (m *memCache) Set(key string, val interface{}, expire time.Duration) bool {

	v := &memCacheValue{
		val: val,
		//过期时间：当前时间加缓存时间
		expire:     time.Now().Add(expire),
		expireTime: expire,
		size:       m.GetValSize(val),
	}
	m.locker.Lock()
	m.del(key)
	add := m.add(key, v)
	defer m.locker.Unlock()

	return add
}

func (m *memCache) get(key string) (*memCacheValue, bool) {
	value, ok := m.values[key]
	return value, ok
}
func (m *memCache) del(key string) {
	val, ok := m.get(key)
	if ok && val != nil {
		m.currMemorySize -= val.size
		delete(m.values, key)
	}
}
func (m *memCache) add(key string, value *memCacheValue) bool {
	if m.currMemorySize+value.size > m.maxMemorySize {
		log.Println("内存不足，无法添加")
		panic("内存不足，无法添加")
		return false
	}
	m.values[key] = value
	m.currMemorySize += value.size
	return true
}
func (m *memCache) Get(key string) (interface{}, bool) {
	m.locker.RLock()
	defer m.locker.RUnlock()
	val, ok := m.get(key)
	if ok {
		if val.expireTime != 0 && time.Now().After(val.expire) {
			log.Println("缓存以过期")
			delete(m.values, key)
			return nil, false
		} else {
			return val.val, ok
		}

	}
	return nil, false
}

func (m *memCache) Del(key string) bool {
	m.locker.Lock()
	defer m.locker.Unlock()
	m.del(key)
	return true
}

func (m *memCache) Exists(key string) bool {
	m.locker.RLock()
	defer m.locker.RUnlock()
	val, ok := m.get(key)
	if ok {
		if val.expireTime != 0 && time.Now().After(val.expire) {
			log.Println("缓存以过期")
			return false
		} else {
			return ok
		}

	}
	return false
}

func (m *memCache) Flush() bool {
	m.locker.Lock()
	defer m.locker.Unlock()
	m.values = make(map[string]*memCacheValue, 0)
	m.currMemorySize = 0
	return true
}

func (m *memCache) Keys() int64 {
	return 1
}

func (m *memCache) GetValSize(val interface{}) int64 {
	m.locker.RLock()
	defer m.locker.RUnlock()
	fmt.Println(len(m.values))
	return int64(len(m.values))

}

// 定期清楚缓存
func (m *memCache) clearExpiredItem() {
	ticker := time.NewTicker(m.clearExpiredItemTime)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			for key, item := range m.values {
				if item.expireTime != 0 && time.Now().After(item.expire) {
					m.locker.Lock()
					m.del(key)
					m.locker.Unlock()
				}
			}
		}
	}
}
