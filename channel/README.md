package channel


#应用场景
1、协程间通信，即协程间数据传递
2、并发场景下利用channel的阻塞机制，做为同步机制
3、利用channel关闭时发送广播的特性，作为协程退出通知


#channel 通过通讯共享内存
1、channel的方向，读、写、读写
2、channel协程间通信信道
3、channel阻塞协程
4、channel 并发场景下的同步机制
5、channel 通知协程退出
6、channel 的多路复用

#注意
1、channel用于协程间通讯，必须存在读写双方，否则将造成死锁
