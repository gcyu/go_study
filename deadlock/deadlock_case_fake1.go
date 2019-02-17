package main

func main() {
	c := make(chan int)

	go func() {
		c <- 1
	}()
}

// 没有死锁。程序正常退出了，很简单，并不是我们那个总结不起作用了，还是因为一个让人很囧的原因，main又没等待其它goroutine，自己先跑完了， 所以没有数据流入c信道，一共执行了一个goroutine, 并且没有发生阻塞，所以没有死锁错误。
