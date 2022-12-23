package ascii

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func FileReader() ([]string, error) {
	// these 3 variables below are hash representation of banner files
	// hash codes are used in this programm to protect template files from ammendment

	// fmt.Printf("%x\n", md5.Sum(data)) - This command was used to find out hash codes for template files and store them in 3 variables

	hashStandard := "ac85e83127e49ec42487f272d9b9db8b"
	hashShadow := "a49d5fcb0d5c59b2e77674aa3ab8bbb1"
	hashThinkertoy := "93e20c2510dfd28993db87352915826a"

	sourceName := "internals/banners/standard.txt"
	// sourceName := ""

	// if len(os.Args) > 2 {
	// 	// _, _, banner, _ := FlagCheck()
	// 	sourceFind := "internals/banners/standard.txt"

	// 	switch sourceFind {
	// 	case "shadow":
	// 		sourceName = "internals/banners/shadow.txt"
	// 	case "thinkertoy":
	// 		sourceName = "internals/banners/thinkertoy.txt"
	// 	case "standard":
	// 		sourceName = "internals/banners/standard.txt"
	// 	default:
	// 		return []string{}, errors.New("Error: Please input correct banner name: standard, shadow or thinkertoy\n")

	// 	}
	// }

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

// Package md5 implements the MD5 hash algorithm as defined in RFC 1321.
// Sum returns the MD5 checksum of the data.
func MD5(data string) string {
	h := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", h)
}
