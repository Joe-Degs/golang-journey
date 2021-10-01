// This package uses an example to show how to broadcast the output of one
// process to many processes using `io.MultiWriter`. Useful for many things
// and a handy snippet to keep in hand in times of trouble
package main

import (
	"bytes"
	"io"
	"log"
	"os/exec"
)

var (
	words   = []string{"Game", "Feast", "Dragons", "of"}
	cmds    = make([][2]*exec.Cmd, len(words))
	writers = make([]io.Writer, len(words))
	buffers = make([]bytes.Buffer, len(words))
	err     error
)

func main() {
	for i := range words {
		// create grep command for every word in `words`
		cmds[i][0] = exec.Command("grep", words[i])
		if writers[i], err = cmds[i][0].StdinPipe(); err != nil {
			log.Fatal("in pipe", i, err)
		}

		// create `wc` to count the number of lines from its stdin pipe
		cmds[i][1] = exec.Command("wc", "-l")

		// pipe `wc` process stdin to `grep` process stdout
		if cmds[i][1].Stdin, err = cmds[i][0].StdoutPipe(); err != nil {
			log.Fatal("in pipe", i, err)
		}

		// redirect stdout of the `wc` process to a write buffer
		cmds[i][1].Stdout = &buffers[i]
	}

	// create a `cat` process
	cat := exec.Command("cat", "book_list.txt")
	// redirect stdout of the cat process to `writers`
	cat.Stdout = io.MultiWriter(writers...)

	// run the cat process and wait for it to finish
	if err := cat.Run(); err != nil {
		log.Fatal("cat run error", err)
	}

	// close the writing pipes, as the cat process is finished and
	// writing to the grep process's stdin is done
	for i := range cmds {
		if err := writers[i].(io.Closer).Close(); err != nil {
			log.Fatalln("close 0", i, err)
		}

		// start the `grep` and `wc` process asynchronously and wait for them
		// to finish executing
		for _, p := range cmds[i] {
			if err := p.Start(); err != nil {
				log.Fatalf("%s start %v\n", p.Args[0], err)
			}
		}
	}

	// now we wait for the wc process to finish executing and then we
	// spill the process's stdout redirected to the buffers to the terminal's
	// stdout
	for i := range cmds {
		if err := cmds[i][1].Wait(); err != nil {
			log.Fatalln("wc wait", i, err)
		}
		count := bytes.TrimSpace(buffers[i].Bytes())
		log.Printf("%10q %s entries", cmds[i][0].Args[1], count)
	}
}
