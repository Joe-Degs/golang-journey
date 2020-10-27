// Sleeping sender
package main

func main() {
	messages := make(chan string)

	// goroutine that does nothing
	go func() {}()

	// send message into channel.
	// theres nobody to give the message to.
	// and the channel is not a buffered channel.
	// results in a deadlock because there is no reciever
	messages <- "Messages that nobody wants to recieve"
}
