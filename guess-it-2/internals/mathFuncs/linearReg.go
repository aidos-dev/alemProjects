package mathFuncs

// This function returns the Sum of the X table and the sum of X table squares
func SumX(limit float64) (float64, float64) {
	var sum, squarSum float64

	for i := 1; i < int(limit); i++ {
		sum += float64(i)
		squarSum += float64(i * i)

	}

	return sum, squarSum
}

// This function returns the Sum of the Y table and the sum of Y table squares
func SumY(sliceY []float64) (float64, float64) {
	var sum, squarSum float64

	for _, el := range sliceY {
		sum += el
		squarSum += el * el

	}

	return sum, squarSum
}

// This function returns the Sum of products of X and Y tables
func SumXY(sliceY []float64) float64 {
	var sum float64

	for i := 1; i < len(sliceY); i++ {
		sum += float64(i) * sliceY[i]
	}
	return sum
}

// This function returns two variables for a Linear regression equation
func LinReg(sliceY []float64) (float64, float64) {
	limit := float64(len(sliceY))

	xSum, xSqSum := SumX(limit)
	ySum, _ := SumY(sliceY)
	xySum := SumXY(sliceY)

	a := ((ySum * xSqSum) - (xSum * xySum)) / (limit*xSqSum - (xSum * xSum))

	b := (limit*xySum - (xSum * ySum)) / (limit*xSqSum - (xSum * xSum))

	return a, b
}
