package main

import (
	"fmt"
	"optimization/basefuncs"
	consts "optimization/constants"
	hookejeeves "optimization/pr7/hookeJeeves"
)

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

	minPoint := hookejeeves.HookeJeeves(func(p hookejeeves.Point) float64 {
		return basefuncs.F5(basefuncs.NewPoint(p.X, p.Y))
	}, hookejeeves.Point{X: 0, Y: 10}, consts.H, consts.Epsilon, consts.Lambda, consts.HDecreaseFactor)
	fmt.Printf("Метод Хука-Дживса: %v; значение функции в этой точке: %g\n",
		minPoint, basefuncs.F5(basefuncs.NewPoint(minPoint.X, minPoint.Y)))
}
