package ascii

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

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
			iputErr := errors.New("invalid symbol\n\n")
			return splited, fmt.Errorf("\nerror: Please input coorect string:\n\n%w", iputErr)

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

// Converter function scans strings from the command line and slices of strings from FileReader function.
// FileReader sends to Converter template strings for ascii-art symbols. Then Converter transforms them into 2D slice of
// slices of ascii-art symbols
func ConverterLastWord(s string, lines []string) (ascLine [][]string) {
	runA := []rune(s)
	ascLine = [][]string{}
	symb := []string{}

	// for _, r := range runA {
	// 	for i := 0; i < 10; i++ {
	// 		symb = append(symb, lines[int((r-32)*9)+i])
	// 	}
	// 	ascLine = append(ascLine, symb)
	// 	symb = []string{}
	// }

	for index := 0; index < len(runA); index++ {

		if len(os.Args) > 2 {
			symb = Coloriser(lines, symb, index, runA[index])
		}

		// this part is activated when Coloriser did not colorised a symbol to create a symbol
		// with defaul color
		if len(symb) == 0 {
			for i := 0; i < 10; i++ {
				symb = append(symb, lines[int((runA[index]-32)*9)+i])
			}
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

	for index := 0; index < len(runA); index++ {

		if len(os.Args) > 2 {
			symb = Coloriser(lines, symb, index, runA[index])
		}

		// this part is activated when Coloriser did not colorised a symbol to create a symbol
		// with defaul color
		if len(symb) == 0 {
			for i := 0; i < 10; i++ {
				symb = append(symb, lines[int((runA[index]-32)*9)+i])
			}
		}

		ascLine = append(ascLine, symb)
		symb = []string{}
	}

	// ascLine = Aligner(ascLine, newLineCount)

	return ascLine
}

func Coloriser(lines []string, symb []string, index int, r rune) []string {
	color := Colors()
	// symb := []string{}
	// ascLineCol := [][]string{}
	colorReset := "\033[0m"

	if len(os.Args) <= 3 {
		for i := 0; i < 10; i++ {
			symb = append(symb, fmt.Sprintf(color+"%s", lines[int((r-32)*9)+i]))
		}
	}

	if len(os.Args) > 3 {
		start, devidFlag, end := ColorIndex()

		// colorReset := fmt.Sprintf("\x1b[1;37m%s", "")
		// test1 := "test 111"
		// test2 := "test 222"

		// if r == ' ' {
		// 	index++
		// } else {
		if start == 0 && !devidFlag && len(os.Args[3]) == 8 { // expample: --index=
			for i := 0; i < 10; i++ {
				symb = append(symb, fmt.Sprintf(color+"%s"+colorReset, lines[int((r-32)*9)+i]))
			}
			// fmt.Println("test 1 ", devidFlag)
		}

		if index >= start && devidFlag && index <= end { // expample: --index=2:5
			for i := 0; i < 10; i++ {
				symb = append(symb, fmt.Sprintf(color+"%s"+colorReset, lines[int((r-32)*9)+i]))
			}
			// fmt.Println("test 2 ", devidFlag)
		}

		if index >= start && devidFlag && end == 0 { // expample: --index=2:
			for i := 0; i < 10; i++ {
				symb = append(symb, fmt.Sprintf(color+"%s"+colorReset, lines[int((r-32)*9)+i]))
			}
			// fmt.Println("test 3 ", devidFlag)
		}

		if index == start { // expample: --index=2
			for i := 0; i < 10; i++ {
				symb = append(symb, fmt.Sprintf(color+"%s"+colorReset, lines[int((r-32)*9)+i]))
				// symb = append(symb, colorReset)
				// fmt.Println("test 4.1 ")
				// symb = append(symb, test2)
			}
			// symb = append(symb, fmt.Sprintf("\x1b[1;37m%s", "")) // white color
			// symb = append(symb, colorReset)
			// symb = append(symb, test1)
			// fmt.Println("test 4 ", devidFlag)
		}

		if start == 0 && devidFlag && index <= end { // expample: --index=:3
			for i := 0; i < 10; i++ {
				symb = append(symb, fmt.Sprintf(color+"%s"+colorReset, lines[int((r-32)*9)+i]))
			}
			// fmt.Println("test 5 ", devidFlag)
		}
		// }
	}

	// fmt.Sprintf("\x1b[0;33m%s", ascLine[sym][row])
	return symb
}
