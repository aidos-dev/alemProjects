package mathFuncs

// This function returns the Variance of the given slice of numbers
func Variance(numbers []float64) float64 {
	// Step 1. Get the mean of numbers by Average function
	mean := Average(numbers)
	// Step 2. Find each number's deviation from the mean
	var dfm float64
	var sDev float64
	var sumSq float64
	var variance float64
	for i := range numbers {
		dfm = numbers[i] - mean
		// Step 3. Square each deviation from the mean
		sDev = dfm * dfm
		// Step 4. Find the sum of squares
		sumSq += sDev
	}
	// Step 5. Divide the sum of squares by number of elements
	variance = sumSq / float64(len(numbers))

	// Step 6. Return the variance and Standard Deviation as float64

	return variance
}
