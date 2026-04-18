package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

const N = 10_000_000

func main() {
	start := time.Now()

	var people []Person
	for range N {
		people = append(people, Person{
			FirstName: "Mario",
			LastName:  "Lazzari",
			Age:       51,
		})
	}

	elapsed := time.Since(start)
	fmt.Println("size:", len(people))
	fmt.Println("time:", elapsed)

	start = time.Now()

	people = make([]Person, N)
	for i := range N {
		people[i] = Person{
			FirstName: "Mario",
			LastName:  "Lazzari",
			Age:       51,
		}
	}

	elapsed = time.Since(start)
	fmt.Println("size with make:", len(people))
	fmt.Println("time with make:", elapsed)
}
