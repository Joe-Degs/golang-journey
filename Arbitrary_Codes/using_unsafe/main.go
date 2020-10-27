package main

import (
	"fmt"
	"os"
	"unsafe"

	"golang.org/x/sys/unix"
)

// UNDERSTANDING TYPE SAFETY AND USING UNSAFE ....

/*
   Got to reading the use of unsafe package in go. Always wanted to know what those things
   were good for. Everybody says "the unsafe package helps to step arounds go's type safety
   (wtf does this mean?)".
   The name of the package even tells you its unsafe to use, but what tastes better than the
   forbidden fruit? ...Nothing LOL!.

   Lets take a look at what type safety is in golang. Nahh what is type safety in general?
   Lets take a quick trip to wikipedia and check it out. According to wikipedia
   "According to wikipedia, type safety is the extent to which a language discourages or prevents
   type errors." Type errors arise performing operations on types you are not supposed to perform
   such operations on. Example "2 / 'Hello'" results in a big fat type error. Because it is illegal
   to perform division on a integer and a string. Atleast in Go. Maybe javascript? LOL.

   Go is type safe in that it doesnt allow you to mess up with types in ways that will make the
   program crush. Also in lowlevel sense you cant do anything you like with pointers. In some langs
   you can do pretty much anything with pointers, like accessing arbitrary memory through incrementing
   or decrementing the pointer or reading or data of another type that will fit that memory. Go
   frowns upon this but provides a package for unsafe pointer usage like the one described above.
   (Its like the way girls say stop it but expect you not to stop, you know LOL!)
   Be careful when you use unsafe, it could result in very unpredictable behaviour of your programs.

   Lets take some codxamples
   {
      var i int8 = -1 // binary rep is: 1111111
      var j int16 = int16(i) // binary rep is: 111111 111111
      println(i, j) // both result in: -1, -1

      // this is a type casting we did right here.
      // We moved our value into a bigger storage space just the way go expects sane people to.
      // since go provides unsafe there a crazy ways to do this and get away with it like

      var k int16 = *(*int16)(unsafe.Pointer(&i)) or
      var l uint8 = *(*uint8)(unsafe.Pointer(&i)) // this will definitely be 8 bits all set to one
      // k in binary should be 11111111 00000000 or 00000000 11111111? but its not(i'm not good at the binary thing).
      // this particular block of code needs to be broken down for me to understand it well so here we go
      {
         // get the pointer of i and change it to an unsafe pointer.
         unsafePtr := unsafe.Pointer(&i)

         // use type assertion to cast the value of unsafePtr to *int16
         int16Ptr := (*int16)(unsafePtr) // this is the syntax for golang type assertion
         // we do (*int16) because the type we are asserting is a pointer variable.

         // we dereference the the *int16 to get int16 value we can use.
         k := *int16Ptr
      }

      // the variable k should be composed of 8 bits set to one and the rest set to 0. but trust me when i tell you
      the value of k is changing with time LOL!. I'm even scared. Maybe it should be 0000000 1111111? probably this.
      But i'm sure i saw it change because when i tested in the morning it was something like 22739 or so, then it
      changed to some value looking like -7928, now its 255 which is believable.
   }

   Let's continue with learning the why we would ever need to use unsafe.
*/

/*
func main() {
	var i int8 = -1
	var j int16 = int16(i)
	println("i and j is ", i, j)

	var k int16 = *(*int16)(unsafe.Pointer(&i))
	println("l is ", k)

	var l uint16 = *(*uint16)(unsafe.Pointer(&i))
	fmt.Println("k is ", l)
}
*/

// USING UNSAFE FOR TYPE CONVERSIONS AND LEARNING HOW TO DO IT RIGHT.
/*

type A struct {
	A int8
	B string
	C float32
}

type B struct {
	D int8
	E string
	F float32
}

*/

/*
   Trying to copy data of a particular type to another type that has same underlying structure
   and size and data. The golang lang type safety will not allow you to do this. Because it is
   very risky. but thank god for unsafe you can do this without breaking a sweat.

   First if you just try to cast the type with the normal type assertion you get an error
   example:
   a := A{1, "foo", 12.33}
   b := (*B)(&a)
   // this will give an error: "cannot convert a (variable of type A) to *B"

   we can get do it without the compiler bitching about it with the unsafe
*/

/*
func main() {
	a := A{1, "foo", 1.33}
	println(a.A, a.B, a.C)
	b := (*B)(unsafe.Pointer(&a))
	println(b.D, b.E, b.F)
}
*/

//                   // A REALWORLD EXAMPLE :>) (unsafe pointer and system calls)

func Ioctl(fd uintptr, request int, argp unsafe.Pointer) error {
	_, _, errno := unix.Syscall(unix.SYS_IOCTL, fd, uintptr(request), uintptr(argp))
	if errno != 0 {
		return os.NewSyscallError("ioctl", fmt.Errorf("%d", int(errno)))
	}
	return nil
}

func main() {
	f, err := os.Open("/dev/vhost-vsock")
	if err != nil {
		fmt.Printf("0 and err is %v\n", err)
		return
	}
	defer f.Close()

	//Context Id to be populated by Ioctl.
	var cid uint32

	// retrieve the context Id of this machine from /dev/vhost-vsock.
	err = Ioctl(f.Fd(), unix.IOCTL_VM_SOCKETS_GET_LOCAL_CID, unsafe.Pointer(&cid))
	if err != nil {
		fmt.Printf("0 and err is %v\n", err)
		return
	}

	fmt.Printf("ContextId is: %d", cid)
}
