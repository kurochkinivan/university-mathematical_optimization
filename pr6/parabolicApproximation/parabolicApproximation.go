package parabolicapproximation

import (
	"fmt"
	"math"
	"math/rand/v2"
)

func ParabolicApproximation(f func(x float64) float64, a, b, eps float64, maxIterations int) (float64, error) {
	x1, x2, x3 := a, (a+b)/2, b
	y1, y2, y3 := f(x1), f(x2), f(x3)

	if !(x1 < x2 && x2 < x3) || !(y1 >= y2 && y2 <= y3) {
		return -1, fmt.Errorf("некорректные входные данные, измените начальные точки")
	}

	for k := range maxIterations {
		k++
		fmt.Println(x1, x2, x3)
		y1, y2, y3 = f(x1), f(x2), f(x3)
		numerator := (y1*(x2*x2-x3*x3) + y2*(x3*x3-x1*x1) + y3*(x1*x1-x2*x2))
		denominator := 2 * (y1*(x2-x3) + y2*(x3-x1) + y3*(x1-x2))

		var x float64
		if denominator == 0 {
			x = x1 + rand.Float64()*(x3-x1)
		} else {
			x = numerator / denominator
		}

		fmt.Println("x", x)
		if math.Abs(x1-x3) < eps {
			fmt.Println("кол-во итераций:", k)
			return x, nil
		}

		fx := f(x)
		if x2 <= x && x <= x3 {
			if fx <= y2 {
				x1, x2 = x2, x
			}
			if fx > y2 {
				x3 = x
			}
		}

		if x1 <= x && x <= x2 {
			if fx <= y2 {
				x2, x3 = x, x2
			}
			if fx > y2 {
				x1 = x
			}
		}
	}

	return -1, fmt.Errorf("кол-во итераций превышено")
}
