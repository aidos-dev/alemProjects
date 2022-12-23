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
	arg := strings.Join(os.Args[1:], " ")

	// anyRune variable adds +1 if there is anything passed in the command line otherwise it remains 0
	anyRune := 0
	for i, r := range arg {
		if r == 10 { // this condition makes an exeption for the check below. Symbols of NewLine('\n) are procided lower
			continue
		}

		if r < 32 || r > 126 { // this condition makes sure that only English letters and only valid symbols
			fmt.Println("Error") // passed into the command line otherwise it stopes the programm
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

	for _, el := range splited {
		if el == "" {
			empty++
		}
		if splited[0] == "" {
			if el == "" && empty > 1 { // Example: "\n\n"
				fmt.Println()
			}
		} else if el == "" {
			fmt.Println()
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

// Printer function prints the result in ascii-art format
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
	// the "hashStandard" variable below is a hash representation of "standard.txt" file
	// hash codes are used in this programm to protect a template file from ammendment

	// fmt.Printf("%x\n", md5.Sum(data)) - This command was used to find out hash code for template file and store it in a variable ("hashStandard")

	hashStandard := "ac85e83127e49ec42487f272d9b9db8b"
	data, err := ioutil.ReadFile("standard.txt")
	if err != nil {
		log.Fatalf("failed reading data from file: %s", err)
		return []string{} // returns ampty slice of strings in case of error
	}

	// the MD5 function checks and stores to "checksum" variable the current hash of the template file (standard.txt)
	checksum := MD5(string(data))
	hashErr := true
	if checksum == hashStandard {
		hashErr = false
	}
	// if the "hashErr" variable remains true it means that a template file has been ammended which is an Error and will lead to incorrect output
	if hashErr {
		fmt.Println("\nError: The banner file must not be changed!\n\nPlease do not change the banner file, use it only in original form.\n")
		os.Exit(1)
	}

	lines := strings.Split(string(data), "\n")

	return lines
}

// Converter function scans strings from the command line and convert them into 2D slice of slices of ascii-art symbols
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
