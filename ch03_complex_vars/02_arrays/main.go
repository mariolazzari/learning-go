package main

import "fmt"

func main() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println(colors)
	fmt.Println("Third color:", colors[2])

	// init array
	var numbers = [5]int16{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// len global function
	fmt.Println("Total numbers:", len(numbers))

}
