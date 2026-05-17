package math

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Float | constraints.Integer
}

// Add returns the sum of two integers.
//
// Addition reference:
// https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
