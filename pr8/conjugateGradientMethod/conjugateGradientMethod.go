package main

import (
	"fmt"
	"math"
	"optimization/constants"
	goldensectionsearch "optimization/pr1/GoldenSectionSearch"
)

type Point struct {
	X, Y float64
}

func NewPoint(x, y float64) Point {
	return Point{X: x, Y: y}
}

func ConjugateGradientMethod(f func(Point) float64, x0 Point, h float64, eps float64, iterations int) Point {
	x := x0
	n := 0

	for norm(grad(x)) > eps && n < iterations {
		y := Point{
			X: -grad(x).X,
			Y: -grad(x).Y,
		}

		minimize := func(alpha float64) float64 {
			return f(Point{
				X: x.X + alpha*y.X,
				Y: x.Y + alpha*y.Y,
			})
		}

		ak := goldensectionsearch.GoldenSectionSearch(minimize, 0, 100, eps)
		x = Point{
			X: x.X + ak*y.X,
			Y: x.Y + ak*y.Y,
		}
		n++
	}

	return x
}

func f(p Point) float64 {
	return math.Pow(p.X+2*p.Y, 2) + math.Pow(p.Y-3, 2)
}

func grad(p Point) Point {
	return Point{
		X: 2 * (p.X + 2*p.Y),
		Y: 4*p.X + 10*p.Y - 6,
	}
}

func norm(p Point) float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
	minPoint := ConjugateGradientMethod(f, NewPoint(0, 0), constants.H, constants.Epsilon, constants.Iterations)
	fmt.Println("minPoint:", minPoint)
}
