package main

import "fmt"

// Constraint: Stringer + underlying type int or float64
type Printable interface {
	fmt.Stringer
	~int | ~float64
}

type MyInt int

func (m MyInt) String() string {
	return fmt.Sprintf("MyInt: %d", m)
}

type MyFloat float64

func (m MyFloat) String() string {
	return fmt.Sprintf("MyFloat: %.2f", m)
}

// generic func
func PrintValue[T Printable](v T) {
	fmt.Println(v) // call String()
}

func main() {
	var a MyInt = 42
	PrintValue(a)

	var b MyFloat = 3.14
	PrintValue(b)
}
