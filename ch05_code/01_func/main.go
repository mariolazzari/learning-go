package main

import "fmt"

func main() {
	doSomething()
}

func doSomething() {
	fmt.Println("Doing something...")

	sum := addValues(1, 2)
	fmt.Println(sum)

	sum, _ = addAllValues(1, 2, 3, 4)
	fmt.Println(sum)
}

func addValues(val1, val2 int) int {
	return val1 + val2
}

func addAllValues(vals ...int) (int, int) {
	total := 0
	for _, v := range vals {
		total += v
	}
	return total, len(vals)
}
