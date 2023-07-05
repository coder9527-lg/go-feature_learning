package ImplementOfCacheSystem

import "time"

type Cache interface {
	//size ：1 KB，100KB，1MB，2MB，1GB
	SetMaxMemory(size string) bool
	//将value写入缓存
	Set(key string, val interface{}, expire time.Duration) bool
	//根据key值获取value
	Get(key string) (interface{}, bool)
	//删除key值
	Del(key string) bool
	//判断key值是否存在
	Exists(key string) bool
	//情况所有的key
	Flush() bool
	//获取缓存中所有的key的数量
	Keys() int64
}
