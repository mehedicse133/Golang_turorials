package main

import "fmt"

type stduent struct {
	name string
	age  int
}

type persondetails interface {
	getname() string
	getage() int
}

type teacher struct {
	name string
	age  int
}

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

	stu := stduent{name: "mehedi", age: 30}
	tea := teacher{name: "amam", age: 30}

	pirnt(stu)
	pirnt(tea)

}
