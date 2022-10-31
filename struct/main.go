package main

import "fmt"

type student struct {
	id      int
	name    string
	age     int
	city    string
	address string
	email   string
}

func details_student(s student) {
	fmt.Println("Details of student:", s.name)
	fmt.Printf("id = %v\nname = %v\nage = %v\ncity =  %v\naddress = %v\nemail = %v", s.id, s.name, s.age, s.city, s.address, s.email)
	fmt.Println()
}

func main() {

	//var student_list []string
	s1 := student{1, "mehedi", 30, "Feni", "Dhaka", "m@gmail.com"}
	s2 := student{2, "Hossain", 31, "Feni", "Dhaka", "h@gmail.com"}

	fmt.Println(s1.name)
	fmt.Println(s2.name)
	fmt.Println(s1, s2)
	fmt.Println()

	details_student(s1)

	details_student(s2)
}
