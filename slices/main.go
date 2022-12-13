package main

import "fmt"

func main() {
	fmt.Println("The slices tutorials ")

	var books = []string{"book1", "book2", "book3", "book4"}
	fmt.Println(books)
	fmt.Println("The length of arry books is : ", len(books))

	books = append(books, "book4", "book5")
	fmt.Println(books)
	fmt.Println("Afer and some books the length of arry books is : ", len(books))

	slice := books[1:3]
	fmt.Println(slice)

	s := books[2]
	fmt.Println(s)

	arr()

	fmt.Println("hi")
}

func arr() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a)
	fmt.Println("The length of arry a is : ", len(a))

	a = append(a, 10, 20, 30, 40, 50, 60)
	fmt.Println("Afert add some value the length of arry a is : ", len(a))
	fmt.Println(a)
}
