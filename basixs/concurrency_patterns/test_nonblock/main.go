package main

import (
	"fmt"
	"strings"
)

func main() {
	pint := make(chan string)

	go func() {
		select {
		case str := <-pint:
			if strings.Contains(str, "pint") {
				fmt.Println("Fancy a pint Love?")
				break
			}
		default:
			fmt.Println("No pint for you boujee boy.")
		}
	}()
	/*
	   type Os struct {}

	   func (*Os) Exit(int) {
	      fmt.Println("lol no")
	   }

	   var os Os
	   os.Exit(100)
	*/
}
