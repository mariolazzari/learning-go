package main

import (
	"fmt"
	"sort"
)

func main() {
	// slice: no dim in []
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	// create slice with make
	var colors2 = make([]string, 0, 3)
	fmt.Println("Make slice", colors2)
	// append multiple elements
	colors2 = append(colors2, "Red", "Green", "Blue")
	fmt.Println("Append miltiple", colors2)

	// add element to slice (returns new slice)
	colors = append(colors, "Black")
	fmt.Println(colors)

	// remove first item
	colors = colors[1:]
	fmt.Println(colors)

	// remove last imte
	colors = colors[:len(colors)-1]
	fmt.Println(colors)

	colors = remove(colors, 1)
	fmt.Println(colors)

	// init slice with make
	numbers := make([]int, 5, 5)
	numbers[0] = 134
	numbers[1] = 21
	numbers[2] = 55
	numbers[3] = 333
	numbers[4] = 3
	fmt.Println(numbers)

	numbers = append(numbers, 1)
	fmt.Println(numbers)

	// sort slice
	sort.Ints(numbers)
	fmt.Println(numbers)

}

// remove element from slice
func remove(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}
