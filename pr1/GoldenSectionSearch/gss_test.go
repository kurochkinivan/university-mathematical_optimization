package goldensectionsearch

import (
	"math/rand"
	"optimization/basefuncs"
	"testing"
)

func TestGSS(t *testing.T) {
	epsilon := 1e-07

	for {
		a, b := -10+rand.Float64()*(10+10), -10+rand.Float64()*(10+10)
		if a > b {
			a, b = b, a
		}

		t.Logf("a: %g, b: %g\n", a, b)
		min := GoldenSectionSearch(basefuncs.F1, a, b, epsilon)

		t.Logf("Минимум функции F на отрезке [%g; %g] = %g\n\n", a, b, min)
	}

}
