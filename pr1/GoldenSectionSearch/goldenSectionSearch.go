package goldensectionsearch

import (
	"fmt"
	"math"
	"os"
)

func GoldenSectionSearch(f func(float64) float64, a, b float64, epsilon float64) float64 {
	if epsilon <= 0 {
		panic("epsilon must be > 0")
	}
	
	sqrt5 := math.Sqrt(5)
	var currentX float64
	
	var oldA, oldB float64 = a, b
	for {
		alpha := a + ((3-sqrt5)/2)*(b-a)
		beta := a + ((sqrt5-1)/2)*(b-a)

		if f(alpha) <= f(beta) {
			a, oldA = oldA, a
			b, beta = beta, alpha
			currentX = alpha
			alpha = a + b - alpha
		} else {
			b, oldB = oldB, b 
			a, alpha = alpha, beta 
			currentX = beta
			beta = a + b - beta 
		}

		delta := b - a

		if delta <= epsilon {
			return currentX
		}

		if alpha == math.Inf(1) || b == math.Inf(1) || a == math.Inf(1) {
			fmt.Printf("\n[%g; %g]\talpha: %g, beta: %g\n", a, b, alpha, beta)
			fmt.Printf("Невозможно вычислить минимум функции F\n")
			os.Exit(1)
		}

	}
}
