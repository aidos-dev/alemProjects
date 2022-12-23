package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	"01.alem.school/git/Optimus/guess-it-2/internals/mathFuncs"
)

func main() {
	var dataArr []float64

	var input float64

	var tempMin, tempMax float64

	var min, max int

	var yFloat, xFloat float64

	var slope, intersept float64

	// var standardDeviasion float64

	// var pearCor float64

	r := bufio.NewScanner(os.Stdin)

	for r.Scan() {

		input, _ = strconv.ParseFloat(r.Text(), 64)

		xFloat++

		if input < 100 || input > 13000 {
			dataArr = append(dataArr, math.Round((mathFuncs.Average(dataArr))))
		} else {
			dataArr = append(dataArr, input)
		}

		// dataArr = append(dataArr, input)

		// pearCor = mathFuncs.PearCor(dataArr)

		// standardDeviasion = mathFuncs.StanDev(dataArr)

		intersept, slope = mathFuncs.LinReg(dataArr)

		if xFloat == 1 {
			yFloat = input
			// pearCor = 1
			// standardDeviasion = 10
			// correlation = 0
		} else {
			yFloat = slope*xFloat + intersept
		}

		// correlation = yFloat * pearCor

		// tempMin = yFloat - standardDeviasion

		// tempMax = yFloat + standardDeviasion

		tempMin = yFloat - 6

		tempMax = yFloat + 6

		min = int(math.Round(tempMin))

		max = int(math.Round(tempMax))

		// min = (int(math.Round((mathFuncs.Average(dataArr)) * slope))) - (int(math.Round((mathFuncs.StanDev(dataArr)))))
		// max = (int(math.Round((mathFuncs.Average(dataArr)) * slope))) + (int(math.Round((mathFuncs.StanDev(dataArr)))))

		// min = (int(math.Round((mathFuncs.Average(dataArr))))) - (int(math.Round((mathFuncs.StanDev(dataArr)))))
		// max = (int(math.Round((mathFuncs.Average(dataArr))))) + (int(math.Round((mathFuncs.StanDev(dataArr)))))

		fmt.Printf("%d %d\n", min, max)

	}
}
