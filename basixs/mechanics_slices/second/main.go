package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

type AnonyType struct {
	random string
}

func GetType(t interface{}) string {
	s := fmt.Sprintf("%T", t)
	if !strings.Contains(s, "main") {
		return s
	}
	return s[strings.LastIndex(s, ".")+1:]
	//return fmt.Sprintf("%T", t)[strings.LastIndex(fmt.Sprintf("%T", t), ".")+1:]
}

func main() {
	n := 2
	var file os.File
	var buf bytes.Buffer
	rand := AnonyType{random: "Random Type"}
	fmt.Println(GetType(rand))
	fmt.Println(GetType(n))
	fmt.Println(GetType(file))
	fmt.Println(GetType(buf))
	fmt.Println(GetType('i'))
}
