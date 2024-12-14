package basefuncs

import "math"

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func F1(x float64) float64 {
	return math.Pow(2*math.Pow((x+1), 2)*(5-x), 1.0/3.0) - 2
}

func F2(x float64) float64 {
	return math.Exp(x) + (1 / (x + 2))
}

func F3(x float64) float64 {
	return math.Cos(x) / math.Pow(x, 2)
}

func F4(x float64) float64 {
	return (x - 4) / (math.Sqrt(math.Pow(x, 2) + 3))
}

func F5(p Point) float64 {
	return math.Pow(p.x+2*p.y, 2) + math.Pow(p.y-3, 2)
}

func Derivative(f func(float64) float64, x float64, h float64) float64 {
	return (f(x+h) - f(x-h)) / (2 * h)
}

func SecondDerivative(f func(float64) float64, x, h float64) float64 {
	return (f(x+h) - 2*f(x) + f(x-h)) / (h * h)
}
