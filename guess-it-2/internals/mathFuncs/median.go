package mathFuncs

import (
	"math"
	"sort"
)

// This function returns the Median of the given slice of numbers
func Median(numbers []float64) float64 {
	// Step 1. Sort the slice in ascending order
	sort.Float64s(numbers)
	// Step 2. Calculate the middle position
	var median float64
	n1 := numbers[len(numbers)/2]
	n2 := numbers[(len(numbers)/2)-1]

	if len(numbers)%2 == 1 {
		median = numbers[(len(numbers) / 2)]
	} else {
		// median = int(math.Floor((float64(n1) + float64(n2)) / 2))
		median = math.Round((float64(n1) + float64(n2)) / 2)
	}
	// 	Step 3. Return the value in the middle position
	return median
}
