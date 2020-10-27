package main

type token string

func (t *token) String() {
	println(*t)
}

func main() {
	newtoken := token("String")
	newtoken.String()
}
