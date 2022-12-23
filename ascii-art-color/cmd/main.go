package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"01.alem.school/git/Optimus/ascii-art-color/internals/ascii"
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
	splited, err := ascii.PrepareStrings()
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
			ascii.NewLinesChecker(splited, el, &empty)
			template, err := ascii.FileReader()
			if err != nil {
				log.Print(err)
				return
			}
			Printer(ascii.Converter111(el, template, &lineCounter))
		} else {
			ascii.NewLinesChecker(splited, el, &empty)
			template, err := ascii.FileReader()
			if err != nil {
				log.Print(err)
				return
			}
			Printer(ascii.ConverterLastWord(el, template))
		}
		lineCounter++
	}
	// colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, "Red")
	// orange := fmt.Sprintf("\x1b[0;33m%s", "Orange")
	// // fmt.Println(colored)
	// colorReset := fmt.Sprintf("\x1b[1;37m%s", "")

	// fmt.Println(orange, colorReset, "space test")

	// fmt.Println(ascii.ColorFlagCheck())
	// fmt.Println(ascii.ColorIndex())
	// fmt.Println(ascii.Colors())
	// fmt.Println(ascii.ColorIndex())

	// fmt.Println("color test")

	// wordPtr := flag.String("word", "foo", "a string")
	// flag.Parse()
	// fmt.Println("word:", *wordPtr)
}

// InputError function handles incorrect inputs in the command line
func InputError() error {
	// if len(os.Args) <= 3 || len(os.Args) > 4 {
	// 	// log.Fatalf("Error: Wrong number of arguments\n\nUsage: go run . [STRING] [BANNER] [OPTION]\n\nExample: go run . something standard --align=right\n")
	// 	return errors.New("error: Wrong number of arguments\n\nUsage: go run . [STRING] [BANNER] [OPTION]\n\nExample: go run . something standard --align=right\n")
	// }

	if len(os.Args) < 2 || len(os.Args) > 4 {
		// log.Fatalf("Error: Wrong number of arguments\n\nUsage: go run . [STRING] [BANNER] [OPTION]\n\nExample: go run . something standard --align=right\n")
		return errors.New("error: Wrong number of arguments\n\nUsage: go run . [STRING] [OPTION]\n\nExample: go run . something --color=<color>\n")
	}

	if len(os.Args) > 2 {
		flagError := ascii.ColorFlagErr()
		if flagError != nil {
			return flagError
		}
	}

	if len(os.Args) > 3 {
		indexFlagErr := ascii.ColorIndexErr()
		if indexFlagErr != nil {
			return indexFlagErr
		}
	}

	// if len(os.Args) > 2 {
	// 	templateError := true
	// 	switch os.Args[2] {
	// 	case "standard":
	// 		templateError = false
	// 	case "shadow":
	// 		templateError = false
	// 	case "thinkertoy":
	// 		templateError = false
	// 	}

	// 	if templateError {s
	// 		// log.Fatalf("Error: Wrong template name. Please input one of 3 template names: standard, shadow or thinkertoy.\n\nUsage: go run . [STRING] [BANNER] [OPTION]\n\nExample: go run . something standard --align=right\n")
	// 		return errors.New("error: Wrong template name. Please input one of 3 template names: standard, shadow or thinkertoy.\n\nUsage: go run . [STRING] [BANNER] [OPTION]\n\nExample: go run . something standard --align=right\n")
	// 	}
	// }

	// if len(os.Args) > 3 {

	// 	flagError := ascii.AlignFlagErr()
	// 	if flagError != nil {
	// 		return flagError
	// 	}
	// }
	return nil
}

// Printer function prints the result in ascii-art format into the terminal
func Printer(ascLine [][]string) {
	// ascLineCol := Coloriser(ascLine)

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
