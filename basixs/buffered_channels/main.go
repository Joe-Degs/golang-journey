package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const (
	numberGoroutines = 4  // Number of goroutines
	taskLoad         = 10 // Amount of work to process
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	// buffered channel to manage the taskLoad
	tasks := make(chan string, taskLoad)

	// launch goroutines to manage the work
	wg.Add(numberGoroutines)
	for gr := 0; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}
	fmt.Printf("We have %d goroutines running\n", runtime.NumGoroutine())

	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task: %d", post)
	}

	// close the channel so the goroutines will quit
	// when all the work is done.
	close(tasks)

	// wait for all the work to get done.
	wg.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			// this means channel is empty and closed.
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}

		// Display we are starting the work
		fmt.Printf("Worker: %d : Started %s\n", worker, task)

		// Randomly wait to simulate work time
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// Display we finished the work
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}
