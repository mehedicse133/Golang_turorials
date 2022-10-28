package main

import (
	"fmt"
)

// main function
func main() {

	fmt.Println("Enter your age: ")
	var age int
	fmt.Scan(&age)
	fmt.Println("Your age is :", age)
	agecheck(age)

	var firstname, lastname string
	fmt.Println("Enter first name : ")
	fmt.Scan(&firstname)

	fmt.Println("Enter your lastname: ")
	fmt.Scan(&lastname)
	y := len(lastname)

	fmt.Println(y)

	// var name []string
	// b := firstname + " " + lastname
	// v := strings.Split(b, "e")
	// for i := 0; i > len(v); i++ {
	// 	name = append(name, name[i])
	// }
	fmt.Println(b)
	fmt.Println(len(b))
	// Addition of two string
	fmt.Println(firstname + " " + lastname)

}

func agecheck(age int) {
	if age > 18 {
		fmt.Println("Your are adult")
	} else {
		fmt.Println("Yor are teen age")
	}
}
