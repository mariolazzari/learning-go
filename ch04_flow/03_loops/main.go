package main

import "fmt"

func main() {
	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	// classic
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	// range
	for i := range colors {
		fmt.Println(colors[i])
	}

	// foreach -> idx, val
	for _, color := range colors {
		fmt.Println(color)
	}

	// range with map
	states := make(map[string]string)
	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"

	for key, value := range states {
		fmt.Println(key, value)
	}

	// while
	value := 0
	for value < 10 {
		fmt.Println("Value:", value)
		value++
	}

	// labels
	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum:", sum)

		if sum > 200 {
			goto theEnd
		}
	}

theEnd:
	fmt.Println("You exceeded 200")
}
