package multimodal

import (
	"math"
)

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
