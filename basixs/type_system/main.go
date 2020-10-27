package main

import (
	"fmt"
	"strconv"
)

// type embedding is kinda confusing if you introduce pointers into the picture
// so here i am trying to understand how the *os.File does this seamlessly by emulating their techniques.
type File struct {
	*file
}

type file struct {
	age  int
	name string
}

func (f *file) changeCredentials(field, value string) {
	fields := [2]string{"name", "age"}
	for idx, _ := range fields {
		if fields[idx] == field && idx == 0 {
			f.name = value
		} else {
			f.age, _ = strconv.Atoi(value)
		}
	}
}

func newFile(name string, age int) *File {
	return &File{&file{
		name: name,
		age:  age,
	}}
}

func (f File) Print() {
	fmt.Printf("Name: %v Age: %d\n", f.file.name, f.file.age)
}

//func interleg(f interface{}) {
//	switch v := f.(type); v {
//	default:
//		fmt.Println(v)
//	}
//}

func main() {
	newFile := newFile("Joe", 20)
	(*newFile).Print() // this is totally unnecessary (go does this implictly), this should have been => newFile.Print()
	newFile.file.changeCredentials("name", "Ama")
	newFile.file.changeCredentials("age", "50")
	newFile.Print()
}
