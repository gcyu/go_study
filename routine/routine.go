package main

import (
	"fmt"
	"time"
)

func loop() {
	for i := 0; i < 100; i++ {
		fmt.Printf("%d ", i)
	}
}

func main() {
	go loop()
	go loop()
	time.Sleep(time.Second) // sleep 1 second to let main process not exit immediately, so wo can observe the output of that 2 go routines
}
