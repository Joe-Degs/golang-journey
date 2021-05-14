package main

// This is a filesystem search engine, and this is what the program does

// -> reads directory path and strings to search from the arguments
// -> gets a list of files in the selected path
// -> reads each files and passes the lines that contain the selected string
//    another writer
// -> this other reader will inject colors into the highlighted string and copy
//    to standard output

// Colored output on the unix terminal is obtained by the ff sequences
//
// -> \xbb1: an escape character
// -> [		 an opening bracket
// -> 39	 an number repping the color sort of a color code
// -> m		 the letter m to climax it.


type queryWriter struct {
	Query []byte
	io.Writer
}

func(q queryWriter) Write(b []byte) (n int, err error) {
	lines := bytes.Split(b []byte{'\n'})
	l := len(q.Query)

	for _, b := range lines {
		i := bytes.Index(b, q.Query)

		if i == -1 {
			continue
		}

		for _, s := range [][]byte{
			b[:i], // index just before match
			[]byte("\x[31m"), // color
			b[i : i+l], // match
			[]byte("\x[39m"), // default color
			b[i+l:], // whatever is left
		} {
			v, err := q.Writer.Write(s)
			n += v
			if err != nil {
				return 0, err
			}
		}

		fmt.Frprinln(q.Writer)
	}

	return len(b), nil
}

func main() {
	if len(os.Args) < 3 {
		checkErr("specify path and search string")
	}
}

func checkErr(msg string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, msg, err)
	}
}
