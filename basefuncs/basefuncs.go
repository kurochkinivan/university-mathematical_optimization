package basefuncs

import "math"

const (
	F2str = "e^x + (1/(x+2))"
)

func F1(x float64) float64 {
	return math.Pow(2*math.Pow((x+1), 2)*(5-x), 1.0/3.0) - 2
}

func F2(x float64) float64 {
	return math.Exp(x) + (1 / (x + 2))
}

// Derivative returns the derivative of f at x using the symmetric difference quotient with stepsize h.
func Derivative(f func(float64) float64, x float64, h float64) float64 {
	return (f(x+h) - f(x-h)) / (2 * h)
}
