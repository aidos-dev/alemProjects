package mathFuncs

import "math"

// This function returns the Standard Deviation of the given slice of numbers
func StanDev(numbers []float64) float64 {
	variance := Variance(numbers)

	StD := math.Sqrt(variance)

	return StD
}
