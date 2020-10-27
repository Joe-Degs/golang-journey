package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Joe-Degs/golang_journey_au_Joe/basixs/concurrency_patterns/runner_pattern/runner"
)

func main() {
	run := runner.New(time.Second * 2)
	run.Add(first, second, third)
	err := run.Start()
	if err != nil {
		log.Println(err)
	}
}

func first(int) {
	time.Sleep(time.Millisecond * 50)
	fmt.Println("First func is done!")
}

func second(int) {
	time.Sleep(time.Second)
	fmt.Println("second func done after sleeping for a second LOL!")
}

func third(int) {
	time.Sleep(time.Second)
	fmt.Println("third func done.")
}
