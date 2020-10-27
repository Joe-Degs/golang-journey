package main

import (
	"bytes"
	"fmt"
)

type path []byte

func (p *path) TruncateAtFinalSlash() {
	i := bytes.LastIndex(*p, []byte("/"))
	if i >= 0 {
		*p = (*p)[0:i]
	}
}

func (p path) ToUpper() {
	for i, b := range p {
		if 'a' <= b && b <= 'z' {
			p[i] = b + 'A' - 'a'
		}
	}
}

func main() {
	pathName := path("usr/bin/tso") // convert from string to path
	//pathName.TruncateAtFinalSlash()
	pathName.ToUpper()
	fmt.Printf("%s\n", pathName)
}
