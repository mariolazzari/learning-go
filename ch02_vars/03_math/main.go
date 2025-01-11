package main

import (
	"fmt"
	"math"
)

func main() {
	i1, i2, i3 := 1, 2, 3
	intSum := i1 + i2 + i3
	fmt.Println("Int sum =", intSum)

	f1, f2, f3 := 1.123, 2.2345, 3.345
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum =", floatSum)
	floatSum = math.Round(floatSum*100) / 100
	fmt.Println("Rounded sum =", floatSum)

	circleRadius := 15.5
	circ := circleRadius * 2 * math.Pi
	fmt.Printf("Circumference = %.2f", circ)
}
