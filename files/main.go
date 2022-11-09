package main

import (
	"fmt"
	"io"
	"io/ioutil"
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
	readfile("./myfile.txt")

}

func readfile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("The file outoput is : ", string(data))

}
