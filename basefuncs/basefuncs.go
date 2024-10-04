package basefuncs

import "math"

func F(x float64) float64 {
	return math.Pow(2*math.Pow((x+1), 2)*(5-x), 1.0/3.0) - 2
}

func Derivative(f func(float64) float64, x float64, h float64) float64 {
	return (f(x+h) - f(x-h)) / (2 * h)
}
