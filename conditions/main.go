package main

import (
	"fmt"
	"strings"
)

// main function
func main() {

	else_if_conition("My namem is Mehedi")

	// fmt.Println("Enter your age: ")
	// var age int
	// fmt.Scan(&age)
	// fmt.Println("Your age is :", age)

	// // Pass age in variable into age_check func as argument
	// age_check(age)

	// var student []string
	// // var firstname, lastname string

	// fmt.Println("Enter how many student you take input:")
	// var num_student int
	// fmt.Scan(&num_student)

	// for i := 0; i < num_student; i++ {
	// 	var firstname, lastname string

	// 	fmt.Println("Enter first name : ")
	// 	fmt.Scan(&firstname)
	// 	fmt.Println(firstname)

	// 	fmt.Println("Enter last name : ")
	// 	fmt.Scan(&firstname)
	// 	full_name := firstname + " " + lastname
	// 	fmt.Println(full_name)
	// 	student = append(student, full_name)
	// }

	// fmt.Println(student)
	// fmt.Println(len(student))
	// fmt.Println(cap(student))

	// fmt.Println("Enter first name : ")
	// fmt.Scan(&firstname)
	// fmt.Println("The length of lastname", len(firstname))

	// fmt.Println("Enter your lastname: ")
	// fmt.Scan(&lastname)
	// fmt.Println("The length of lastname", len(lastname))

	// var name []string
	// b := firstname + " " + lastname
	// v := strings.Split(b, "e")
	// for i := 0; i > len(v); i++ {
	// 	name = append(name, name[i])
	// }
	// fmt.Println(b)
	// fmt.Println(len(b))
	// // Addition of two string
	// fmt.Println(firstname + " " + lastname)

}

func age_check(age int) {
	if age > 18 {
		fmt.Println("Your are adult")
	} else {
		fmt.Println("Yor are teen age")
	}
}

func else_if_conition(str string) {
	fmt.Println(len(str))

	s := strings.SplitAfter(str, "")

	var m_volume []string

	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == " " {
			continue
		}
		count += 1

		if s[i] == "m" {
			m_volume = append(m_volume, s[i])
		}

	}
	fmt.Println(count)
	fmt.Println(len(m_volume))
	fmt.Printf("the type ost s : %T", s)
}
