package multimodal

import (
	"math"
)

func BruteForceModified(f func(float64) float64, a, b float64, L, eps float64) float64 {
	h := 2 * eps / L

	var minPoint float64 
	var minValue float64 = math.MaxFloat64 
	for x := a + h/2; x < b + h; x += h {
		if f(x) < minValue {
			minPoint = x
			minValue = f(x)
		}
	}

	return minPoint
}

func BruteForce(f func(float64) float64, A, B float64, n int) float64 {
	if n <= 0 {
		panic("n must be > 0")
	}
	
	var minValue float64 = math.MaxFloat64
	var minPoint float64 = A

	for i := 0; i <= n; i++ {
		point := A + (float64(i) * (B - A) / float64(n))
		value := f(point)

		if value < minValue {
			minPoint = point
			minValue = value
		}
	}

	return minPoint
}
