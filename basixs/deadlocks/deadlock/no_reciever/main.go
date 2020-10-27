package main

func main() {
	messages := make(chan string)

	go func() {}()

	// main function trying to recieve from channel.
	// but there is nothing to recieve
	// this causes a deadlock
	<-messages
}

//deadlocks happen when two goroutines are waiting for each other and none is able to proceed.
