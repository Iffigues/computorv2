package util

import (
	"testing"
)

type def struct {
	give int
	wait int
}

func getTest() (e []def) {
	return []def{
		def{-0, 1},
		def{-1, -1},
		def{-2, -1},
		def{-3, -1},
		def{-4, -1},
		def{0, 1},
		def{1, 1},
		def{2, 2},
		def{3, 6},
		def{4, 24},
	}
}

func TestFactorial(t *testing.T) {
	given := getTest()
	factorial(-0.0)
	for _, val := range given {
		rest := factorial(val.give)
		if rest != val.wait {
			t.Fatalf("%d %d %d", val.give, val.wait, rest)
		}
	}
}
