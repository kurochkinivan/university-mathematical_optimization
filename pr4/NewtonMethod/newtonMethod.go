package newtonmethod

import (
	"math"

	"gonum.org/v1/gonum/diff/fd"
)

func NewtonMethod(f func(float64) float64, A, B float64, epsilon float64) float64 {
	if epsilon <= 0 {
		panic("epsilon must be > 0")
	}
	
	x := (A + B) / 2
	for {
		firstDeriv := fd.Derivative(f, x, &fd.Settings{
			Formula: fd.Central,
		})
		secondDeriv := fd.Derivative(f, x, &fd.Settings{
			Formula: fd.Central2nd,
		})
		// fmt.Println("1st derivative:", basefuncs.Derivative(f, x, epsilon))
		// fmt.Println("2nd derivative:", basefuncs.SecondDerivative(f, x, epsilon))
		// fmt.Printf("[%g; %g]\tfirstDeriv: %g, secondDeriv: %g\n", x, f(x), firstDeriv, secondDeriv)

		if x - firstDeriv / secondDeriv < A || x - firstDeriv / secondDeriv > B {
			return x
		}

		x -= firstDeriv / secondDeriv

		if math.Abs(firstDeriv) <= epsilon {
			return x
		}
	}
}

