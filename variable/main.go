package main

import "fmt"

func main() {
	// Variablel decleration
	const COUNTRY = "Bangladesh"

	var name string
	var age int
	var phone string
	var city string
	var merital_status bool

	// value assign
	name = "Mehedi"
	age = 30
	phone = "4444444"
	city = "Feni"
	merital_status = false

	// display the
	fmt.Println("Name is", name)
	fmt.Println("Age", age)
	fmt.Println("City", city)
	fmt.Println("Phone", phone)
	fmt.Println("Merital status is", merital_status)
	fmt.Println("Country", COUNTRY)

	// multiple variable initial
	var a, b, c int
	a = 10
	b = 20
	c = 30
	result := a + b + c // Dynamiclly type assigined
	fmt.Println("The sum of a,b,c is  = ", result)

	// type casting
	i, j := 10, 20.5
	r := float64(i) + j
	fmt.Println(r)

	v := pass_value_function("Amam", 30) // Dynamiclly type assigined
	fmt.Println(v)

	fmt.Println()
	variable_type_check()

}

func pass_value_function(name string, age int) string {
	nam := name
	return nam
}

func variable_type_check() {
	name := "Mehdi" // Dynamiclly type assigined
	age := 30
	is_meried := false
	car_price := 2000000.500

	fmt.Printf("Name is %v and\n%v years old\n", name, age)
	fmt.Printf("Name type is :  %T\n", name)
	fmt.Printf("Age type is :  %T\n", age)
	fmt.Printf("Merital Status  is :  %T", is_meried)
	fmt.Printf("Car Price type is :  %T", car_price)

}
