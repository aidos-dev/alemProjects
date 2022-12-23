package mathFuncs

// This function returns the Average of the given slice of numbers
func Average(numbers []float64) float64 {
	var sum float64

	var res float64
	for _, el := range numbers {
		sum += el
	}
	res = sum / float64(len(numbers))
	return res
}
