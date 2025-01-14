package main

import "fmt"

func main() {
	poodle := Dog{"Poodle", 10}
	fmt.Println(poodle)

	// debug
	fmt.Printf("%+v\n", poodle)

	// interpolation
	fmt.Printf("Breed is %v\nWeight is %v\n", poodle.Breed, poodle.Weight)

	// set props
	poodle.Weight = 9
	fmt.Printf("Breed is %v\nWeight is %v\n", poodle.Breed, poodle.Weight)

}

// struct is a custom type
type Dog struct {
	Breed  string
	Weight int
}
