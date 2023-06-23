package maths

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Abs returns the absolute value of x.
func Abs[T Number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
