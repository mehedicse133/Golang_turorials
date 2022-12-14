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
	fmt.Printf("id = %v\nname = %v\nage = %v\ncity =  %v\naddress = %v\nemail = %v\n", s.id, s.name, s.age, s.city, s.address, s.email)
	fmt.Println("--------------------------------")
}

type student_list struct {
	students_list []student
}

func (slist *student_list) add_student(s student) {
	slist.students_list = append(slist.students_list, s)
	//	slist.students_list =
}

func (s *student) get_student_id() int {
	return s.id
}

func (s *student) edit_name(name string) {
	s.name = name
}

func (s *student) get_name() string {
	return s.name
}

func main() {

	//created object of student struct
	s1 := student{22, "mehedi", 30, "Feni", "Dhaka", "m@gmail.com"}
	s2 := student{2, "Hossain", 31, "Feni", "Dhaka", "h@gmail.com"}

	//print student objects
	fmt.Println(s1.name)
	fmt.Println("--------------------------------")
	fmt.Println(s2.name)
	fmt.Println("--------------------------------")

	// print both students objents
	fmt.Println(s1, s2)
	fmt.Println("--------------------------------")

	//print single student object using function
	details_student(s1)
	details_student(s2)

	// created object of student_list struct
	s := student_list{}
	s.add_student(s1)
	s.add_student(s2)
	// fmt.Println(s1)

	fmt.Println(s.students_list)
	fmt.Println("Length of slice students list : ", len(s.students_list))
	fmt.Println("--------------------------------")

	// use edit_name func and get stusdnet_id func
	s1.edit_name("Subbir")
	fmt.Println(s1)

	fmt.Println("--------------------------------")
	fmt.Println(s1.get_student_id())
	fmt.Println(s1.get_name())

}
