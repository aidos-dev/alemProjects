package ascii

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func AsciiConv(textInp string, fontValue string) (string, error, int) {
	// Programm starts with checking if the command input is correct with IputError function
	err, errHttpCode := InputError()
	if err != nil {
		log.Print(err)
		return "", err, errHttpCode
	}
	// "splited" is a variable that gets strings slice from the PrepareStrings function.
	// actually it is a string from the command line that needs to be printed out in ascii-art form.
	// PrepareStrings function splits this string in slice of strings(chunks) delimiting it by a NewLine symbol so the
	// programm can handle them separatly

	splited, err, errHttpCode := PrepareStrings(textInp)
	if err != nil {
		// log.Print(err)
		return "", err, errHttpCode
	}

	// the for - range block bellow proceedes different cases of NewLines. All these conditions are required since we
	// don't know where a user would put a NewLine(\n)
	// Also it actually makes the final output(saving) to a file the final result(an argument converted to ascii-art format)

	// NewLineChecker function create new lines in the output if NewLine symbols(\n) is passed to the command line
	// the "empty" variable counts the number of empty strings (an empty string is a new line in this programm)
	// in order to allow to NewLineChecker function to change the "empty" variable it is passed to the function as an andess to this variable (as a pointer)

	empty := 0

	res := ""

	newLine := ""

	for _, el := range splited {

		newLine = NewLinesChecker(splited, el, &empty)
		res += newLine
		template, err, errHttpCode := FileReader(fontValue)
		if err != nil {
			log.Print(err)
			return "", err, errHttpCode
		}

		res += Printer(Converter(el, template))

	}

	return res, nil, http.StatusOK
}

// InputError function handles incorrect inputs in the command line
func InputError() (error, int) {
	// if len(os.Args) < 2 || len(os.Args) > 4 {
	// 	return errors.New("error: Wrong number of arguments\n\nUsage: go run . [STRING] [BANNER]\n\nExample: go run . something standard\n")
	// }

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
			return errors.New("error: Wrong template name. Please input one of 3 template names: standard, shadow or thinkertoy.\n\nUsage: go run . [STRING] [BANNER]\n\nExample: go run . something standard\n"), http.StatusBadRequest
		}
	}

	return nil, http.StatusOK
}

func PrepareStrings(textInp string) ([]string, error, int) {
	// "arg" is a variabel containing a string from the command line that needs to be printed out in ascii-art format
	arg := textInp

	// this replace is required to make all NewLine symbols in one common format.
	// Mainly it's required to proceed NewLines created in terminal.

	// arg = strings.ReplaceAll(arg, "\n", "\\n")

	// arg = strings.ReplaceAll(arg, "\r", "\\n")

	// // arg = strings.ReplaceAll(arg, "\r\n", "\\n")

	// splited := strings.Split(arg, "\\n")

	// web newlines section

	arg = strings.ReplaceAll(arg, "\r", "\\r")

	arg = strings.ReplaceAll(arg, "\n", "\\r")

	arg = strings.ReplaceAll(arg, "\r\n", "\\r")

	splited := strings.Split(arg, "\\r")

	// anyRune variable adds +1 if there is anything passed in the command line otherwise it remains 0
	anyRune := 0
	argRune := []rune(arg)
	for _, r := range argRune {
		// this condition makes an exeption for the check below.
		// 10 is Symbols of NewLine('\r), they are procided lower
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
			return splited, fmt.Errorf("<br></br>\nError: Please input coorect symbols:\n\n%w", iputErr), http.StatusBadRequest

		} else {
			anyRune++
		}
	}

	if anyRune == 0 {
		splited = splited[1:]
	}

	return splited, nil, http.StatusOK
}

func NewLinesChecker(splited []string, el string, empty *int) string {
	if el == "" {
		*empty++
	}
	if splited[0] == "" {
		if el == "" && *empty > 1 { // Example: "\n\n"
			// fmt.Println()
			return "\n"
		}
	} else {
		if el == "" {
			// fmt.Println()
			return "\n"
		}
	}
	// the condition bellow proceedes cases when user inputs NewLine(\n) in the beginnig and then types any other printable symbols
	// Example: "\nHello"
	if splited[0] == "" && splited[1] != "" && *empty == 1 {
		// fmt.Println()
		*empty++ // we increase empty in order to restrict addition of new lines by other conditions above
		return "\n"

	}
	if splited[0] == "" && splited[1] == "" && *empty >= 1 && el != "" {
		// fmt.Println()
		*empty++
		return "\n"

	}
	return ""
}

// Printer function prints the result in ascii-art format into the terminal
func Printer(ascLine [][]string) string {
	chunckRes := ""
	for row := 1; row < 9; row++ {
		if len(ascLine) == 0 {
			break
		}
		for sym := range ascLine {
			chunckRes += ascLine[sym][row]
		}
		// chunckRes += "\n"
		chunckRes += "\r"
	}

	return chunckRes
}

func FileReader(fontValue string) ([]string, error, int) {
	// these 3 variables below are hash representation of banner files
	// hash codes are used in this programm to protect template files from ammendment

	// fmt.Printf("%x\n", md5.Sum(data)) - This command was used to find out hash codes for template files and store them in 3 variables

	hashStandard := "ac85e83127e49ec42487f272d9b9db8b"
	hashShadow := "a49d5fcb0d5c59b2e77674aa3ab8bbb1"
	hashThinkertoy := "93e20c2510dfd28993db87352915826a"

	sourceName := "banners/standard.txt"

	sourceFind := fontValue

	switch sourceFind {
	case "shadow":
		sourceName = "webFiles/banners/shadow.txt"
	case "thinkertoy":
		sourceName = "webFiles/banners/thinkertoy.txt"
	case "standard":
		sourceName = "webFiles/banners/standard.txt"
	default:
		return []string{}, errors.New("Error: Please input correct banner name: standard, shadow or thinkertoy\n"), http.StatusInternalServerError

	}

	data, err := ioutil.ReadFile(sourceName)
	if err != nil {
		return []string{}, errors.New("failed reading data from  a banner file"), http.StatusInternalServerError // returns ampty slice of strings in case of error
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
		return []string{}, errors.New("\nError: Banner files must not be changed!\n\nPlease do not change banner files, use them only in original form.\n"), http.StatusInternalServerError
	}

	return lines, nil, http.StatusOK
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
