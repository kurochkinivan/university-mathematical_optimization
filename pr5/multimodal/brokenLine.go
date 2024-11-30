package multimodal

import (
	"math"
)

func BrokenLine(f func(float64) float64, a, b, L, eps float64) float64 {
	if eps <= 0 {
		panic("eps must be > 0")
	}

	x0 := (f(a) - f(b) + L*(a+b)) / (2 * L)
	y0 := (f(a) + f(b) + L*(a-b)) / 2
	p0 := y0

	pairs := make([][2]float64, 0)
	delta1 := (f(x0) - p0) / (2 * L)
	x11, x12 := x0-delta1, x0+delta1
	p1 := (f(x0) + p0) / 2
	pairs = append(pairs, [2]float64{x11, p1}, [2]float64{x12, p1})

	minPair := findMinPair(pairs)
	pairs = removePair(pairs, minPair)
	delta2 := (f(minPair[0]) - minPair[1]) / (2 * L)
	x21 := minPair[0] - delta2
	x22 := minPair[0] + delta2
	p2 := (f(minPair[0]) + minPair[1]) / 2
	pairs = append(pairs, [2]float64{x21, p2}, [2]float64{x22, p2})

	for {
		minPair = findMinPair(pairs)
		pairs = removePair(pairs, minPair)

		delta := (f(minPair[0]) - minPair[1]) / (2 * L)
		x11 = minPair[0] - delta
		x12 = minPair[0] + delta
		p := (f(minPair[0]) + minPair[1]) / 2

		pairs = append(pairs, [2]float64{x11, p}, [2]float64{x12, p})

		if 2*L*delta <= eps {
			return f(minPair[0])
		}
	}
}

func findMinPair(pairs [][2]float64) [2]float64 {
	var minP float64 = math.MaxInt
	var result [2]float64
	for _, pair := range pairs {
		if pair[1] < minP {
			minP = pair[1]
			result = pair
		}
	}
	return result
}

func removePair(pairs [][2]float64, pairToRemove [2]float64) [][2]float64 {
	result := make([][2]float64, 0)
	for _, pair := range pairs {
		if pair[0] != pairToRemove[0] || pair[1] != pairToRemove[1] {
			result = append(result, pair)
		}
	}
	return result
}
