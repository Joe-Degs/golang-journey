package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name, email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	u := &user{name: "Joe Attah", email: "joeattah@pmail.com"}
	u.notify()
	sendNotification(u)
}

/*
   interfaces do a lot of things which i dont quite understand yet and still trying to understand.
   how useful it is to me is still a mystery i am trying to unravel, what the future holds i know not of, so all i do is learn, learn and hope for a less dim future than this one.

   Apparently, the internals of interfaces are weird and not cool at all. internally, it is just a two-word data structure that has one word pointing to an internal itable that contains type information about the stored value and its method set, the other one contains a pointer to the stored value.

*/
