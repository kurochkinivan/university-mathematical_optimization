package main

import (
	"fmt"
	"math"
	"optimization/constants"
	gss "optimization/pr2/GoldenSectionSearch"

	"gonum.org/v1/gonum/mat"
)

// Variant 6
// func targetFunction(v *mat.VecDense) float64 {
// 	x1, x2 := v.AtVec(0), v.AtVec(1)
// 	return math.Pow(x1+2*x2, 2) + math.Pow(x2-3, 2)
// }

// Variant 22
// func targetFunction(v *mat.VecDense) float64 {
// 	x1, x2 := v.AtVec(0), v.AtVec(1)
// 	return math.Pow(x1-8*x2, 2) + math.Pow(x2+1, 2)
// }

func targetFunction(v *mat.VecDense) float64 {
	x1, x2 := v.AtVec(0), v.AtVec(1)
	return -2*x1*x1 + 10*x1*x2 - 2*x2*x2 - 26*x1 + 2*x2
}

func norm(v *mat.VecDense) float64 {
	x1, x2 := v.AtVec(0), v.AtVec(1)
	return math.Sqrt(x1*x1 + x2*x2)
}

func grad(f func(v *mat.VecDense) float64, v *mat.VecDense) *mat.VecDense {
	x1, x2 := v.AtVec(0), v.AtVec(1)
	h := 1e-05
	return mat.NewVecDense(2, []float64{
		partialDerivX(f, x1, x2, h),
		partialDerivY(f, x1, x2, h),
	})
}

func partialDerivX(f func(v *mat.VecDense) float64, x1, x2, h float64) float64 {
	addH, subH := mat.NewVecDense(2, []float64{x1 + h, x2}), mat.NewVecDense(2, []float64{x1 - h, x2})
	return (f(addH) - f(subH)) / (2 * h)
}

func partialDerivY(f func(v *mat.VecDense) float64, x1, x2, h float64) float64 {
	addH, subH := mat.NewVecDense(2, []float64{x1, x2 + h}), mat.NewVecDense(2, []float64{x1, x2 - h})
	return (f(addH) - f(subH)) / (2 * h)
}

func secondPartialDerivX(f func(v *mat.VecDense) float64, x1, x2, h float64) float64 {
	addH, subH := mat.NewVecDense(2, []float64{x1 + h, x2}), mat.NewVecDense(2, []float64{x1 - h, x2})
	stable := mat.NewVecDense(2, []float64{x1, x2})
	return (f(addH) - 2*f(stable) + f(subH)) / math.Pow(h, 2)
}

func secondPartialDerivY(f func(v *mat.VecDense) float64, x1, x2, h float64) float64 {
	addH, subH := mat.NewVecDense(2, []float64{x1, x2 + h}), mat.NewVecDense(2, []float64{x1, x2 - h})
	stable := mat.NewVecDense(2, []float64{x1, x2})
	return (f(addH) - 2*f(stable) + f(subH)) / math.Pow(h, 2)
}

func secondMixedDeriv(f func(v *mat.VecDense) float64, x1, x2, h float64) float64 {
	addx1addx2, subx1addx2 := mat.NewVecDense(2, []float64{x1 + h, x2 + h}), mat.NewVecDense(2, []float64{x1 - h, x2 + h})
	addx1subx2, subx1subx2 := mat.NewVecDense(2, []float64{x1 + h, x2 - h}), mat.NewVecDense(2, []float64{x1 - h, x2 - h})

	return (f(addx1addx2) - f(subx1addx2) - f(addx1subx2) + f(subx1subx2)) / (4 * math.Pow(h, 2))
}

func computeHessian(f func(*mat.VecDense) float64, v *mat.VecDense) *mat.Dense {
	x1, x2 := v.AtVec(0), v.AtVec(1)
	h := 1e-05
	d2_dx := secondPartialDerivX(f, x1, x2, h)
	d2_dxdy := secondMixedDeriv(f, x1, x2, h)
	d2_dy := secondPartialDerivY(f, x1, x2, h)
	return mat.NewDense(2, 2, []float64{d2_dx, d2_dxdy, d2_dxdy, d2_dy})
}

func NewtonMethodModified(f func(*mat.VecDense) float64, x0 *mat.VecDense, eps1, eps2 float64, M int) *mat.VecDense {
	x := x0
	k := 0

	for {
		gradient := grad(f, x)
		if norm(gradient) <= eps1 {
			return x
		}

		if k >= M {
			return x
		}

		hessian := computeHessian(f, x)
		fmt.Println("hessian", hessian)

		invHessian := new(mat.Dense)
		invHessian.Inverse(hessian)

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
	x0 := mat.NewVecDense(2, []float64{0, 0})
	min := NewtonMethodModified(targetFunction, x0, constants.Epsilon, constants.Epsilon, 100000)
	fmt.Println(min, targetFunction(min))
}
