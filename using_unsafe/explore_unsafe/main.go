package main

import (
	"fmt"
	"unsafe"
)

func to_2bytes(a uint16) [2]byte {
	// so basically this (*[2]byte) is a function.
	// which you pass the unsafe pointer to.
	return *(*[2]byte)(unsafe.Pointer(&a))
}

func read_arbitrary_mem() [4]byte {
	// if the type you are casting to is larger than the type you are casting from
	// then you stand the chance of reading arbitrary mem which is bad.
	a := uint16(1)
	b := (*[4]byte)(unsafe.Pointer(&a))
	return *b
}

func write_arbitrary_mem() {
	var v uint32

	// overwrite whatever is at this address.
	// this is extremely terrible thing to do, never try this
	// on anything possible.
	// starting to think i could write a device driver in golang.
	// with all this low_level stuff.
	*(*[4]byte)(
		unsafe.Pointer(uintptr(unsafe.Pointer(&v)) - 0xffffffff),
	) = [4]byte{0xff, 0xff, 0xff, 0xff}
}

func iterate_slice(a []uint32) {
	// this is the unsafest way to iterate over a slice ever!
	// but its fun to do and it looks like a safe thing to do, yet
	// very difficult.
	// i heard this is how array iteration is done in c(can you imagine!).
	for i := uintptr(0); i < uintptr(len(a)); i++ {
		fmt.Printf("%d\n", *(*uint32)(
			unsafe.Pointer(uintptr(unsafe.Pointer(&a[0])) + i*unsafe.Sizeof(a[0]))),
		)
	}
}

func main() {
	fmt.Println("uint16 to [2]byte", to_2bytes(0x2b4b))
	fmt.Println("casting to a larger type", read_arbitrary_mem())
	// output of this code is
	// uint16 to [2]byte [75 43]
	// casting to a larger type [1 0 75 43]

	//always do a size check before converting from types.

	iterate_slice([]uint32{0x2a1, 0xb2, 0xff3, 0x14, 0x95, 0x05})
	// write_arbitrary_mem()
}
