package ascii

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

func ColorIndex() (int, bool, int) {
	start, deviderFlag, end, _ := IndexFlagCheck()

	return start, deviderFlag, end
}

// ColorIndexErr function makes sure that there are no errors in flag input
func ColorIndexErr() error {
	if len(os.Args[3]) < 8 || os.Args[3][:8] != "--index=" {
		return errors.New("error: Wrong flag. Please input the correct flag: --index=\n\nUsage: go run cmd/main.go [STRING] OPTION] [INDEX]\n\nExample: go run . something --color=red --index=0\n")
	}

	start, deviderFlag, end, flagNameError := IndexFlagCheck()

	if start < 0 || end < 0 {
		return errors.New("error: please input correct index range. Negative numbers are not allowed\n\nUsage: go run cmd/main.go [STRING] [OPTION] [INDEX]\n\nExample: go run . something --color=red --index=0:2\n")
	}

	if start > end && deviderFlag {
		return errors.New("error: please input correct index range. Starting index should not be less than ending index\n\nUsage: go run cmd/main.go [STRING] [OPTION] [INDEX]\n\nExample: go run . something --color=red --index=0:2\n")
	}

	if flagNameError != nil && os.Args[3][8:] != "" {
		return errors.New("error: please input correct index as an integer\n\nUsage: go run . [STRING] [OPTION] [INDEX]\n\nExample: go run cmd/main.go something --color=red --index=0\n")
	}

	// this check saves the index range input from irrelevant and too high values
	// _, runeCounter := SpaceCounter()

	runeLimit := len(os.Args[1])

	if end-start > runeLimit || start > runeLimit-1 {
		return errors.New("error: please input correct index range. The range should not exceed the number of symbols\n\nUsage: go run cmd/main.go [STRING] [OPTION] [INDEX]\n\nExample: go run . something --color=red --index=0:2\n")
	}

	return nil
}

func IndexFlagCheck() (int, bool, int, error) {
	startStr := ""
	endStr := ""
	start := 0
	end := len(os.Args[1])
	var err1, err2 error
	var indexRange string
	deviderFlag := false

	if os.Args[3][8:] != "" {
		indexRange = os.Args[3][8:] // this line checks for the flag parameter from a user and stores to the inputColor variable
	}

	for _, el := range indexRange {
		if el == ':' {
			deviderFlag = true
		}
		if el != ':' && deviderFlag == false {
			startStr += string(el)
		}
		if el != ':' && deviderFlag == true {
			endStr += string(el)
		}
	}

	if startStr != "" {
		start, err1 = strconv.Atoi(startStr)
		if err1 != nil {
			return 0, deviderFlag, 0, err1
		}
	}

	if endStr != "" {
		end, err2 = strconv.Atoi(endStr)
		if err2 != nil {
			return 0, deviderFlag, 0, err2
		}
	}

	// colorInp, index, _, _ := FlagCheck()

	return start, deviderFlag, end, nil
}

func SpaceCounter() (int, int) {
	arg := os.Args[1]
	arg = strings.TrimSpace(arg)
	spaceCounter := 0
	runeCounter := 0
	for _, el := range arg {
		if el == ' ' {
			spaceCounter++
		}
		if el != ' ' && el != '\n' {
			runeCounter++
		}

	}
	return spaceCounter, runeCounter
}
