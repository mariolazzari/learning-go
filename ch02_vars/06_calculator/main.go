package main

import (
	"fmt"
	"strconv"
	"strings"
)

// calculate() returns the sum of the two parameters
func calculate(value1 string, value2 string) float64 {
	// Your code goes here.String('\n')

	// Convert the first string to a float64
	f1, err := strconv.ParseFloat(strings.TrimSpace(value1), 64)
	if err != nil {
		fmt.Println("Error", err)
		panic(err)
	}

	// Convert the second string to a float64
	f2, err := strconv.ParseFloat(strings.TrimSpace(value2), 64)
	if err != nil {
		fmt.Println("Error", err)
		panic(err)
	}

	// Calculate and return the result
	return f1 + f2
}

func main() {
	sum := calculate("10", "20")
	fmt.Println(sum)
}
