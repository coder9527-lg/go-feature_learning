
设计实现一个内存缓存的系统
要求：
1、支持设定过期时间，精度到秒
2、支持设定最大内存，当内存超出时做出合适的处理
3、支持并发安全
4、按照以下接口要求实现

type Cache interface{
    //size ：1 KB，100KB，1MB，2MB，1GB
    SetMaxMemory(size string) bool
    //将value写入缓存
    Set(key string,val interface{},expire time.Duration) bool
    //根据key值获取value
    Get(key string)(interface{},bool)
    //删除key值
    Del(key string)bool
    //判断key值是否存在
    Exists(key string) bool
    //情况所有的key
    Flush() bool
    //获取缓存中所有的key的数量
    Keys() int64
}