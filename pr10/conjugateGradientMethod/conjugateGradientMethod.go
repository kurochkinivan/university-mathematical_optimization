package main

import (
	"fmt"
	"math"
	"optimization/constants"
	goldensectionsearch "optimization/pr2/GoldenSectionSearch"
)

type Vars [2]float64

func NewVars(x1, x2 float64) Vars {
	var vars Vars
	vars[0], vars[1] = x1, x2
	return vars
}

func ConjugateGradientMethod(f func(Vars) float64, x0 Vars, eps float64, iterations int) Vars {
	x := x0

	for n := 0; n < iterations; n++ {
		if norm(grad(x)) < eps {
			return x
		}

		gradient := grad(x)
		p := NewVars(-gradient[0], -gradient[1])

		minimize := func(alpha float64) float64 {
			return f(NewVars(x[0]+alpha*p[0], x[1]+alpha*p[1]))
		}

		alpha := goldensectionsearch.GoldenSectionSearch(minimize, -100, 100, eps)
		x = NewVars(x[0]+alpha*p[0], x[1]+alpha*p[1])
	}

	return x
}

func f(v Vars) float64 {
	return math.Pow(v[0]+2*v[1], 2) + math.Pow(v[1]-3, 2)
}

func grad(v Vars) Vars {
	return Vars{
		2 * (v[0] + 2*v[1]),
		4*v[0] + 10*v[1] - 6,
	}
}

func norm(v Vars) float64 {
	return math.Sqrt(v[0]*v[0] + v[1]*v[1])
}

func main() {
	minPoint := ConjugateGradientMethod(f, NewVars(0, 0), constants.Epsilon, constants.Iterations)
	fmt.Println("minPoint:", minPoint)
}
