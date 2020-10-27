package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func incCounter(id string) {
	fmt.Printf("%s is running\n", id)
	defer wg.Done()

	for count := 0; count < 2; count++ {
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock()
	}
}

func main() {
	wg.Add(2)

	go incCounter("A")
	go incCounter("B")

	wg.Wait()
	fmt.Printf("Counter is now %d\n", counter)
	fmt.Printf("Done\n")
}
