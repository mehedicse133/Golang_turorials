package main

import "fmt"

// interface contains colletions of functions
type persondetails interface {
	getname() string
	getage() int
}

// struct contains collections of fields
type teacher struct {
	name string
	age  int
}

type stduent struct {
	name string
	age  int
}

// create teacher and student functions
func (s stduent) getname() string {
	return s.name
}

func (s stduent) getage() int {
	return s.age
}

func (t teacher) getname() string {
	return t.name
}

func (t teacher) getage() int {
	return t.age
}

func pirnt(p persondetails) {
	fmt.Println(p.getname())
	fmt.Println(p.getage())

}

func main() {
	fmt.Println("welcome to interface")

	// create object of student and teacher
	stu := stduent{name: "mehedi", age: 30}
	tea := teacher{name: "amam", age: 30}

	// pass objects into interface method
	pirnt(stu)
	pirnt(tea)

}
