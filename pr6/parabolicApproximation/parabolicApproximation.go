package parabolicapproximation

import (
	"fmt"
	"math"
	"math/rand"
)


func ParabolicApproximation(f func(x float64) float64, a, b, eps float64, maxIterations int) (float64, error) {
	if eps <= 0 {
		return 0, fmt.Errorf("epsilon must be > 0")
	}

	// Initial points
	x1, x2, x3 := a, (a+b)/2, b
	y1, y2, y3 := f(x1), f(x2), f(x3)

	// Ensure that the conditions for the method are met
	if !(x1 < x2 && x2 < x3 && y1 >= y2 && y2 <= y3) {
		return 0, fmt.Errorf("invalid input data, change the initial points")
	}

	// Start the iterative process
	for i := 0; i < maxIterations; i++ {
		// Calculate the numerator and denominator for the new point
		numerator := y1*(x2*x2-x3*x3) + y2*(x3*x3-x1*x1) + y3*(x1*x1-x2*x2)
		denominator := 2*y1*(x2-x3) + y2*(x3-x1) + y3*(x1-x2)

		var xTilda float64

		// Handle the case where denominator is zero
		if denominator == 0 {
			xTilda = rand.Float64()*(x3-x1) + x1
		} else {
			xTilda = numerator / denominator
		}

		// Check if the calculated xTilda is within bounds, else generate a new one
		if xTilda < x1 || xTilda > x3 {
			xTilda = rand.Float64()*(x3-x1) + x1
		}

		// Calculate function value at the new point
		yTilda := f(xTilda)

		// Check if the stopping condition is met
		if math.Abs(x1-x3) < eps {
			return xTilda, nil
		}

		// Update points based on the comparison of yTilda with current y values
		if xTilda >= x2 && xTilda <= x3 {
			if yTilda <= y2 {
				x1, y1 = x2, y2
				x2, y2 = xTilda, yTilda
			} else {
				x3, y3 = xTilda, yTilda
			}
		} else if xTilda <= x2 && xTilda >= x1 {
			if yTilda <= y2 {
				x3, y3 = x2, y2
				x2, y2 = xTilda, yTilda
			} else {
				x1, y1 = xTilda, yTilda
			}
		}
	}

	return 0, fmt.Errorf("method did not converge within the maximum number of iterations")
}
