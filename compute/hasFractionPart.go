package compute

import (
	"math"
)

//HasFractionPart() checks if a float number has a fraction part
func HasFractionPart(f float64) bool {

	_, fractionPart := math.Modf(f)

	return fractionPart != 0.0  
}
