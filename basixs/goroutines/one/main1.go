package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// allocate one logical processor for the scheduler to use
	runtime.GOMAXPROCS(1)

	// wg is used to wait for a goroutine to finish.
	// Add a count of two, one for each goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Starting Goroutines")

	// declare anonymous function and create goroutine
	go func() {
		// schedule the call to Done to tell main we are done
		defer wg.Done()

		// display the english alphabet three times
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				if count == 2 && char == 'z' {
					fmt.Printf("%c\n", char)
				} else {
					fmt.Printf("%c ", char)
				}
			}
		}
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				if count == 2 && char == 'Z' {
					fmt.Printf("%c\n", char)
				} else {
					fmt.Printf("%c ", char)
				}
			}
		}
	}()

	// Wait for goroutines to complete executing
	wg.Wait()

	fmt.Println("\nTerminating program")

}
