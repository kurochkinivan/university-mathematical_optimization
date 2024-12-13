package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// f вычисляет значение функции в точке Point.
func f(p Point) float64 {
	return math.Pow(p.X+4*p.Y, 2) + math.Pow(p.Y-3, 2)
}

// grad вычисляет градиент функции в точке Point.
func grad(p Point) Point {
	return Point{
		X: 2*p.X + 8*p.Y,
		Y: 8*p.X + 34*p.Y - 6,
	}
}

// norm вычисляет евклидову норму вектора Point.
func norm(p Point) float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// ConjugateGradient реализует метод сопряженного градиента.
func ConjugateGradient(eps float64, p0 Point) (Point, int) {
	N := 0
	pk := Point{}
	p := p0

	for norm(grad(p)) > eps && N < 100 {
		g := grad(p)
		pk = Point{X: -g.X, Y: -g.Y}

		// Минимизируем f(p + ak * pk) по ak.
		ak := minimizeAlpha(func(alpha float64) float64 {
			return f(Point{
				X: p.X + alpha*pk.X,
				Y: p.Y + alpha*pk.Y,
			})
		})

		p.X += ak * pk.X
		p.Y += ak * pk.Y
		N++
	}

	return p, N
}

// minimizeAlpha находит минимум одномерной функции по скалярному параметру.
func minimizeAlpha(f func(alpha float64) float64) float64 {
	// Простая реализация методом деления отрезка.
	low, high := -10.0, 10.0
	eps := 1e-6
	for high-low > eps {
		mid1 := low + (high-low)/3
		mid2 := high - (high-low)/3
		if f(mid1) < f(mid2) {
			high = mid2
		} else {
			low = mid1
		}
	}
	return (low + high) / 2
}

func main() {
	p0 := Point{X: -5, Y: 6}
	eps := 0.22

	pMin, N := ConjugateGradient(eps, p0)

	fmt.Printf("Минимум: p = (%f, %f), f(p) = %f, итераций: %d\n", pMin.X, pMin.Y, f(pMin), N)
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}