package ascii

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func FlagCheck() (string, bool) {
	// alignFlag := flag.String("align", "left", "flag is used to to change the alignment of the output")
	// flag.Parse()

	alignFlag := "left" // default value for the "alignFlag" variable is "left"

	if os.Args[3][8:] != "" {
		alignFlag = os.Args[3][8:] // this line checks for the flag parameter from a user and stores to the alignFlag variable
	}

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

// FlagErr function makes sure that there are no errors in flag input
func FlagErr() error {
	if len(os.Args[3]) < 8 || os.Args[3][:8] != "--align=" {
		// log.Fatalf("Error: Wrong flag. Please input the correct flag: --align=\n\nUsage: go run . [STRING] [BANNER] [OPTION]\n\nExample: go run . something standard --align=right\n")
		return errors.New("error: Wrong flag. Please input the correct flag: --align=\n\nUsage: go run . [STRING] [BANNER] [OPTION]\n\nExample: go run . something standard --align=right\n")
	}

	_, flagNameError := FlagCheck()
	if flagNameError {
		// log.Fatalf("Error: please input one of flag names: center, right or justify (left is the default value)\n\nUsage: go run . [STRING] [BANNER] [OPTION]\n\nExample: go run . something standard --align=right\n")
		return errors.New("error: please input one of flag names: center, right or justify (left is the default value)\n\nUsage: go run . [STRING] [BANNER] [OPTION]\n\nExample: go run . something standard --align=right\n")
	}
	return nil
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
	align, _ := FlagCheck()
	symb := 0

	if len(ascLine) < 1 {
		return ascLine
	}

	stringsAll, err := PrepareStrings("TEST")
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

	// fmt.Println(len(ascLine))

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

// Converter function scans strings from the command line and slices of strings from FileReader function.
// FileReader sends to Converter template strings for ascii-art symbols. Then Converter transforms them into 2D slice of
// slices of ascii-art symbols
func Converter111(s string, lines []string, newLineCount *int) (ascLine [][]string) {
	// s = strings.TrimSpace(s)
	// s = strings.TrimRight(s, " ")
	// s = TrimRightSpaceCustom(s)
	runA := []rune(s)
	ascLine = [][]string{}
	symb := []string{}

	for _, r := range runA {

		for i := 0; i < 10; i++ {
			symb = append(symb, lines[int((r-32)*9)+i])
		}
		ascLine = append(ascLine, symb)
		symb = []string{}
	}

	ascLine = Aligner(ascLine, newLineCount)

	return ascLine
}
