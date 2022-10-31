package main

import "fmt"

func main() {
	fmt.Println("welcome to functions")

	sum := sum(2, 3)
	sub := sub(2, 3)
	div := div(45, 2)
	variodic := variodic(2, 3, 3, 4)

	fmt.Println(sum)
	fmt.Println(sub)
	fmt.Println(div)
	fmt.Println(variodic)
}

func sum(a int, b int) int {
	r := a + b
	return r
}

func variodic(values ...int) int {
	r := 0
	for _, value := range values {
		r += value
	}
	return r
}

func sub(a int, b int) int {
	r := a - b
	return r
}

func div(a float64, b float64) float64 {
	if a == 0 {
		fmt.Println("Enta a positive mumber of a : ")
	}
	if b < 0 {
		fmt.Println("Enter the value of b grater than Zero")
	}

	r := a / b
	return r
}
