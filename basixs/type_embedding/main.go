package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name, email string
}

func (u *user) notify() {
	fmt.Printf("Sending email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	user  //embedded type
	level string
}

func main() {
	ad := admin{
		user{
			"Billy Punkii",
			"billylovesrockandroll@gmail.com",
		},
		"super",
	}

	// accessing the inner methods of ad
	ad.user.notify()

	// accessing the same methods on the outer struct even though its not declared because thats how cool go is :)
	ad.notify()
}

/*
   So basically what go is doing up above is graduating the methods of the inner struct to the other struct whilst at the same time maintaining the methods on the inner struct so they can be accessed from the inner struct and also overridden on the other struct. pretty smart right.. Its called inner type promotion
   yeah! :)
*/
