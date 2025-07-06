// Package num provides functions to do math.
package num

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add computes the sum of two integer values and returns the result.
//
// The details can be found at [Addition].
//
// [Addition]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](x, y T) T {
	return x + y
}
