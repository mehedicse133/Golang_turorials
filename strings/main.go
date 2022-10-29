package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("welcome to String")

	var s string = "string"
	fmt.Println(s)

	var str string
	str = "Mehedi"
	fmt.Println("String is : ", str)

	// using string Split Method
	split_str := strings.Split(str, "")
	for i := 0; i < len(split_str); i++ {
		fmt.Println("Split string char is : ", split_str[i])
	}

	// using string Replace method
	str2 := "Mehedi Hosain"
	v := strings.Replace(str2, "e", "a", 2)
	fmt.Println(v)

	y := strings.Split(str2, "")
	fmt.Println(y)

	// using string Count method
	char_volume_count := strings.Count(str2, "e")
	fmt.Println(char_volume_count)

}
