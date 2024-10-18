package bitwisesearch

import (
	"math"
)

func BitwiseSearch(f func(float64) float64, a, b float64, epsilon float64) float64 {
	h := (b - a) / 4
	x0 := a

	var x1 float64
	var i int = 1
	for {
		x0 = math.Round(x0*1e10) / 1e10
		x1 = math.Round(x1*1e10) / 1e10

		// fmt.Printf("%2.2d: h %g,\t[x0: %g, x1: %g]\n", i, h, x0, x1)

		x1 = x0 + h
		x1 = math.Round(x1*1e10) / 1e10

		// fmt.Printf("%2.2d: h %g,\t[x0: %g, x1: %g]\n", i, h, x0, x1)
		i++

		if f(x0) > f(x1) {
			x0 = x1
			if a < x0 && x0 < b {
				continue
			}
		}

		if math.Abs(h) <= epsilon {
			return x0
		}

		x0 = x1
		h /= -4
	}
}
