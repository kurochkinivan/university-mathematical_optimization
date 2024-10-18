package main

import (
	"fmt"
	"optimization/basefuncs"
	consts "optimization/constants"

	gss "optimization/pr1/GoldenSectionSearch"
	bs "optimization/pr2/BitwiseSearch"
	mda "optimization/pr3/MidpointCircleAlgorithm"
)

func main() {
	fmt.Printf("Отрезок: [%g; %g]\nТочность: %g\n\n", consts.A, consts.B, consts.Epsilon)

	min := gss.GoldenSectionSearch(basefuncs.F2, consts.A, consts.B, consts.Epsilon)
	fmt.Printf("Метод золотого сечения: %g\n\n", min)

	min = bs.BitwiseSearch(basefuncs.F2, consts.A, consts.B, consts.Epsilon)
	fmt.Printf("Метод поразрядного поиска: %g\n\n", min)

	min = mda.MidpointAlgorithm(basefuncs.F2str, consts.A, consts.B, consts.Epsilon)
	fmt.Printf("Метод средней точки: %g\n\n", min)

}
