package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	var dataArr []float64

	var input float64

	var min, max int

	r := bufio.NewScanner(os.Stdin)

	for r.Scan() {

		input, _ = strconv.ParseFloat(r.Text(), 64)

		if input < 100 || input > 200 {
			dataArr = append(dataArr, math.Round((Average(dataArr))))
		} else {
			dataArr = append(dataArr, input)
		}

		min = int(math.Round((Average(dataArr)))) - int(math.Round((StanDev(dataArr))))
		max = int(math.Round((Average(dataArr)))) + int(math.Round((StanDev(dataArr))))

		fmt.Printf("%d %d\n", min, max)

	}
}

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

// This function returns the Median of the given slice of numbers
func Median(numbers []float64) float64 {
	// Step 1. Sort the slice in ascending order
	// for j := 0; j < len(numbers)-1; j++ {
	// 	for i := 0; i < len(numbers)-1; i++ {
	// 		if numbers[i] > numbers[i+1] {
	// 			numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
	// 		}
	// 	}
	// }
	// buble sort was replaced with "sort.Float64s()" function as it is a standart library function in golang
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

// This function returns the Standard Deviation of the given slice of numbers
func StanDev(numbers []float64) float64 {
	variance := Variance(numbers)

	StD := math.Sqrt(variance)

	return StD
}
