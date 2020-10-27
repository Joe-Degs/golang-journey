package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stdin, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(-1)
	}
	name := os.Args[1]

	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
		os.Exit(-1)
	}
	fmt.Printf("The address is %s\n", addr.String())
	os.Exit(0)
}