package bitwisesearch

import "math"

func BitwiseSearch(f func(float64) float64, a, b float64, epsilon float64) float64 {
	h := (b - a) / 4
	x0 := a

	var x1 float64
	for {
		x1 = x0 + h
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
		h = -h/4
	}
}
