package main

import "fmt"

func main() {
	msg := "Hi 😂 and 😭"
	// convert to rune slice
	runes := []rune(msg)
	// print the 4th rune
	fmt.Printf("%c\n", runes[3])
	fmt.Printf("%c\n", runes[9])
}
