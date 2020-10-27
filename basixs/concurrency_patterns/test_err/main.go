package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("random error")
	fmt.Println(shootError(&err))
}

func shootError(err *error) error {
	return *err
}
