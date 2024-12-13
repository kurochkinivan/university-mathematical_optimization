package parabolicapproximation

import (
	"fmt"
	"math"
	"math/rand/v2"
)

func ParabolicApproximation(f func(x float64) float64, a, b, eps float64, maxIterations int) (float64, error) {
	x1, x2, x3 := a, (a+b)/2, b
	y1, y2, y3 := f(x1), f(x2), f(x3)

	if !(x1 < x2 && x2 < x3 && y1 >= y2 && y2 <= y3) {
		return math.NaN(), fmt.Errorf("некорректные входные данные, измените начальные точки")
	}

	for k := 0; k < maxIterations; k++ {
		numerator := y1*(math.Pow(x2, 2)-math.Pow(x3, 2)) + y2*(math.Pow(x3, 2)-math.Pow(x1, 2)) + y3*(math.Pow(x1, 2)-math.Pow(x2, 2))
		denominator := 2 * (y1*(x2-x3) + y2*(x3-x1) + y3*(x1-x2))

		var x float64
		if denominator == 0 {
			x = x1 + rand.Float64() * (x3 - x1)
		} else {
			x = numerator / denominator
		}

		if x < x1 || x > x3 {
			x = x1 + rand.Float64()*(x3-x1)
		}

		if math.Abs(x1-x3) < eps {
			fmt.Println("кол-во итераций:", k)
			return x, nil
		}

		y := f(x)
		if x2 <= x && x <= x3 {
			if y <= y2 {
				x1, y1 = x2, y2
				x2, y2 = x, y
			} else if y > y2 {
				x3, y3 = x, y
			}
		} else if x1 <= x && x <= x2 {
			if y <= y2 {
				x3, y3 = x2, y2
				x2, y2 = x, y
			} else if y > y2 {
				x1, y1 = x, y
			}
		}
	}

	return math.NaN(), fmt.Errorf("кол-во итераций превышено")
}
