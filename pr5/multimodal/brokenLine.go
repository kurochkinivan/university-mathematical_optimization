package multimodal

func BrokenLine(f func(float64) float64, a float64, b float64, L float64, eps float64) float64 {
	if eps <= 0 {
		panic("eps must be > 0")
	}

	// Шаг 1
	x0 := (f(a) - f(b) + L*(a+b)) / (2 * L)
	p0 := (f(a) + f(b) + L*(a-b)) / 2

	p := p0

	// Список пар точка-значение функции
	pairs := make([][2]float64, 0)

	// 1-й шаг
	y0 := f(x0)
	delta := (y0 - p0) / (2 * L)
	x1 := x0 - delta
	x2 := x0 + delta
	p = (y0 + p0) / 2
	pairs = append(pairs, [2]float64{x1, p})
	pairs = append(pairs, [2]float64{x2, p})
	x0, p0 = x2, p

	// 2-й шаг
	y0 = f(x0)
	delta = (y0 - p0) / (2 * L)
	x1 = x0 - delta
	x2 = x0 + delta
	p = (y0 + p0) / 2
	pairs = append(pairs, [2]float64{x1, p})
	pairs = append(pairs, [2]float64{x2, p})

	// 3-й шаг
	x0, p0 = minPair(pairs)
	pairs = removePair(pairs, x0, p0)

	for {
		y0 = f(x0)
		delta = (y0 - p0) / (2 * L)
		if 2*L*delta <= eps {
			return x0
		}
		x1 = x0 - delta
		x2 = x0 + delta
		p = (y0 + p0) / 2
		pairs = append(pairs, [2]float64{x1, p})
		pairs = append(pairs, [2]float64{x2, p})
		x0, p0 = minPair(pairs)
		pairs = removePair(pairs, x0, p0)
	}
}

func minPair(pairs [][2]float64) (float64, float64) {
	minPair := pairs[0]
	for _, pair := range pairs[1:] {
		if pair[1] < minPair[1] {
			minPair = pair
		}
	}
	return minPair[0], minPair[1]
}

func removePair(pairs [][2]float64, x float64, p float64) [][2]float64 {
	result := make([][2]float64, 0)
	for _, pair := range pairs {
		if pair[0] != x || pair[1] != p {
			result = append(result, pair)
		}
	}
	return result
}
