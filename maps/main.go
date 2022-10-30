package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to Maps ")

	//crate Maps

	var map1 = map[string]string{"name": "Mehedi", "city": "Feni"}
	map2 := map[string]string{"name": "Mehedi", "city": "Feni"}

	fmt.Println(map1)
	fmt.Println(map2)

	student := map[int]string{
		1: "Mehdi",
		2: "Amam",
		3: "Subbir",
		4: "Tanvir",
	}

	// print map values
	fmt.Println(student)

	// print map values using loop
	for id, name := range student {
		if name == "Subbir" {
			continue
		}
		fmt.Println(id, name)
	}
}
