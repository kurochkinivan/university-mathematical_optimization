package main

import (
	"fmt"
	"math"
	"optimization/constants"
	gss "optimization/pr2/GoldenSectionSearch"

	"gonum.org/v1/gonum/mat"
)

var hessian = mat.NewDense(2, 2, []float64{2, 4, 4, 10})

func f(v *mat.VecDense) float64 {
	x1, x2 := v.AtVec(0), v.AtVec(1)
	return math.Pow(x1+2*x2, 2) + math.Pow(x2-3, 2)
}

func grad(v *mat.VecDense) *mat.VecDense {
	x1, x2 := v.AtVec(0), v.AtVec(1)
	return mat.NewVecDense(2, []float64{
		2 * (x1 + 2*x2),
		4*(x1+2*x2) + 2*(x2-3),
	})
}

func norm(v *mat.VecDense) float64 {
	x1, x2 := v.AtVec(0), v.AtVec(1)
	return math.Sqrt(x1*x1 + x2*x2)
}

func NewtonMethod(f func(*mat.VecDense) float64, H *mat.Dense, x0 *mat.VecDense, eps1, eps2 float64, M int) *mat.VecDense {
	x := x0
	k := 0

	for {
		gradient := grad(x)
		if norm(gradient) <= eps1 {
			return x
		}

		if k >= M {
			return x
		}

		invHessian := new(mat.Dense)
		invHessian.Inverse(H)

		// Положительно определенная матрица
		var t float64
		d := mat.NewVecDense(2, nil)
		basisVec := mat.NewVecDense(2, []float64{1, 1})
		if mat.Inner(basisVec, invHessian, basisVec) > 0 {
			d.MulVec(invHessian, gradient)
			d.ScaleVec(-1, d)
			t = 1
		} else {
			d = gradient
			d.ScaleVec(-1, d)
			t = gss.GoldenSectionSearch(func(v float64) float64 { // возможно затенение
				expr := mat.NewVecDense(2, nil)
				expr.ScaleVec(v, gradient)
				expr.AddVec(x, expr)
				return f(expr)
			}, 0, 1, 1e-03)
		}

		xNew := mat.NewVecDense(2, nil)
		xNew.ScaleVec(t, d)
		xNew.AddVec(xNew, x)

		sub := mat.NewVecDense(2, nil)
		sub.SubVec(xNew, x)
		if norm(sub) < eps2 && math.Abs(f(xNew)-f(x)) < eps2 {
			return xNew
		}

		x = xNew
		k++
	}
}

func main() {
	x0 := mat.NewVecDense(2, []float64{-100, 100})
	min := NewtonMethod(f, hessian, x0, constants.Epsilon, constants.Epsilon, constants.Iterations)
	fmt.Printf("Метод Ньютона: %v\nЗначение функции в точке минимума: %g\n", min, f(min))
}
