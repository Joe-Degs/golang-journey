package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"unicode"
	"unicode/utf8"
)

type ScrambleWriter struct {
	w io.Writer
	r *rand.Rand
	c float64
}

func NewScrambleWriter(w io.Writer, r *rand.Rand, c float64) *ScrambleWriter {
	return &ScrambleWriter{w, r, c}
}

func (s *ScrambleWriter) Write(b []byte) (n int, err error) {
	runes := make([]rune, 0, 10)
	for r, i, w := rune(0), 0, 0; i < len(b); i += w {
		//decode first rune return char and size
		r, w = utf8.DecodeRune(b[i:])
		if unicode.IsLetter(r) {
			runes = append(runes, r)
			continue
		}
		// scramble if unicode char is not a letter
		v, err := s.shambleWrite(runes, r)
		if err != nil {
			return n, err
		}
		n += v
		runes = runes[:0]
	}
	// after decoding all runes in b and len(runes) is not zero
	// scramble some runes!
	if len(runes) != 0 {
		v, err := s.shambleWrite(runes, 0)
		if err != nil {
			return n, err
		}
		n += v
	}
	return
}

func (s *ScrambleWriter) shambleWrite(runes []rune, sep rune) (n int, err error) {
	// scramble after first letter
	for i := 1; i < len(runes)-1; i++ {
		if s.r.Float64() > s.c {
			continue
		}
		j := s.r.Intn(len(runes)-1) + 1
		runes[i], runes[j] = runes[j], runes[i]
	}
	if sep != 0 {
		runes = append(runes, sep)
	}
	b := make([]byte, 10)
	for _, r := range runes {
		v, err := s.w.Write(b[:utf8.EncodeRune(b, r)])
		if err != nil {
			return n, err
		}
		n += v
	}
	return
}

func main() {
	buf := &bytes.Buffer{}
	newScrambler := NewScrambleWriter(buf, rand.New(rand.NewSource(334)), 2.0)
	newScrambler.Write([]byte("some random string shit like something lika that\n that that that\n"))
	fmt.Println(buf.String())
}
