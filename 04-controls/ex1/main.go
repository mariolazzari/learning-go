package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	// preallocate
	ints := make([]int, 100)

	for i := range 100 {
		ints[i] = rand.IntN(100)
	}

	fmt.Println("ints:", ints)
}
