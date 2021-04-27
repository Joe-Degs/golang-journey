package main

import "syscall"

func main() {
	fd, err := syscall.Open("dev/stdout")
}
