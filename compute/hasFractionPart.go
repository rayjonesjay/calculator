package compute

import (
	"math"
)

func HasFractionPart(f float64) bool {

	_, fractionPart := math.Modf(f)

	return fractionPart != 0.0
}
