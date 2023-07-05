package main

import (
	"testProject/ImplementOfCacheSystem"
	"time"
)

func main() {
	//通信共享内存测试
	//channel.ConcurrentSync()
	//channel.NoticeAndMultiplexing()
	////channel.Communication()
	//ch := make(chan os.Signal, 0)
	//signal.Notify(ch, os.Interrupt, os.Kill)
	////
	//<-ch

	//文件复制
	//_case.CopyDirToDir()
	//文件写入
	//_case.ReadWriteFiles()
	//文件边读边写
	//_case.OneSideReadWriteToDest()
	//文件读取-按行读取
	//_case.ReadLine1()
	//文件读取-通过bufio按行读取
	//_case.ReadLine2()
	//文件读取 -通过scanner按行读取
	//_case.ReadLine3()

	//原子操作
	//_case.AtomicCase2()

	//context控制协程退出
	//_case.ContextCase()
	//ctx, stop := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
	//defer stop()
	//<-ctx.Done()

	//泛型使用
	//_case.SimpleCase()

	//通用占位符
	//character.TestCharacterNumber()

	//cache缓存使用示例
	cache := ImplementOfCacheSystem.NewMemCache()
	cache.SetMaxMemory("300B")

	cache.Set("int", 1, time.Second)
	cache.Set("bool", false, time.Second)
	cache.Set("data", map[string]interface{}{"a": 1}, time.Second)
}

//func NewMemCache() ImplementOfCacheSystem.Cache {
//
//}
