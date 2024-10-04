package main

import (
	"fmt"
	"optimization/basefuncs"
	consts "optimization/constants"
	gss "optimization/pr1/GoldenSectionSearch"
	bs "optimization/pr2/BitwiseSearch"
)

func main() {
	fmt.Printf("Отрезок: [%d; %d]\nТочность: %g\n\n", consts.A, consts.B, consts.Epsilon)

	min := gss.GoldenSectionSearch(basefuncs.F, consts.A, consts.B, consts.Epsilon)
	fmt.Printf("Метод золотого сечения: %g\n\n", min)

	min = bs.BitwiseSearch(basefuncs.F, consts.A, consts.B, consts.Epsilon)
	fmt.Printf("Метод поразрядного поиска: %g\n\n", min)
}