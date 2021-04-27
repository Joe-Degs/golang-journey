package main

import (
	"fmt"
	"unsafe"
)

type Bytes400 struct {
	val [100]int32
}

type TestStruct struct {
	a [9]int64
	b byte
	c *Bytes400
	d int64
}

// unsafe array iteration and type conversion on the fly with zero allocations
// get info about structs in memory
// change struct fields directly in memory, with struct pointer and field offset.

func main() {
	array := [10]uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var sum int8

	// unsafe array iteration
	sizeOfUint64 := unsafe.Sizeof(array[0])
	for i := uintptr(0); i < 10; i++ {
		sum += *(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(&array)) + sizeOfUint64*i))
	}

	fmt.Println("Sum = ", sum)

	// size of structs and offset of structs.
	// the offset of the struct field plus the size of the field's type
	// will give the full size of struct in memory. which is unsafe.
	t := TestStruct{b: 42}
	fmt.Println("Size of b ->", unsafe.Sizeof(t))
	fmt.Println(unsafe.Offsetof(t.a), unsafe.Offsetof(t.b), unsafe.Offsetof(t.c), unsafe.Offsetof(t.d))

	fmt.Println(unsafe.Sizeof(Bytes400{}))

	fmt.Println("old value of t.b -> ", t.b)

	// change the value of struct field t.b
	*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&t)) + unsafe.Offsetof(t.b)))++

	// print out value of t.b
	fmt.Println("new value of t.b -> ", t.b)
}
