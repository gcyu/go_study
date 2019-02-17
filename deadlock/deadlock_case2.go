package main

import "fmt"

func main() {
	ch := make(chan int)
	//go func() {
	//	<- ch
	//}()
	ch <- 1 // 1流入信道，堵塞当前线, 没人取走数据信道不会打开
	fmt.Println("This line code wont run") //在此行执行之前Go就会报死锁
}

// 只在单一的goroutine里操作无缓冲信道，一定死锁。比如你只在main函数里操作信道
