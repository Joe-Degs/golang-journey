package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		fmt.Println("Func: Trying to recieve") // this will execute
		<-messages                             // this will block
		fmt.Println("Recieved to recieve")     // never executed
	}()

	fmt.Println("Main: Trying to recieve")
	<-messages
	fmt.Println("Never executed")
}
