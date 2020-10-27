package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)

	// spawned goroutine sleeps for sometime before sending message into the channel.
	go func() {
		time.Sleep(1 * time.Second)
		messages <- "Hello"
	}()

	// recieving from channel can't be done for sometime because there is no value to recieve
	// it is blocked
	fmt.Println(<-messages)
}
