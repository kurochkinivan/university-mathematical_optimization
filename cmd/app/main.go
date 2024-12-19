package main

import (
	"fmt"
	"math"
	goldensectionsearch "optimization/pr2/GoldenSectionSearch"
	"optimization/pr6/multimodal"
	parabolicapproximation "optimization/pr7/parabolicApproximation"
)

// hookejeeves "optimization/pr7/hookeJeeves"
// gss "optimization/pr1/GoldenSectionSearch"
// bs "optimization/pr2/BitwiseSearch"
// mda "optimization/pr3/MidpointCircleAlgorithm"
// nw "optimization/pr4/NewtonMethod"
// mm "optimization/pr5/multimodal"
// pa "optimization/pr6/parabolicApproximation"

func main() {
	// fmt.Printf("Отрезок: [%g; %g]\nТочность: %g\n\n", consts.A, consts.B, consts.Epsilon)

	// min := gss.GoldenSectionSearch(basefuncs.F2, consts.A, consts.B, consts.Epsilon)
	// fmt.Printf("Метод золотого сечения: %g\n\n", min)

	// min = bs.BitwiseSearch(basefuncs.F4, consts.A, consts.B, consts.Epsilon)
	// fmt.Printf("Метод поразрядного поиска: %g\n\n", min)

	// min = mda.MidpointAlgorithm(basefuncs.F4, consts.A, consts.B, consts.Epsilon)
	// fmt.Printf("Метод средней точки: %g\n\n", min)

	// min = nw.NewtonMethod(basefuncs.F4, consts.A, consts.B, consts.Epsilon)
	// fmt.Printf("Метод Ньютона: %g\n\n", min)

	// lipschitz := mm.LipschitzConstant(basefuncs.F3, consts.A, consts.B, consts.Iterations)
	// fmt.Printf("Константа Липшицца: %g\n", lipschitz)
	// min = mm.BruteForceModified(basefuncs.F3, consts.A, consts.B, lipschitz, consts.Epsilon)
	// fmt.Printf("Метод равномерного поиска: %g\n", min)
	// min = mm.BrokenLine(basefuncs.F3, consts.A, consts.B, lipschitz, consts.Epsilon)
	// fmt.Printf("Метод ломанной линии: %v\n\n", min)

	// ans, err := pa.ParabolicApproximation(basefuncs.F4, consts.A, consts.B, consts.Epsilon, consts.Iterations)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Метод параболической аппроксимации: %v\n\n", ans)

	// minPoint := hookejeeves.HookeJeeves(func(p hookejeeves.Point) float64 {
	// 	return basefuncs.F5(basefuncs.NewPoint(p.X, p.Y))
	// }, hookejeeves.Point{X: 0, Y: 10}, consts.H, consts.Epsilon, consts.Lambda, consts.HDecreaseFactor)
	// fmt.Printf("Метод Хука-Дживса: %v; значение функции в этой точке: %g\n",
	// 	minPoint, basefuncs.F5(basefuncs.NewPoint(minPoint.X, minPoint.Y)))

	f := func(x float64) float64 {
		return math.Sqrt(math.Pow(800+x, 2)+600*600) + math.Sqrt(math.Pow(400+(200-x), 2)+800*800)
	}
	min := goldensectionsearch.GoldenSectionSearch(f, 0, 200, 1e-08)
	fmt.Println(min, f(min))
	fmt.Println(f(0))
	fmt.Println(multimodal.BruteForce(f, 0, 200, 100000))

	f = func(x float64) float64 {
		return math.Pow(x, 2) + 8/x
	}
	parabolicapproximation.ParabolicApproximation(f, 1, 3, 1e-03, 10000)
}
