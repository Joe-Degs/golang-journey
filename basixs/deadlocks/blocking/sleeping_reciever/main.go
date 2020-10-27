package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)

	// reciever waits for message to be transmitted
	go func() {
		fmt.Println("Waiting to recieve from channel")
		msg := <-messages
		fmt.Println("Message recieved")
		fmt.Println(msg)
	}()

	// sender takes a 22 second nap before transmitting
	time.Sleep(2 * time.Second)
	messages <- "Message: I have arrived"

	// wait for goroutine
	time.Sleep(1 * time.Second)
}
