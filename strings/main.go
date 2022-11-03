package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("welcome to String")

	// string variable diclare and value assign
	var s string = "string string string gggg"
	var str string
	str = "Mehedi"
	str2 := "hossain"

	fmt.Println(s)
	fmt.Println(str2)
	fmt.Println("String is : ", str)
	print_line()

	// using string Split Method
	split_method()

	// using string Replace method
	replace_method()

	// using string join method
	join_method()

	// Name collect form list of email using string split() method
	name_extract_form_emails()

	// using Field() method split the string which work on based os white space
	field_method()

	// use count() method calculate and return perticular pattern how many times exists.
	count_method()

	//use HasPrefix()
	has_prefix_method()

}

func count_method() {
	str := "Lorem ipsum dolor sit amet, consectetur Lorem adipiscing elit."
	count := strings.Count(str, "Lorem")
	fmt.Println(count)
	print_line()

	// // using string Count method
	// str2 := "Mehedi"
	// char_volume_count := strings.Count(str2, "e")
	// fmt.Println(char_volume_count)
}

func split_method() {
	var str = "Mehedi"
	split_str := strings.Split(str, "")
	for i := 0; i < len(split_str); i++ {
		fmt.Println("Split string char is : ", split_str[i])
	}
	print_line()
	// y := strings.Split(str2, "")
	// fmt.Println(y)
}

func replace_method() {
	str := "Mehedi Hosain"
	r_str := strings.Replace(str, "e", "a", 2)
	fmt.Println(r_str)
	print_line()
}

func join_method() {
	s_str := []string{"My", "name", "is"}
	join_str := strings.Join(s_str, "-")
	fmt.Println(join_str)
	print_line()
}

func has_prefix_method() {
	lorem_string := "Lorem ipsum dolor sit amet, consectetur Lorem adipiscing elit. sed do eiusmod tempor."
	l_string := strings.Fields(lorem_string)
	var work_list []string
	var word_count = 0
	for _, l := range l_string {
		if strings.HasPrefix(l, "Lo") {
			word_count += 1
			work_list = append(work_list, l)
		}
	}
	fmt.Println(work_list)
	fmt.Println(word_count)
	fmt.Println(len(work_list))
	print_line()
}

func field_method() {
	lorem_string := "Lorem ipsum dolor sit amet, consectetur Lorem adipiscing elit. sed do eiusmod tempor."
	l_string := strings.Fields(lorem_string)
	fmt.Println(l_string)

	fmt.Printf("Type of l_string is %T\n", l_string)

	for _, l := range l_string {
		fmt.Println(l)
	}
	fmt.Println(len(l_string))
	print_line()
}

func name_extract_form_emails() {
	emaillist := []string{"me@gmail.com", "ta@gmail.com", "am@gmail.com", "su@gmail.com"}
	var name_list []string
	for _, email := range emaillist {
		split_name := strings.Split(email, "@")
		name := split_name[0]
		name_list = append(name_list, name)

	}
	fmt.Println("Output os Extract name list form email:", name_list) //Output is [me ta am su]
	print_line()
}

func string_declare_and_value_assign() {
	var s string = "string string string gggg"
	var str string
	str = "Mehedi"
	str2 := "hossain"

	fmt.Println(s)
	fmt.Println(str2)
	fmt.Println("String is : ", str)
	print_line()
}

func print_line() {
	fmt.Println("--------------------------------")
}
