package main

import "fmt"

func main() {
	theAnswer := 42
	var result string

	if theAnswer < 0 {
		result = "negative"
	} else if theAnswer == 0 {
		result = "zero"
	} else {
		result = "positive"
	}

	fmt.Println("result is", result)

	// assignment in if
	if x := -42; x < 0 {
		fmt.Println("x is neative")
	} else if x == 0 {
		fmt.Println("x is 0")
	} else {
		fmt.Println("x is positive")
	}

}
