package main

import "fmt"

func main() {
	fmt.Println("welcome to loop")

	// print 0 to 50 valuie using for loop
	x := 50
	for i := 0; i < x; i++ {
		fmt.Println(i)
	}

	//  for loop using continue statement
	value := 20
	for i := 0; i < value; i++ {
		if i%5 == 0 {
			fmt.Println("skip value is ", i)
			continue

		}
		fmt.Println(i)

	}

	// for loop using brek statement
	val := 20
	for i := 0; i < val; i++ {
		if i == 10 {
			fmt.Println("Execution terminate 10 position ")
			break
		}
		fmt.Println(i)

	}

	// odd and even value find and calculate sum
	var even []int
	var odd []int
	num := 50
	odd_sum := 0
	even_sum := 0

	for i := 0; i < num; i++ {
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
