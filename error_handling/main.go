package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Welcome to error handling")

	result, err := div(7, 2)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

}

func div(a int, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("connot divide by 0")
	}
	return a / b, nil

}
