package multimodal

import (
	"math"

	"gonum.org/v1/gonum/diff/fd"
)

func LipschitzConstant(f func(float64) float64, a float64, b float64, n int) float64 {
	if n <= 0 {
		panic("n must be > 0")
	}
	
	step := (b - a) / float64(n-1)
	
	var maxDeriv float64
	for i := 0; i < n; i++ {
		x := a + float64(i)*step
		deriv := math.Abs(fd.Derivative(f, x, &fd.Settings{
			Formula: fd.Central,
			Step:   1e-6,
		})) 
		maxDeriv = max(maxDeriv, deriv)
	}

	return maxDeriv
}

func max(a, b float64) float64 {
	if a > b {
		return a 
	}
	return b
}
