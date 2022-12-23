package ascii

import (
	"errors"
	"os"
	"strings"
)

// ColorFlagErr function makes sure that there are no errors in flag input
func ColorFlagErr() error {
	if len(os.Args[2]) < 8 || os.Args[2][:8] != "--color=" {
		return errors.New("error: Wrong flag. Please input the correct flag: --color=\n\nUsage: go run . [STRING] OPTION]\n\nExample: go run cmd/main.go something --color=red\n")
	}

	if Colors() == "" {
		return errors.New("error: please input one of flag names: gray, red, lime, yellow, dodger blue, magenta, cyan, white, black, brown, green, orange, blue, purple, light sea green, silver\n\nUsage: go run . [STRING] [OPTION]\n\nExample: go run . something --color=red\n")
	}

	return nil
}

func ColorFlagCheck() string {
	inputColor := "white" // default value for the "inputColor" variable is "white"

	if os.Args[2][8:] != "" {
		inputColor = os.Args[2][8:] // this line checks for the flag parameter from a user and stores to the inputColor variable
	}

	// colorInp, index, _, _ := FlagCheck()

	return inputColor
}

func Colors() string {
	color := ColorFlagCheck()

	// convertion to lower case is provided here since a user might intput a desired color in upper or lower case.
	// Therefore it is required to bring an iput to the standard case
	color = strings.ToLower(color)
	// removing all white spaces is required since a user might input a name of a color with or without white spaces.
	// Therefore it is required to bring an iput to one standard
	color = strings.ReplaceAll(color, " ", "")

	// ANSI Color codes
	colorCode := ""
	const (
		GrayCol          = "\x1b[1;30m"
		RedCol           = "\x1b[1;31m"
		LimeCol          = "\x1b[1;32m"
		YellowCol        = "\x1b[1;33m"
		DodgerBlueCol    = "\x1b[1;34m"
		MagentaCol       = "\x1b[1;35m"
		CyanCol          = "\x1b[1;36m"
		WhiteCol         = "\x1b[1;37m"
		BlackCol         = "\x1b[0;30m"
		BrownCol         = "\x1b[0;31m"
		GreenCol         = "\x1b[0;32m"
		OrangeCol        = "\x1b[0;33m"
		BlueCol          = "\x1b[0;34m"
		PurpleCol        = "\x1b[0;35m"
		LightSeaGreenCol = "\x1b[0;36m"
		SilverCol        = "\x1b[0;37m"
	)

	switch color {
	case "gray":
		colorCode = GrayCol
	case "red":
		colorCode = RedCol
	case "lime":
		colorCode = LimeCol
	case "yellow":
		colorCode = YellowCol
	case "dodgerblue":
		colorCode = DodgerBlueCol
	case "magenta":
		colorCode = MagentaCol
	case "cyan":
		colorCode = CyanCol
	case "white":
		colorCode = WhiteCol
	case "black":
		colorCode = BlackCol
	case "brown":
		colorCode = BrownCol
	case "green":
		colorCode = GreenCol
	case "orange":
		colorCode = OrangeCol
	case "blue":
		colorCode = BlueCol
	case "purple":
		colorCode = PurpleCol
	case "lightseagreen":
		colorCode = LightSeaGreenCol
	case "silver":
		colorCode = SilverCol
	}

	return colorCode
}
