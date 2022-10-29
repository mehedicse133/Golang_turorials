package main

import "fmt"

func main() {
	fmt.Println("welcome to loop")

	x := 5
	for i := 0; i < x; i++ {
		fmt.Println(i)
	}

	var even []int
	var odd []int

	odd_sum := 0
	even_sum := 0

	for i := 0; i < x; i++ {
		if i%2 == 0 {
			even_sum += i
			even = append(even, i)
		} else {
			odd_sum += i
			odd = append(odd, i)
		}
	}

	fmt.Println(even)
	fmt.Println(even_sum)

	fmt.Println(odd)
	fmt.Println(odd_sum)

}
