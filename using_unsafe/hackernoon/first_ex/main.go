package main

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

func main() {
	a := "Hello. Currrent time is " + time.Now().String()

	fmt.Println(a)

	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&a))
	stringSize := unsafe.Sizeof(stringHeader.Data)

	// add 5 bytes to the string header data pointer, and insert a '!' there.
	// stringHeader.Data is a uintptr and not a real pointer so its addable.
	*(*byte)(unsafe.Pointer(stringHeader.Data + 5)) = '!'

	// Okay, here i was trying find out if i could get the exact size of the string data
	// in bytes and then insert then insert '!' at the end of the string. Which after i tried,
	// I realized that it was a stupid thing to think about in the first place.
	// And it seems the standard memory space allocated for small strings is around 8 bytes.
	// Using the len is better(which i am aware of) but i wanted to know if i could do it with
	// the size of the string in memory.
	*(*byte)(unsafe.Pointer(stringHeader.Data + stringSize)) = '!'

	fmt.Println(a)

	// looks like the size of a string header is 8 bytes standard or something.
	// but maybe if the size exceeds the 8 bytes they might allocate more or something.
	// i dont really know what the fuck is happening with go memory allocation and stuff.
	fmt.Println("string size is", stringSize)
	fuckingString := "fuck you and me and you and me"

	// I tried to do some small experimentation to find out some more things about runtime memory
	// allocation and all that stuff.
	// So i got the size of the string header data pointer in a confusingly difficult way and then
	// printed it out. Then i didnt even know what to do again.
	fmt.Println(fuckingString+"'s size is", unsafe.Sizeof((*reflect.StringHeader)(unsafe.Pointer(&fuckingString)).Data), "bytes")

	// cast the bits of fucking string to a byte array
	// Yeah, this is where i tried to flex some of my unsafe go muscles.
	// i just copied all the bytes of the string into a byte slice in the most
	// unsafe way possible.
	// and then i print the bytes out to see what the fuck i just did.
	// if i was curious enough, i would iterate over the byte array convert them to runes
	// or something and then see if i'll get the string back but not today bro not today.
	byteArr := (*[]byte)(unsafe.Pointer(&fuckingString))
	fmt.Printf("fuckingString to byte array = %#v\n", byteArr)

}
