package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Fuck you!")
		return
	}
	path, err := filepath.Abs(os.Args[1])
	checkErr(err)

	b, err := os.ReadFile(path)
	checkErr(err)

	fmt.Println(string(b))
}

func checkErr(err error) {
	if err == nil {
		return
	}
	fmt.Fprintln(os.Stderr, "fuck you again!")
	return
}

// always remember to close files you have opened.
// because golang performs io operations when the internal
// buffer is full. Closing the file flushes the buffer before
// closing. soo theres that. Though i'll get it off my chest.
