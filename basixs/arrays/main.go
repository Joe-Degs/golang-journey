package main

import "fmt"

func main() {

	// array internals
	array := [5]*int{0: new(int), 1: new(int)}
	*array[0] = 5
	*array[1] = 10

	for i := 0; i < len(array); i++ {
		if array[i] != nil {
			fmt.Println(*array[i])
		} else {
			fmt.Println(array[i])
		}
	}

	// multidimensional arrays
	ThreedArray := [3][3][3]int{{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}, {{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}, {{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}}
	fmt.Println("3dArray: ", ThreedArray)

	// slice internals
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("slice: ", slice)
	newSlice := slice[1:3]
	fmt.Println("newSlice: ", newSlice)
	fmt.Println("cap(newSlice) =>", cap(newSlice))

	// fucking go doesnt have steps when slicing slices. so i guess if you wanna do that you devise your own algorithm for doing it.
	// but they have some wierd looking structure that looks just like stepping in python but it does some other job related to
	// the cap() of the slice
	// でㄩぽれ => discovered that pressing <ctrl + k> with and two other random characters on the keyboard generate a utf8 character.
	source := []string{"Apple", "Orange", "Pear", "Plum", "Banana"}
	nuSlice := source[2:3:4] // limiting the capacity of the new slice to protect the underlying data of the original slice
	fmt.Println(nuSlice)

	// iterating over slices with range and for
	someSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 8, 9}
	for idx, val := range someSlice {
		fmt.Printf("Index: %d, Value: %d\n", idx, val)
	}

	// devil in the details: range copies the value and not references to them.
	// thats pretty memory expensive or something, systems engineers will be very concerned with shit like this.
	ptrSlice := []int{1, 2, 3, 4, 5}
	for idx, val := range ptrSlice {
		fmt.Printf("Value: %d,  Value-Addr: %X    Elem-Addr: %X\n", val, &val, &ptrSlice[idx])
	}
}
