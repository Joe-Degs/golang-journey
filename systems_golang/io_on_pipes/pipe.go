package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {
	pr, pw := io.Pipe()
	//closeError := errors.New("closing reader!")
	//defer pr.CloseWithError(closeError)

	var wg sync.WaitGroup

	wg.Add(1)
	go func(w io.WriteCloser) {
		defer wg.Done()
		data := []string{
			"The boy is going to school",
			"And he is not joking with nobody",
			"But people are joking with him thoo",
		}

		for _, line := range data {
			_, err := w.Write([]byte(line))
			if err != nil {
				fmt.Fprintf(os.Stderr, "io on pipes write error: %v\n", err)
				break
			}
		}
		w.Close()
	}(pw)

	wg.Add(1)
	go func(rc io.ReadCloser) {
		defer wg.Done()
		buf := make([]byte, 100)
		for {
			_, err := rc.Read(buf)
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Fprintf(os.Stderr, "io on pipes read error: %v\n", err)
				return
			}
			fmt.Fprintf(os.Stdout, "io on pipes: %s\n", string(buf))
		}
	}(pr)

	wg.Wait()
}
