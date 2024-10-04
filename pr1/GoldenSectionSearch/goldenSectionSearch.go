package goldensectionsearch

import (
	"math"
)

func GoldenSectionSearch(f func(float64) float64, a float64, b float64) float64 {
	l := 0.001

	y := a + (3-math.Sqrt(5))/2*(b-a)
	z := a + b - y

	for {
		if f(y) <= f(z) {
			b = z
			z = y
			y = a + b - y 
		} else {
			a = y
			y = z
			z = a + b - z
		}
		
		delta := math.Abs(a-b)
	
		if l >= delta {
			return (a+b)/2
		}
	}
}
