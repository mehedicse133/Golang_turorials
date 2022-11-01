package main

import "fmt"

//https://www.golinuxcloud.com/go-cast/
func main() {
	fmt.Println("welcome to functions")

	sum := sum(2, 50)
	sub := sub(8, 3)
	div := div(0, 0)

	//fmt.Println(sum)
	print(sum)
	print(sub)
	// fmt.Println(sub)
	fmt.Println(div)
	print_line()

	variodic := variodic(2, 3, 3, 4)
	fmt.Println(variodic)
	print_line()

	num_ := -8
	fmt.Println(square(num_))
	print_line()

	// Pass by value and multi values return
	name, phone := multi_value_return("Mehedi", 222222)
	fmt.Println(name, phone)
	print_line()

	// Ignore returnde values form fumctions using _
	name, _ = multi_value_return("amam", 8888888)
	fmt.Println(name)
	print_line()

	// Ignore returnde values form fumctions using _
	_, phone = multi_value_return("amam", 8888888)
	fmt.Println(phone)
	print_line()

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
	} else if b == 0 {
		fmt.Println("Enter the value of b grater than Zero")

	}
	r := a / b
	return r

}

func square(num int) int {
	if num == 0 {
		return 1
	}
	if num < 0 {
		fmt.Println("Enter positvie value")
		return -1
	}
	return num * num
}

func multi_value_return(name string, phone int) (string, int) {
	return name, phone
}

func print_line() {
	fmt.Println("--------------------------------")
}

func print(values ...int) {
	for _, val := range values {
		fmt.Println(val)
	}

}
