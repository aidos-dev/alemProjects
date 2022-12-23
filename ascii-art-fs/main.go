package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Error: Wrong number of arguments\n\nUsage: go run . [STRING] [BANNER]\n\nExample: go run . something standard\n")
		return
	}

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
		fmt.Printf("Error: Wrong template name. Please input one of 3 template names: standard, shadow or thinkertoy.\n\nUsage: go run . [STRING] [BANNER]\n\nExample: go run . something standard\n")
		return
	}

	// arg := strings.Join(os.Args[1:], " ")
	arg := os.Args[1]

	// anyRune variable adds +1 if there is anything passed in the command line otherwise it remains 0
	anyRune := 0
	for i, r := range arg {
		// this condition makes an exeption for the check below.
		// 10 is Symbols of NewLine('\n), they are procided lower
		if r == 10 {
			continue
		}
		// this condition makes sure that only English letters and only valid symbols
		// passed into the command line otherwise it stops the programm
		// numbers 32 and 126 are ascii table rune numbers
		if r < 32 || r > 126 {
			fmt.Println("Error")
			fmt.Println("Incorrect", "symbol #", i+1)
			fmt.Println("Please input coorect string")
			return
		} else {
			anyRune++
		}
	}

	// this replace is required to make all NewLine symbols in one common format.
	// Mainly it's required to proceed NewLines created in terminal.
	arg = strings.ReplaceAll(arg, "\n", "\\n")

	splited := strings.Split(arg, "\\n")

	empty := 0

	if anyRune == 0 {
		splited = splited[1:]
	}

	// the for - range block bellow proceedes different cases of NewLines. All these conditions are required since we
	// don't know where a user would put a NewLine(\n)
	// Also it actually makes the final output(saving) to a file the final result(an argument converted to ascii-art format)

	for _, el := range splited {
		if el == "" {
			empty++
		}
		if splited[0] == "" {
			if el == "" && empty > 1 { // Example: "\n\n"
				fmt.Println()
			}
		} else {
			if el == "" {
				fmt.Println()
			}
		}
		// the condition bellow proceedes cases when user inputs NewLine(\n) in the beginnig and then types any other printable symbols
		// Example: "\nHello"
		if splited[0] == "" && splited[1] != "" && empty == 1 {
			fmt.Println()
			empty++ // we increase empty in order to restrict addition of new lines by other conditions above
		}

		Printer(Converter(el, FileReader()))

	}
}

// Printer function prints the result in ascii-art format into the file created by FileWriter function
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

func FileReader() []string {
	// these 3 variables below are hash representation of banner files
	// hash codes are used in this programm to protect template files from ammendment

	// fmt.Printf("%x\n", md5.Sum(data)) - This command was used to find out hash codes for template files and store them in 3 variables

	hashStandard := "ac85e83127e49ec42487f272d9b9db8b"
	hashShadow := "a49d5fcb0d5c59b2e77674aa3ab8bbb1"
	hashThinkertoy := "93e20c2510dfd28993db87352915826a"

	sourceFind := os.Args[2]
	sourceName := ""

	switch sourceFind {
	case "shadow":
		sourceName = "shadow.txt"
	case "thinkertoy":
		sourceName = "thinkertoy.txt"
	case "standard":
		sourceName = "standard.txt"
	default:
		fmt.Printf("Error: Please input correct banner name: standard, shadow or thinkertoy\n")
		log.Fatal()
	}

	data, err := ioutil.ReadFile(sourceName)
	if err != nil {
		log.Fatalf("failed reading data from file: %s", err)
		return []string{} // returns ampty slice of strings in case of error
	}
	// the MD5 function checks and stores to checksum variable the current hash of the file (chosen by user)
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
	// if the "hashErr" variable remains true it means that one ore more template files has been ammended which is an Error and will lead to incorrect output
	if hashErr {
		fmt.Println("\nError: Banner files must not be changed!\n\nPlease do not change banner files, use them only in original form.\n")
		os.Exit(1)
	}

	lines := strings.Split(string(data), "\n")

	return lines
}

// Converter function scans strings from the command line and slices of strings from FileReader function.
// FileReader sends to Converter template strings for ascii-art symbols. Then Converter transforms them into 2D slice of
// slices of ascii-art symbols
func Converter(s string, lines []string) (ascLine [][]string) {
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

// Package md5 implements the MD5 hash algorithm as defined in RFC 1321.
// Sum returns the MD5 checksum of the data.
func MD5(data string) string {
	h := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", h)
}
