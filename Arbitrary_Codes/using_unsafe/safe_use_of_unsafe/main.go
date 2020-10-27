package main

import (
	"fmt"
	"unsafe"
)

// COMPILE TIME VERIFICATION WITH THE  GO VET TOOL

func main() {
	// array of uint32 values stored in contigous memory
	arr := [...]uint32{1, 2, 3}

	// the size of a uint32 value in bytes
	size := unsafe.Sizeof(uint32(0))
	fmt.Printf("\nSIZE OF UINT32 VALUE  IS %xbytes\n\n", size)

	// take the initial memory address of the slice and begin iteration
	// always remember becuase a slice is just a header and pointer to an
	// underyling array, always take the pointer of the first idx's value.
	/*
				ptr := uintptr(unsafe.Pointer(&arr[0]))
				for i := 0; i < len(arr); i++ {
					// print out the value at the the current address we took
					// then update the pointer to point at the next address in memory
					fmt.Printf("ptr: %x, value: %d\n", ptr, (*(*uint32)(unsafe.Pointer(ptr))))
					ptr += size

		         // when we run this we get the following result;
		         // ptr: c00006af1c, value: 1
		         // ptr: c00006af20, value: 2
		         // ptr: c00006af24, value: 3

					// Now this move we just made is illegal but it feels good. I Know!(in Monica Geller's voice).
					// we are reading array elements directly from memory and not using the index.
					// always remember slices are just describing a specific array in memory, to get the address
					// of that array you need to get the address of the first element in the array then use
					// that address to determine the end address of that array in memory. COOL stuff! LOL!
				}
	*/
	/*
	   Even though this program runs as expected, we get a go vet warning
	   "|24 col 55 warning| possible misuse of unsafe.Pointer"
	   Reason being that uintptrs are just integers big enough to hold memory addresses.
	   Why we would use this is to print the memory address of a value. The go syscall package
	   uses it a bunch because system calls require access to certain memory addresses and there
	   are standard ways of doing it right.
	   What we did is a security flaw and can be used to break your system or steal vital info.
	   When you convert a pointer to uintptr, the garbage collector does not update the uintptr
	   when the value is moved from that address it first stored at. Meaning if the object the
	   uintptr is moved, uintptr now contains the memory address of nothing or something else
	   other that what was once there which is dangerous. What if a user's credential is what is
	   now stored in memory? Bro you don fuck up be that..

	   There is a better way to do this and not get judged by the go vet tool.
	   And that is comming right up. That is instead of storing the pointer in a uintpr variable
	   before starting advancing through the array, we could just take the address on the fly and
	   access the value with it without storing it first.
	   Lets take a look:
	*/
	for i := 0; i < len(arr); i++ {
		// Now we can print get the value at an index by;
		//    - taking the address of the first element in the array
		//    - applying the offset of (i * size) bytes to advance into the memory addcress of the value
		//    - convert uintptr back to *uint32 and dereference it to print the value.
		fmt.Printf("ptr: %x, value: %d\n", uintptr(unsafe.Pointer(&arr[i])), *(*uint32)(unsafe.Pointer(
			uintptr(unsafe.Pointer(&arr[0])) + (uintptr(i) * size),
		)))

		// we run this to get the same result as the previous one but no warning from govet :).
		// ptr: c00006af1c, value: 1
		// ptr: c00006af20, value: 2
		// ptr: c00006af24, value: 3
	}
}
