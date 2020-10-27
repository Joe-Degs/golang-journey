package main

import "fmt"

func main() {
	msg := make(chan string)

	go func() { msg <- "messages" }()

	fmt.Println(<-msg)
}

// a goroutine can send messages via any channel.
// a channel is like a communication bridge between running goroutines.
// it makes it possible to pass messages or values or variables or resources between running goroutines.
// the main function is also a goroutine by the way.
// go routines run on different threads or so.
// goroutines are lightweight threads that help to craft brilliantly concurrent programs.
// learning to reason about them is quite difficult for the beginner so ive come to realise.
// but once it gets locked down by the brain cells, going back to other methods of managing concurrency does not make sense any longer
// different goroutines communicate using channels, once a message is passed into the channel, it is broadcasted to other goroutines.

// one important thing is that because the channel is used for communication between multiple
// goroutines and goroutines are lightweight threads, data races might be possible but no its
// not possible with channels in the picture. reason being that when a channel is recieving
// data or a resource, it locks it down, no other thing can touch that resource or do anything
// with it unless the channel releases it which it only does when its sending that resource
// to other goroutines
