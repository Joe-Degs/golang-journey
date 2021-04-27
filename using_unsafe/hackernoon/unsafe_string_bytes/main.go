package main

import (
	"reflect"
	"unsafe"
)

func SafeBytestoString(b []byte) string {
	return string(b)
}

func SafeStringtoBytes(s string) []byte {
	return []byte(s)
}

func UnsafeStringToBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
	}))
}

func UnsafeBytesToString(b []byte) string {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	return *(*string)(unsafe.Pointer(&reflect.StringHeader{
		Data: sh.Data,
		Len:  sh.Len,
	}))
}

//func main() {
//	ts := "random string"
//
//	s := SafeStringtoBytes(ts)
//	us := UnsafeStringToBytes(ts)
//	fmt.Println(s)
//	fmt.Println(us)
//}
