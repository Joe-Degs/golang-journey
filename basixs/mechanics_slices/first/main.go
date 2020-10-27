package main

import "fmt"

/*
   // trying to use this idea but it doesnt seem like a nice idea.
   // type assertion would be tedious with this approach.
   type buff []interface{}

   type RingBuffer struct {
   	buf  *buff
   	size int
   }

*/

var buffer [256]byte

func AddOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i]++
	}
}

func subtractOneFromSliceLen(slice []byte) []byte {
	slice = slice[0 : len(slice)-1]
	return slice
}

func main() {
	slice := buffer[10:20]
	/*
		for i := 0; i < len(slice); i++ {
			slice[i] = byte(i)
		}
		fmt.Println("before", slice)
		AddOneToEachElement(slice)
		fmt.Println("After", slice)
	*/

	fmt.Println("Before: len(slice) =", len(slice))
	newSlice := subtractOneFromSliceLen(slice)
	fmt.Println("After: len(slice) = ", len(slice))
	fmt.Println("After: len(newSlice) =", len(newSlice))
}

/*
QUE
------
if you have a buffer; `var buffer [150]byte` and it is sliced; `slice := buffer[1-len(buffer)-1]`
what will the underlying slice header look like?

ANS
-------
slice := sliceHeader{
   Length:        148
   ZerothElement: &buffer[1]
}

*/
