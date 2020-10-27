package main

import (
	"fmt"
	"unsafe"
)

// Run - time verification with the Go compilers `checkptr` debugging flag.

/*
   Looking at another example. Suppose we are passing a Go structure to a Linux kernel API
   which would typically accept a C union(i dont know what a c union is(quick trip to wikipedia)).
   One pattern for doing this is to have an overarching Go structure which contains a raw byte array
   (mimicking a C union), and then to create typed variants for possible argument combinations.
*/

// one is typed Go struct containing structured data to pass to kernel.

type one struct{ v uint64 }

// two mimics a C union type which passes a blob of data to the kernel.
type two struct{ b [32]byte }

/*
func main() {
	// we want to send the contents of a to the kernel as raw bytes.
	in := one{v: 0xff}
	out := (*two)(unsafe.Pointer(&in))

	// Assume the kernel will only access the first 8 bytes. But what is stored
	// in the remaining 24 bytes?
	fmt.Printf("%#v\n", out.b[0:8])
}
*/

/*
   We have a little prblem with the above code. If we cast a smaller structure into a big one, we
   we enable the reading of arbitrary beyond the end of the end of the smaller structures data. This
   is another way the careless use `unsafe` can introduce problems into the application.
*/

// to do this well, we have to initialize the union structure first before copying data into it.
// to ensure arbitrary memory is not accessed.

func newTwo(in one) *two {
	// initialize out and its array.
	var out two

	// explicitly copy the contents of in into out by casting both into byte arrays then
	// slicing the arrays. This will produce the correct union packed structure, without relying on
	// unsafe casting into a smaller type of a larger type
	copy(
		(*(*[unsafe.Sizeof(two{})]byte)(unsafe.Pointer(&out)))[:],
		(*(*[unsafe.Sizeof(one{})]byte)(unsafe.Pointer(&in)))[:],
	)

	return &out
}

func main() {
	// the two structures are now appropriately initialized.
	out := newTwo(one{v: 0xff})
	fmt.Printf("%#v\n", out.b)
}
