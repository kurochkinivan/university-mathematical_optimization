package midpointalgorithm

import (
	"math"

	"github.com/TheDemx27/calculus"
)

func MidpointAlgorithm(f0 string, a, b float64, epsilon float64) float64 {
	f := calculus.NewFunc(f0)
	for {
		x := (a + b) / 2
		
		deriv := f.Diff(x)
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
