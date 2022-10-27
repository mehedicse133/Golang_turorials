package main

import "fmt"

func main() {
	fmt.Println("This is arry module")

	var animals [3]string

	animals[0] = "Dog"
	animals[1] = "Cat"
	animals[2] = "Tiger"

	fmt.Println("The animals list are : ", animals)

	var student = [5]string{"Mehedi", "Sharif", "jahir"}
	fmt.Println("The student list are :", student)
	fmt.Println("The lenth of student array is :", len(student))
	fmt.Println("The capacity  of student  arry is:", cap(student))
}
