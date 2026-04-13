package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	ints := make([]int, 100)
	for i := range 100 {
		ints[i] = rand.IntN(100)

		switch {
		case ints[i]%2 == 0 && ints[i]%3 == 0:
			fmt.Println("Six!")
		case ints[i]%2 == 0:
			fmt.Println("Two!")
		case ints[i]%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}
}
