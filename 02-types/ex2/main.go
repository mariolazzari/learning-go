package main

import "fmt"

func main() {
	const value = 1
	i := value
	var f float64 = value // inferred

	fmt.Printf("i = %d, f = %.3f\n", i, f)
}
