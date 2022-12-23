package ascii

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// FlagErr function makes sure that there are no errors in flag input
// func AlignFlagErr() error {
// 	if len(os.Args[3]) < 8 || os.Args[3][:8] != "--align=" {
// 		// log.Fatalf("Error: Wrong flag. Please input the correct flag: --align=\n\nUsage: go run . [STRING] [BANNER] [OPTION]\n\nExample: go run . something standard --align=right\n")
// 		return errors.New("error: Wrong flag. Please input the correct flag: --align=\n\nUsage: go run . [STRING] [BANNER] [OPTION]\n\nExample: go run . something standard --align=right\n")
// 	}

// 	_, flagNameError := AlignFlagCheck()
// 	if flagNameError {
// 		// log.Fatalf("Error: please input one of flag names: center, right or justify (left is the default value)\n\nUsage: go run . [STRING] [BANNER] [OPTION]\n\nExample: go run . something standard --align=right\n")
// 		return errors.New("error: please input one of flag names: center, right or justify (left is the default value)\n\nUsage: go run . [STRING] [BANNER] [OPTION]\n\nExample: go run . something standard --align=right\n")
// 	}
// 	return nil
// }

func AlignFlagCheck() (string, bool) {
	alignFlag := "left" // default value for the "alignFlag" variable is "left"

	if os.Args[3][8:] != "" {
		alignFlag = os.Args[3][8:] // this line checks for the flag parameter from a user and stores to the alignFlag variable
	}

	// _, _, _, alignFlag := FlagCheck()

	error := true
	switch alignFlag {
	case "left":
		error = false
	case "center":
		error = false
	case "right":
		error = false
	case "justify":
		error = false
	}

	if error {
		return "", true
	}

	return alignFlag, false
}

// AsciiWidth function calculates the width of a printable string in ascii-art format (big ascii-art letters)
// it is required to calculate how much space is going to take the printable string in ascii-art format
func AsciiWidth(ascLine [][]string) int {
	width := 0
	firstLine := 1
	for sym := range ascLine {
		width += len(ascLine[sym][firstLine])
	}
	// fmt.Println(width)
	return width
}

// DeltaCount function calculates the difference between the terminal window width and the width of a printable string in ascii-art format(big ascii-art letters)
func DeltaCount(ascLine [][]string) int {
	stringWidth := AsciiWidth(ascLine)
	termWidth := consoleSize()
	delta := termWidth - stringWidth - 1
	return delta
}

func Aligner(ascLine [][]string, newLineCount *int) [][]string {
	if len(os.Args) != 4 {
		return ascLine
	}

	delta := DeltaCount(ascLine)
	align, _ := AlignFlagCheck()
	symb := 0

	if len(ascLine) < 1 {
		return ascLine
	}

	stringsAll, err := PrepareStrings()
	if err != nil {
		return ascLine
	}
	argString := stringsAll[*newLineCount]
	argString = strings.TrimSpace(argString)

	if align == "right" {
		for row := 1; row < 9; row++ {
			for i := 0; i <= delta; i++ {
				ascLine[symb][row] = " " + ascLine[symb][row]
			}
		}
	}

	if align == "center" {
		for row := 1; row < 9; row++ {
			for i := 0; i <= delta/2; i++ {
				ascLine[symb][row] = " " + ascLine[symb][row]
			}
		}
	}

	if align == "justify" {

		// spaceCount is required to count the number of strings in the string argument to count the space for justify
		spaceCount := 0
		for _, el := range argString {
			if el == ' ' {
				spaceCount++
			}
		}
		// if there are no any spaces the function returns ascLine with no any changes
		if spaceCount == 0 {
			return ascLine
		}

		for i := 0; i < len(ascLine); i++ {
			if argString[i] == ' ' {
				for row := 1; row < 9; row++ {
					for j := 0; j <= delta/spaceCount-1; j++ {
						ascLine[i][row] = ascLine[i][row] + " "
					}
				}
			}
		}

	}

	return ascLine
}

func consoleSize() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)
	s = strings.TrimSpace(s)
	sArr := strings.Split(s, " ")

	width, err := strconv.Atoi(sArr[1])
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(width)
	return width
}
