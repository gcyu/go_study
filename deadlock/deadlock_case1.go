package main

func main() {
	ch := make(chan int)
	<- ch // 阻塞main goroutine, 信道c被锁
}

// 何谓死锁? 操作系统有讲过的，所有的线程或进程都在等待资源的释放。如上的程序中, 只有一个goroutine, 所以当你向里面加数据或者存数据的话，都会锁死信道并且阻塞当前goroutine, 也就是所有的goroutine(其实就main线一个)都在等待信道的开放(没人拿走数据信道是不会开放的)，也就是死锁。
