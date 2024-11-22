package midpointalgorithm

import (
	"math"

	"gonum.org/v1/gonum/diff/fd"
)

func MidpointAlgorithm(f func(float64) float64, a, b float64, epsilon float64) float64 {
	for {
		x := (a + b) / 2

		deriv := fd.Derivative(f, x, nil)
		if b-a <= epsilon || math.Abs(deriv) <= epsilon {
			return x
		}

		if deriv > 0 {
			b = x
		} else {
			a = x
		}

	}
}
