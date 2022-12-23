package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	err := InpErr()
	if err != nil {
		log.Print(err)
		return
	}
	//"numbers" is variable containing a slice of float64 from test file
	numbers, err := ReadFile()
	if err != nil {
		log.Print(err)
		return
	}
	a, b, r := Calculator(numbers)
	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", b, a)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r)
}

func InpErr() error {
	if len(os.Args) != 2 {
		return errors.New("error: wrong number of arguments\n\nUsage: go run main.go data.txt\n")
	}

	return nil
}

// This function gets the data (numbers) from test file and coverts the data to slice of float64 ---> []float64{}
func ReadFile() (nums []float64, err error) {
	fileName := os.Args[1]
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return []float64{}, errors.New("error: failed reading data from file")
	}

	lines := strings.Split(string(data), "\n")
	// Assign cap to avoid resize on every append.
	nums = make([]float64, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		// Atoi better suits the job when we know exactly what we're dealing
		// with. Scanf is the more general option.
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, float64(n))
	}

	return nums, nil
}

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
func LinReg(xSum float64, xSqSum float64, ySum float64, xySum float64, sliceY []float64, limit float64) (float64, float64) {
	a := ((ySum * xSqSum) - (xSum * xySum)) / (limit*xSqSum - (xSum * xSum))

	b := (limit*xySum - (xSum * ySum)) / (limit*xSqSum - (xSum * xSum))

	return a, b
}

// This function returns a Pearson correlation coefficient
func PearCor(xSum float64, xSqSum float64, ySum float64, ySqSum float64, xySum float64, sliceY []float64, limit float64) float64 {
	r := (limit*xySum - xSum*ySum) / math.Sqrt((limit*xSqSum-(xSum*xSum))*(limit*ySqSum-(ySum*ySum)))

	return r
}

// This function prepares, collects and distributes variables requred for all the functions calculations
// Also it collects and returns the final result to the main function
func Calculator(sliceY []float64) (float64, float64, float64) {
	limit := float64(len(sliceY))

	xSum, xSqSum := SumX(limit)
	ySum, ySqSum := SumY(sliceY)
	xySum := SumXY(sliceY)

	a, b := LinReg(xSum, xSqSum, ySum, xySum, sliceY, limit)
	r := PearCor(xSum, xSqSum, ySum, ySqSum, xySum, sliceY, limit)

	return a, b, r
}
