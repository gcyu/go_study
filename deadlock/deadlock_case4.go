package main

func main() {
	c, quit := make(chan int), make(chan int)

	go func() {
		c <- 1  // c通道的数据没有被其他goroutine读取走，堵塞当前goroutine
		quit <- 0 // quit始终没有办法写入数据
	}()

	<- quit // quit 等待数据的写
}

// 其实，总结来看，为什么会死锁？非缓冲信道上如果发生了流入无流出，或者流出无流入，也就导致了死锁。或者这样理解 Go启动的所有goroutine里的非缓冲信道一定要一个线里存数据，一个线里取数据，要成对才行

// 仔细分析的话，是由于：主线等待quit信道的数据流出，quit等待数据写入，而func被c通道堵塞，所有goroutine都在等，所以死锁。

// 简单来看的话，一共两个线，func线中流入c通道的数据并没有在main线中流出，肯定死锁。
