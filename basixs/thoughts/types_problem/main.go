package main

import "fmt"

func check(v interface{}) {
	num := 2
	switch v.(type) {
	case fmt.Sprintf("%T", num):
		fmt.Println("works!")
		break
	default:
		fmt.Println("doesnt!")
	}
}

func main() {
	check(7)
}
