package main

import _case "testProject/file-rw/case"

func main() {
	//通信共享内存测试
	//channel.ConcurrentSync()
	//channel.NoticeAndMultiplexing()
	////channel.Communication()
	//ch := make(chan os.Signal, 0)
	//signal.Notify(ch, os.Interrupt, os.Kill)
	//
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
	_case.ReadLine3()
}
