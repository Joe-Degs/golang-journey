package main

import (
	"fmt"
	"os"
	"path/filepath"
	"io/fs"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "Please specify a path")
		return
	}
	root, err := filepath.Abs(os.Args[1])
	checkErr(err)
	fmt.Println("Listing files in root", root)

	var c struct {
		file int
		dir  int
	}

	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		// Walk through and count
		if info.IsDir() {
			c.dir++
		} else {
			c.file++
		}
		fmt.Println("-", path)
		return nil
	})
	fmt.Printf("Total: %d file and %d directories", c.file, c.dir)
}

func checkErr(err error) {
	if err == nil {
		return
	}
	fmt.Fprint(os.Stderr, err)
	os.Exit(-1)
}
