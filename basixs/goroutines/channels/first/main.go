package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var w sync.WaitGroup

func main() {
	var court chan int = make(chan int)
	w.Add(2)

	go player("Joseph", court)
	go player("Bernice", court)
	court <- 1
	w.Wait()
	fmt.Println("Done")
}

func player(name string, court chan int) {
	defer w.Done()

	for {
		// wait for ball to hit back at us
		ball, ok := <-court
		if !ok {
			// if the channel is closed we won
			fmt.Printf("Player %s won", name)
			return
		}

		// pick a random number
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			runtime.Gosched()

			// close the channel to signal we lost
			close(court)
			return
		}

		// display and increment the hit count by one
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
	}
}
