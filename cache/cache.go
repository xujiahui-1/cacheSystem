package cache

import "time"

type Cache interface {
	//SetMaxMemory 最大内存设置
	SetMaxMemory(size string) bool

	//Set 存数据 ,注意缓存超时后的内存清理
	Set(key string, val interface{}, expire time.Duration) bool

	//Get 获取数据
	Get(key string) (interface{}, bool)

	//Del 删除数据
	Del(key string) bool

	//Exists 判断数据是否存在
	Exists(key string) bool

	//Flush 清空所有数据
	Flush() bool

	//Keys 获取key数量
	Keys() int64
}
