package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	// struct literal style without names
	mario := Employee{"Mario", "Lazzari", 1}

	// struct literal style with names
	maria := Employee{
		firstName: "Maria",
		lastName:  "Lazzari",
		id:        2,
	}

	// var declaration
	var mariarosa Employee
	mariarosa.id = 3
	mariarosa.firstName = "Mariarosa"
	mariarosa.lastName = "Lazzari"

	fmt.Println("Mario:", mario)
	fmt.Println("Maria:", maria)
	fmt.Println("Mariarosa:", mariarosa)
}
