package main

import "fmt"

func main() {
	var rc chan string = make(chan string)
	//var sc <-chan string = make(<-chan string)
	recv_chan(rc)
	send_chan(rc)
}

func recv_chan(c chan<- string) {
	c <- "string"
}

func send_chan(c <-chan string) {
	fmt.Println(<-c)
}
