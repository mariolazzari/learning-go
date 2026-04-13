package main

import "fmt"

func main() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	first2 := greetings[0:2]
	firstMid := greetings[1:4]
	last2 := greetings[3:]

	fmt.Println("First 2 slice:", first2)
	fmt.Println("Middle  slice:", firstMid)
	fmt.Println("Last 2 slice:", last2)
}
