package main

import (
	"fmt"
	"math"
)

func main() {
	var b byte = math.MaxUint8
	b++

	var smallI int32 = math.MaxInt32
	smallI++

	var bigI uint64 = math.MaxUint64
	bigI++

	fmt.Printf("b = %d\n", b)
	fmt.Printf("smallI = %d\n", smallI)
	fmt.Printf("bigI = %d\n", bigI)
}
