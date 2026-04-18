package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func main() {
	mario := MakePerson("mario", "lazzari", 51)
	fmt.Println("Mario", mario)

	marioPtr := MakePersonPointer("mario", "lazzari", 51)
	fmt.Println("Mario pointer", marioPtr)
}
