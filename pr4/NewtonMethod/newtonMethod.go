package newtonmethod

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/diff/fd"
)

func NewtonMethod(f func(float64) float64, x float64, epsilon float64) float64 {
	for {
		firstDeriv := fd.Derivative(f, x, &fd.Settings{
			Formula: fd.Forward,
			Step:    1e-1,
		})
		secondDeriv := fd.Derivative(f, x, &fd.Settings{
			Formula:     fd.Central2nd,
			Concurrent:  true,
			OriginKnown: true,
			OriginValue: f(0),
		})

		x -= firstDeriv / secondDeriv
		fmt.Println("firstDeriv: ", firstDeriv, "secondDeriv: ", secondDeriv, "x: ", x)

		if math.Abs(firstDeriv) <= epsilon {
			return x
		}
	}
}
