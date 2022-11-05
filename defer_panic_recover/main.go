package main

import "fmt"

func main() {
	//defer
	fmt.Println("welcome to defer_panic_recover")
	defer fmt.Println("end the execution")
	fmt.Println("welcome to defer")

	defer fmt.Println("A")
	defer fmt.Println("B")
	defer fmt.Println("C")

	team1 := panicc("amam0", "tanvir", "mehedi", "subbir")
	fmt.Println(team1)

	team2 := panicc("anik", "mamun", "rabbi")
	fmt.Println(team2)
}

func panicc(names ...string) []string {
	var team []string
	for _, name := range names {
		if name == "mehedi" {
			panic("The team is rejected,becuse there is a wrong person")
		} else {
			team = append(team, name)
		}

	}
	return team
}
