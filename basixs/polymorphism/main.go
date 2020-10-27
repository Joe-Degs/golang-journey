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

type admin struct {
	name, email string
}

func (a *admin) notify() {
	fmt.Printf("Sending admin to %s<%s>\n", a.name, a.email)
}

func main() {
	bill := &user{"Bill", "userbill@rockmail.com"}
	lisa := &admin{"Lisa", "adminlisa@punkmail.com"}
	sendNotification(bill)
	sendNotification(lisa)
}

func sendNotification(n notifier) {
	n.notify()
}
