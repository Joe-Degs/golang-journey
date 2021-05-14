package main

import (
	"fmt"
	"io"
	"strings"
	"unicode"
	"unicode/utf8"
)

type angry_reader struct {
	r io.Reader
}

func new_reader(r io.Reader) *angry_reader {
	return &angry_reader{r: r}
}

func (a *angry_reader) Read(b []byte) (n int, err error) {
	n, err = a.r.Read(b)
	for r, i, w := rune(0), 0, 0; i < n; i += w {
		r, w = utf8.DecodeRune(b[i:])
		if !unicode.IsLetter(r) {
			continue
		}
		ru := unicode.ToUpper(r)
		if wu := utf8.EncodeRune(b[i:], ru); w != wu {
			return n, fmt.Errorf("%c->%c, size mismatch %d->%d", r, ru, w, wu)
		}
	}
	return
}

func a_fucking_loop_example(n int) {
	for i := 0; i < n; i++ {
		println(i)
	}
}

func main() {
	some_random_fucking_turkish_string := "üşîöğİıâçğûüşîöİığççççççâğİİ—İı–ççû ûüğâçıİİ"

	//make a new reader from the string
	reader := strings.NewReader(some_random_fucking_turkish_string)

	// a new angry_reader from our prev reader
	angry := new_reader(reader)
	buf := make([]byte, len(some_random_fucking_turkish_string))
	_, err := angry.Read(buf)
	if err != nil {
		fmt.Errorf("%v %s\n", err, "fuck off!")
	}
	a_fucking_loop_example(10)
	fmt.Printf("%s\n", string(buf))
}
