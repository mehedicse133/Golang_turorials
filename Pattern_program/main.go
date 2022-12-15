package main

import (
	"fmt"
	"strconv"
)

func main() {

	var row int

	fmt.Print("Enter the rows : ")
	fmt.Scan(&row)
	value := 1

	for i := 0; i <= row; i++ {
		for j := 1; j <= i; j++ {
			if j == i {
				fmt.Printf(strconv.Itoa(value))
				value = value + 1
			} else {
				fmt.Printf(strconv.Itoa(value) + "*")
				value = value + 1
			}
		}
		fmt.Println()
	}
	s_value := value - 1
	for i := row; i >= 0; i-- {
		for j := 1; j <= i; j++ {
			if j == i {
				fmt.Printf(strconv.Itoa(s_value))
				s_value = s_value - 1
			} else {
				fmt.Printf(strconv.Itoa(s_value) + "*")
				s_value = s_value - 1
			}

		}
		fmt.Println()
	}

}
