package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("welcome to Golang files")

	str := "I love my country"

	file, err := os.Create("./myfile.txt")
	if err != nil {
		panic(err)
	}
	l, _ := io.WriteString(file, str)
	if err != nil {
		panic(err)
	}
	fmt.Println(l)
	defer file.Close()

}
