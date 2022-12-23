package mathFuncs

import "math"

// This function returns the Sum of the X table and the sum of X table squares
func sumX(limit float64) (float64, float64) {
	var sum, squarSum float64

	for i := 1; i < int(limit); i++ {
		sum += float64(i)
		squarSum += float64(i * i)

	}

	return sum, squarSum
}

// This function returns the Sum of the Y table and the sum of Y table squares
func sumY(sliceY []float64) (float64, float64) {
	var sum, squarSum float64

	for _, el := range sliceY {
		sum += el
		squarSum += el * el

	}

	return sum, squarSum
}

// This function returns the Sum of products of X and Y tables
func sumXY(sliceY []float64) float64 {
	var sum float64

	for i := 1; i < len(sliceY); i++ {
		sum += float64(i) * sliceY[i]
	}
	return sum
}

// This function returns a Pearson correlation coefficient
func PearCor(sliceY []float64) float64 {
	limit := float64(len(sliceY))

	xSum, xSqSum := sumX(limit)
	ySum, ySqSum := sumY(sliceY)
	xySum := sumXY(sliceY)

	r := (limit*xySum - xSum*ySum) / math.Sqrt((limit*xSqSum-(xSum*xSum))*(limit*ySqSum-(ySum*ySum)))

	return r
}
