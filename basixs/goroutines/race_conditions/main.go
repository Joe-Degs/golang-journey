package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		//value := counter
		atomic.AddInt64(&counter, 1)

		// yield thread and be placed back in the queue
		runtime.Gosched()

		// Increment our local value of counter
		//value++

		// store the value back in counter
		//counter = value
	}
}
