package hookejeeves

import (
	"math"
)

type Point struct {
	X, Y float64
}

func HookeJeeves(f func(Point) float64, x0 Point, h0, epsilon, lambdaVal, hDecreaseFactor float64) Point {
	x1 := x0
	h := h0

	for {
		x2 := exploratorySearch(f, x1, h)
		if norm(x2, x1) < epsilon {
			return x2
		}

		x3 := patternSearch(x1, x2, lambdaVal)
		x4 := exploratorySearch(f, x3, h)

		if norm(x4, x3) > epsilon {
			x1 = x2
		} else {
			h /= hDecreaseFactor
		}
	}
}

func exploratorySearch(f func(Point) float64, p Point, h float64) Point {
	bestP := p
	bestValue := f(p)

	directions := []Point{
		{h, 0},
		{-h, 0},
		{0, h},
		{0, -h},
	}

	for _, dir := range directions {
		newP := Point{p.X + dir.X, p.Y + dir.Y}
		newValue := f(newP)
		if newValue < bestValue {
			bestP = newP
			bestValue = newValue
		}
	}

	return bestP
}

func patternSearch(p1, p2 Point, lambda float64) Point {
	return Point{
		X: p1.X + lambda*(p2.X-p1.X),
		Y: p1.Y + lambda*(p2.Y-p1.Y),
	}
}

func norm(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.X-p1.X, 2) + math.Pow(p2.Y-p1.Y, 2))
}
