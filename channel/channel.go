package channel

import (
	"fmt"
	"time"
)

//协程间通信

func Communication() {
	//定义一个可读可写的通道
	ch := make(chan int, 0)
	go communicationF1(ch)
	go communicationF2(ch)
}

//F1 接收一个只写通道
func communicationF1(ch chan<- int) {
	//通过循环想通道写入 0~99
	for i := 0; i < 100; i++ {
		ch <- i
	}
}

//F2 接收一个只读通道
func communicationF2(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}

}

//ConcurrentSync 并发场景下的同步机制
func ConcurrentSync() {

	//带缓冲的通导
	ch := make(chan int, 10)
	//向ch写入消息
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()
	//向ch写入消息
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()
	//从ch读取消息并打印
	go func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()
}

//NoticeAndMultiplexing 通知协程退出与多路复用
func NoticeAndMultiplexing() {
	ch := make(chan int, 0)
	strCh := make(chan string, 0)
	done := make(chan struct{}, 0)
	go noticeAndMultiplexingF1(ch)
	go noticeAndMultiplexingF2(strCh)
	go noticeAndMultiplexingF3(ch, strCh, done)

	time.Sleep(5 * time.Second)
	close(done)
}

func noticeAndMultiplexingF1(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}

}

func noticeAndMultiplexingF2(ch chan<- string) {
	for i := 0; i < 100; i++ {
		ch <- fmt.Sprintf("数字:%d", i)
	}

}

//select 子句作为一个整体阻塞，其中任意channel准备就绪则继续执行
func noticeAndMultiplexingF3(ch <-chan int, strCh <-chan string, done <-chan struct{}) {
	i := 0
	for {
		select {
		case i := <-ch:
			fmt.Println(i)
		case str := <-strCh:
			fmt.Println(str)
		case <-done:
			fmt.Println("收到退出通知，退出当前协程")
			return
			//default:
			//	fmt.Println("执行default语句")
		}
		i++
		fmt.Println("累积执行次数", i)
	}
}
