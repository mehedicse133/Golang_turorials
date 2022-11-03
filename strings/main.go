package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("welcome to String")

	var s string = "string string string gggg"
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

	// using string join method
	s_str := []string{"My ", "name", "is"}
	// g := uuuu[1:]
	// fmt.Println(g)
	fmt.Println(strings.Join(s_str, "-"))

	mystr := "mehe@gmai.com"
	split_mystr := strings.Split(mystr, "@")
	take_name := split_mystr[0]
	fmt.Println(take_name)
	print_line()

	// Name collect form list of email using string split method
	emaillist := []string{"me@gmail.com", "ta@gmail.com", "am@gmail.com", "su@gmail.com"}
	var name_list []string
	for _, email := range emaillist {
		split_name := strings.Split(email, "@")
		name := split_name[0]
		name_list = append(name_list, name)

	}
	fmt.Println("Output os Extract name list form email:", name_list) //Output is [me ta am su]
	print_line()

	// using Field method split the string which work on based os white space
	lorem_string := "Lorem ipsum dolor sit amet, consectetur Lorem adipiscing elit. sed do eiusmod tempor."
	l_string := strings.Fields(lorem_string)
	fmt.Println(l_string)

	fmt.Printf("Type of l_string is %T\n", l_string)

	for _, l := range l_string {
		fmt.Println(l)
	}
	fmt.Println(len(l_string))
	print_line()

	// use count method calculate and return perticular pattern how many times exists.
	CountMethod()

	var uu []string
	var word_count = 0
	for _, l := range l_string {
		if strings.HasPrefix(l, "Lo") {
			word_count += 1
			uu = append(uu, l)
		}
	}
	fmt.Println(uu)
	fmt.Println(word_count)
	fmt.Println(len(uu))
	print_line()
}

func CountMethod() {
	str := "Lorem ipsum dolor sit amet, consectetur Lorem adipiscing elit."
	count := strings.Count(str, "Lorem")
	fmt.Println(count)
	print_line()
}

func print_line() {
	fmt.Println("--------------------------------")
}
