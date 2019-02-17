package main

import "fmt"

func main() {
	var messages = make(chan string)
	go func(message string) {
		messages <- message // save a message into channel
	}("Ping!")
	fmt.Println(<-messages) // pop a message from channel
}
