package calc

import (
	"math"
)

func Case1(n float64, outputs chan float64) {
	outputs <- math.Sqrt((5 * math.Pow(n, 2)) + 4)
}

func Case2(n float64, outputs chan float64) {
	outputs <- math.Sqrt((5 * math.Pow(n, 2)) - 4)
}
