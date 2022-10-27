package main

import "fmt"

func main() {
	fmt.Println("This is arry module")

	// staic array
	var animals [3]string

	animals[0] = "Dog"
	animals[1] = "Cat"
	animals[2] = "Tiger"

	fmt.Println("The animals list are : ", animals)

	var student = [5]string{"Mehedi", "Sharif", "jahir"}
	// array print using index
	fmt.Println("The student list are :", student)
	fmt.Println("The lenth of student array is :", len(student))
	fmt.Println("The capacity  of student  arry is:", cap(student))

	// arry print using loop
	var a = [5]string{"a", "b", "c"}
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// Dynamic arry and slice
	var num []int
	num = append(num, 2)
	num = append(num, 3)
	num = append(num, 4)
	num = append(num, 5)
	num = append(num, 6)
	fmt.Println(num)
	fmt.Println("Length of num arrry is:", len(num))

}
