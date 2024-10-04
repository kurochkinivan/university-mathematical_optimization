package main

import (
	"fmt"
	"optimization/basefuncs"
	consts "optimization/constants"
	gss "optimization/pr1/GoldenSectionSearch"
)

func main() {
	fmt.Println(basefuncs.F(consts.X))

	fmt.Println(gss.GoldenSectionSearch(basefuncs.F, consts.A, consts.B))
}


