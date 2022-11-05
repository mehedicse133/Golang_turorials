package main

import "fmt"

func main() {
	//defer
	fmt.Println("welcome to defer_panic_recover")
	defer fmt.Println("end the execution")
	fmt.Println("welcome to defer")

	defer fmt.Println("A")
	defer fmt.Println("B")
	defer fmt.Println("C")

}
