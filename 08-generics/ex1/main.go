package main

import "fmt"

// Numeric constraint: all ints + all floats
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func double[T Number](n T) T {
	return 2 * n
}

func main() {
	n := double(2)
	fmt.Println(n)
}
