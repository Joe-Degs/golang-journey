package main

import (
	"fmt"
	"strings"
)

// never copy a non zero string builder
// never ever do that shit.
// always use standard implementation of things they are efficient.

func main() {
	b := strings.Builder{}
	b.WriteString("hey")
	b.WriteRune('!')
	b.WriteByte(0xa)
	s := b.String()
	fmt.Println(s)
	panicker(s)
}

func panicker(str string) {
	b := strings.Builder{} // nil builder
	b.WriteString(str)
	b.WriteString("Joe")
	s := b
	s.WriteString("Something")
	fmt.Println(s.String())
}
