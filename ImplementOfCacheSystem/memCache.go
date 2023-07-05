package ImplementOfCacheSystem

import (
	"fmt"
	"time"
)

// 用来设计一个cache缓存系统

type memCache struct {
	//最大内存
	maxMemorySize int64
	//最大内存字符串
	maxMemorySizeStr string
	//当前已使用内存
	currentMem int64
}

func NewMemCache() Cache {
	return &memCache{}

}

//size ：1 KB，100KB，1MB，2MB，1GB
func (mc *memCache) SetMaxMemory(size string) bool {
	mc.maxMemorySize, mc.maxMemorySizeStr = ParseSize(size)
	fmt.Println(mc.maxMemorySize, mc.maxMemorySizeStr)
	fmt.Println("called SetMaxMemory")
	return false
}

//将value写入缓存
func (mc *memCache) Set(key string, val interface{}, expire time.Duration) bool {
	fmt.Println("called set")
	return false
}

//根据key值获取value
func (mc *memCache) Get(key string) (interface{}, bool) {
	fmt.Println("called get")
	return nil, false
}

//删除key值
func (mc *memCache) Del(key string) bool {
	fmt.Println("called del")
	return false
}

//判断key值是否存在
func (mc *memCache) Exists(key string) bool {
	fmt.Println("called Exists")
	return false
}

//清空所有的key
func (mc *memCache) Flush() bool {
	fmt.Println("called Flush")
	return false
}

//获取所有缓存中的key的数量
func (mc *memCache) Keys() int64 {
	fmt.Println("called Keys")
	return 0
}
