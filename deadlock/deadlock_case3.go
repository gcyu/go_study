package main

import "fmt"

var ch1 = make(chan int)
var ch2 = make(chan int)

func say(s string) {
	fmt.Println(s)
	ch1 <- <- ch2 // ch1 等待 ch2流出的数据
}

func main() {
	go say("hello")
	<- ch1  // 堵塞主线
}

// 多个go routine都在等。其中主线等ch1中的数据流出，ch1等ch2的数据流出，但是ch2等待数据流入，两个goroutine都在等，也就是死锁。
