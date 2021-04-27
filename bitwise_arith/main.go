package main

import "fmt"

func bitwise_arith() {
	ans := 1 << 8
	abs := 4 >> 3

	fmt.Printf("%#b is %d :: hex is %#[1]x and %#[2]x\n", ans, ans)
	fmt.Printf("%#b is %d\n", abs, abs)
}

func ringbuf() {
	a := 5 % (5 - 3)
	b := 5 % (5 - 2)
	c := 5 % (5 - 1)
	fmt.Println(a, b, c)
}

// this returns the number occuring the odd number of times.
func odd() {
	var res int
	arr := []int{1, 1, 1, 1, 2, 2, 2}
	for _, el := range arr {
		res ^= el
	}
	fmt.Printf("odd occuring numer is %d", res)
}

func main() {
	odd()
}
