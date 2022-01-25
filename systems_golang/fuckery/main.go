// fuckery is a binary for testing out weird things, its like my go playground
// that compiles to the binary fuckery(which is a command line tool)
// write code, compile and you've got fuckery to use
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	text := make([]string, 0, 10)
	for reader.Scan() {
		text = append(text, reader.Text())
	}
	if err := reader.Err(); err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("scanner terminated due to error: %v", err))
	}
	if len(text) >= 1 {
		fmt.Fprintln(os.Stdout, "\nYou entered:")
		for d, t := range text {
			fmt.Fprintf(os.Stdout, "%d) %s\n", d+1, t)
		}
	}
}
