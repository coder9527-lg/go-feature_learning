package main

import (
	"os"
	"os/signal"
	"testProject/channel"
)

func main() {
	//fmt.Println("hello world")
	//fmt.Println(_case.Fib(10))
	channel.ConcurrentSync()
	channel.NoticeAndMultiplexing()
	//channel.Communication()
	ch := make(chan os.Signal, 0)
	signal.Notify(ch, os.Interrupt, os.Kill)

	<-ch
}
