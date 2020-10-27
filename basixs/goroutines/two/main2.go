package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// allocating one logical processor for the scheduler to use
	runtime.GOMAXPROCS(1)

	// declaring the semaphor used to wait for program to finish
	var wg sync.WaitGroup

	// increment the semaphor to two
	wg.Add(2)

	// creating two goroutines
	fmt.Println("Starting goroutines")

	go printPrime("A", &wg)
	go printPrime("B", &wg)

	// wait for goroutines to finish
	fmt.Println("Waiting to Finish")
	wg.Wait()

	fmt.Println("Terminating program")
}

func printPrime(prefix string, w *sync.WaitGroup) {
	defer w.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s: %d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)

}
