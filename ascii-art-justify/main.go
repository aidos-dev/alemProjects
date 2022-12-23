package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	// Programm starts with checking if the command input is correct with IputError function
	err := InputError()
	if err != nil {
		log.Print(err)
		return
	}
	// "splited" is a variable that gets strings slice from the PrepareStrings function.
	// actually it is a string from the command line that needs to be printed out in ascii-art form.
	// PrepareStrings function splits this string in slice of strings(chunks) delimiting it by a NewLine symbol so the
	// programm can handle them separatly
	splited, err := PrepareStrings()
	if err != nil {
		log.Print(err)
		return
	}

	// the for - range block bellow proceedes different cases of NewLines. All these conditions are required since we
	// don't know where a user would put a NewLine(\n)
	// Also it actually makes the final output(saving) to a file the final result(an argument converted to ascii-art format)

	// NewLineChecker function create new lines in the output if NewLine symbols(\n) is passed to the command line
	// the "empty" variable counts the number of empty strings (an empty string is a new line in this programm)
	// in order to allow to NewLineChecker function to change the "empty" variable it is passed to the function as an andess to this variable (as a pointer)

	empty := 0
	lineCounter := 0

	for i, el := range splited {
		// TrimSpace function cuts(trims) all spaces in the beginning and the end of a string form the command line.
		// It is required to do clear alignment with no any extra spaces on both sides of the terminal window
		el = strings.TrimSpace(el)
		if i != len(splited) {
			NewLinesChecker(splited, el, &empty)
			template, err := FileReader()
			if err != nil {
				log.Print(err)
				return
			}
			Printer(Converter111(el, template, &lineCounter))
		} else {
			NewLinesChecker(splited, el, &empty)
			template, err := FileReader()
			if err != nil {
				log.Print(err)
				return
			}
			Printer(ConverterLastWord(el, template))
		}
		lineCounter++
	}
}

// InputError function handles incorrect inputs in the command line
func InputError() error {
	// if len(os.Args) <= 3 || len(os.Args) > 4 {
	// 	// log.Fatalf("Error: Wrong number of arguments\n\nUsage: go run . [STRING] [BANNER] [OPTION]\n\nExample: go run . something standard --align=right\n")
	// 	return errors.New("error: Wrong number of arguments\n\nUsage: go run . [STRING] [BANNER] [OPTION]\n\nExample: go run . something standard --align=right\n")
	// }

	if len(os.Args) < 2 || len(os.Args) > 4 {
		// log.Fatalf("Error: Wrong number of arguments\n\nUsage: go run . [STRING] [BANNER] [OPTION]\n\nExample: go run . something standard --align=right\n")
		return errors.New("error: Wrong number of arguments\n\nUsage: go run . [STRING] [BANNER] [OPTION]\n\nExample: go run . something standard --align=right\n")
	}

	if len(os.Args) > 2 {
		templateError := true
		switch os.Args[2] {
		case "standard":
			templateError = false
		case "shadow":
			templateError = false
		case "thinkertoy":
			templateError = false
		}

		if templateError {
			// log.Fatalf("Error: Wrong template name. Please input one of 3 template names: standard, shadow or thinkertoy.\n\nUsage: go run . [STRING] [BANNER] [OPTION]\n\nExample: go run . something standard --align=right\n")
			return errors.New("error: Wrong template name. Please input one of 3 template names: standard, shadow or thinkertoy.\n\nUsage: go run . [STRING] [BANNER] [OPTION]\n\nExample: go run . something standard --align=right\n")
		}
	}

	if len(os.Args) > 3 {

		flagError := FlagErr()
		if flagError != nil {
			return flagError
		}
	}
	return nil
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

func PrepareStrings() ([]string, error) {
	// "arg" is a variabel containing a string from the command line that needs to be printed out in ascii-art format
	arg := os.Args[1]

	// this replace is required to make all NewLine symbols in one common format.
	// Mainly it's required to proceed NewLines created in terminal.
	arg = strings.ReplaceAll(arg, "\n", "\\n")

	splited := strings.Split(arg, "\\n")

	// anyRune variable adds +1 if there is anything passed in the command line otherwise it remains 0
	anyRune := 0
	for _, r := range arg {
		// this condition makes an exeption for the check below.
		// 10 is Symbols of NewLine('\n), they are procided lower
		if r == 10 {
			continue
		}
		// this condition makes sure that only English letters and only valid symbols
		// passed into the command line otherwise it stops the programm
		// numbers 32 and 126 are ascii table rune numbers
		if r < 32 || r > 126 {
			// log.Fatalf("Error")
			// fmt.Println("Error: Incorrect", "symbol #", i+1)
			// log.Fatalf("/ Please input coorect string")
			iputErr := errors.New("Invalid symbol\n\n")
			return splited, fmt.Errorf("\nError: Please input coorect string:\n\n%w", iputErr)

		} else {
			anyRune++
		}
	}

	if anyRune == 0 {
		splited = splited[1:]
	}

	return splited, nil
}

func NewLinesChecker(splited []string, el string, empty *int) {
	if el == "" {
		*empty++
	}
	if splited[0] == "" {
		if el == "" && *empty > 1 { // Example: "\n\n"
			fmt.Println()
		}
	} else {
		if el == "" {
			fmt.Println()
		}
	}
	// the condition bellow proceedes cases when user inputs NewLine(\n) in the beginnig and then types any other printable symbols
	// Example: "\nHello"
	if splited[0] == "" && splited[1] != "" && *empty == 1 {
		fmt.Println()
		*empty++ // we increase empty in order to restrict addition of new lines by other conditions above
	}
	if splited[0] == "" && splited[1] == "" && *empty >= 1 && el != "" {
		fmt.Println()
		*empty++
	}
}

// Printer function prints the result in ascii-art format into the terminal
func Printer(ascLine [][]string) {
	for row := 1; row < 9; row++ {
		if len(ascLine) == 0 {
			break
		}
		for sym := range ascLine {
			fmt.Printf("%v", ascLine[sym][row])
		}
		fmt.Println()
	}
}

func FileReader() ([]string, error) {
	// these 3 variables below are hash representation of banner files
	// hash codes are used in this programm to protect template files from ammendment

	// fmt.Printf("%x\n", md5.Sum(data)) - This command was used to find out hash codes for template files and store them in 3 variables

	hashStandard := "ac85e83127e49ec42487f272d9b9db8b"
	hashShadow := "a49d5fcb0d5c59b2e77674aa3ab8bbb1"
	hashThinkertoy := "93e20c2510dfd28993db87352915826a"

	sourceName := "standard.txt"

	if len(os.Args) > 2 {
		sourceFind := os.Args[2]

		switch sourceFind {
		case "shadow":
			sourceName = "shadow.txt"
		case "thinkertoy":
			sourceName = "thinkertoy.txt"
		case "standard":
			sourceName = "standard.txt"
		default:
			return []string{}, errors.New("Error: Please input correct banner name: standard, shadow or thinkertoy\n")

		}
	}

	data, err := ioutil.ReadFile(sourceName)
	if err != nil {
		return []string{}, errors.New("failed reading data from file: %s") // returns ampty slice of strings in case of error
	}
	// the MD5 function checks and stores to checksum variable the current hash if the file (chosen by user)
	checksum := MD5(string(data))
	hashErr := true
	switch checksum {
	case hashStandard:
		hashErr = false
	case hashShadow:
		hashErr = false
	case hashThinkertoy:
		hashErr = false
	}

	lines := strings.Split(string(data), "\n")

	// if the "hashErr" variable remains true it means that one ore more template files has been ammended which is an Error and will lead to incorrect output
	if hashErr {
		return []string{}, errors.New("\nError: Banner files must not be changed!\n\nPlease do not change banner files, use them only in original form.\n")
	}

	return lines, nil
}

// Converter function scans strings from the command line and slices of strings from FileReader function.
// FileReader sends to Converter template strings for ascii-art symbols. Then Converter transforms them into 2D slice of
// slices of ascii-art symbols
func ConverterLastWord(s string, lines []string) (ascLine [][]string) {
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

	return ascLine
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

// Package md5 implements the MD5 hash algorithm as defined in RFC 1321.
// Sum returns the MD5 checksum of the data.
func MD5(data string) string {
	h := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", h)
}

// func consoleSize() (int, int) {
// 	cmd := exec.Command("stty", "size")
// 	cmd.Stdin = os.Stdin
// 	out, err := cmd.Output()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	s := string(out)
// 	s = strings.TrimSpace(s)
// 	sArr := strings.Split(s, " ")

// 	heigth, err := strconv.Atoi(sArr[0])
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	width, err := strconv.Atoi(sArr[1])
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return heigth, width
// }
