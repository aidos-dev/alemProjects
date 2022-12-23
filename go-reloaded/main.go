package main

import (
	"fmt"
	"os"
	"strconv"
)

// Завершает если нету ошибок
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func compare(a, b string) int {
	if a == b {
	} else if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// Преобразовние слова в верхний регистр
func upper(s string) string {
	doom := []rune(s)
	for i := 0; i < len(doom); i++ {
		letter := doom[i]
		if letter >= 'a' && letter <= 'z' {
			doom[i] = doom[i] - 32
		}
	}
	return string(doom)
}

// Преобразование слова в нижний регистр
func lower(s string) string {
	doom := []rune(s)
	for i := 0; i < len(doom); i++ {
		letter := doom[i]
		if letter >= 'A' && letter <= 'Z' {
			doom[i] = doom[i] + 32
		}
	}
	return string(doom)
}

// Первая РУНА строки
func frune(s string) string {
	ls := []rune(s)
	return string(ls[0])
}

// Первый символ слова
func capitalise(s string) string {
	runes := []rune(s)

	strlen := 0
	for i := range runes {
		strlen = i + 1
	}

	for i := 0; i < strlen; i++ {
		if i != 0 && (!((runes[i-1] >= 'a' && runes[i-1] <= 'z') || (runes[i-1] >= 'A' && runes[i-1] <= 'Z'))) {
			if runes[i] >= 'a' && runes[i] <= 'z' {
				runes[i] = rune(runes[i] - 32)
			}
		} else if i == 0 {
			if runes[i] >= 'a' && runes[i] <= 'z' {
				runes[i] = rune(runes[i] - 32)
			}
		} else {
			if runes[i] >= 'A' && runes[i] <= 'Z' {
				runes[i] = rune(runes[i] + 32)
			}
		}
	}
	return string(runes)
}

// Строка через пробелы
func split_white_spaces(s string) []string {
	var str []string
	var word string
	l := len(s) - 1

	for i, v := range s {
		if i == l {
			word = word + string(v)
			str = append(str, word)
		} else if v == 32 || v == 15 || v == 10 {
			if s[i+1] == ' ' || s[i+1] == '	' || s[i+1] == 10 {
			} else {
				str = append(str, word)
				word = ""
			}
		} else {
			word = word + string(v)
		}
	}
	return str
}

func quotes(s string) string {
	str := ""
	var removeSpace bool
	for i, char := range s {
		if char == 39 && s[i-1] == ' ' {
			if removeSpace {
				str = str[:len(str)-1]
				str = str + string(char)
				removeSpace = false
			} else {
				str = str + string(char)
				removeSpace = true
			}
		} else if i > 1 && s[i-2] == 39 && s[i-1] == ' ' {
			if removeSpace {
				str = str[:len(str)-1]
				str = str + string(char)
			} else {
				str = str + string(char)
			}
		} else {
			str = str + string(char)
		}
	}
	return str
}

func remove_tags(s []string) string {
	str := ""

	for i, tag := range s {
		if tag == "(cap," || tag == "(low," || tag == "(up," {
			s[i] = ""
			s[i+1] = ""
		} else if tag != "(up)" && tag != "(hex)" && tag != "(bin)" && tag != "(cap)" && tag != "(low)" && tag != "" {
			if i == 0 {
				str = str + tag
			} else {
				str = str + " " + tag
			}
		}
	}
	return str
}

func remove_spaces(s string) string {
	len := len(s) - 1
	if s[len-1] == ' ' {
		return remove_spaces(s[:len])
	}
	return s[:len]
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("You should pass 2 arguments")
		return
	}
	GoReloaded()
}

// Читает файлы и сохраняет их в базе данных
func GoReloaded() {
	data, err := os.ReadFile(os.Args[1])
	check(err)
	input := string(data)
	result := split_white_spaces(input)

	// запускает цикл диапазона
	for i, v := range result {
		// Заменяет слово перед его десятичной версей
		if compare(v, "(hex)") == 0 {
			if i-1 < 0 {
				return
			}
			j, _ := strconv.ParseInt(result[i-1], 16, 64)
			result[i-1] = fmt.Sprint(j)
		}

		if compare(v, "(bin)") == 0 {
			if i-1 < 0 {
				return
			}
			j, _ := strconv.ParseInt(result[i-1], 2, 64)
			result[i-1] = fmt.Sprint(j)
		}
		// Конвертирует слово перед нижним регистром
		if compare(v, "(low)") == 0 {
			if i-1 < 0 {
				return
			}
			result[i-1] = lower(result[i-1])
		}
		// конфертирует предыдущие слова в нижний регистр
		if compare(v, "(low,") == 0 {
			result[i-1] = lower(result[i-1])

			le := len(result[i+1])
			numb := result[i+1][:le-1]
			nu, err := strconv.Atoi(numb)
			check(err)

			for j := 1; j <= nu; j++ {
				if i-j < 0 {
					break
				}
				result[i-j] = lower(result[i-j])
			}
		}
		// ковертирует слова перед в верхний регистр
		if compare(v, "(up)") == 0 {
			if i-1 < 0 {
				return
			}
			result[i-1] = upper(result[i-1])
		}

		if compare(v, "(up,") == 0 {
			result[i-1] = upper(result[i-1])

			le := len(result[i+1])
			numb := result[i+1][:le-1]
			nu, err := strconv.Atoi(numb)
			check(err)

			for j := 1; j <= nu; j++ {
				if i-j < 0 {
					break
				}
				result[i-j] = upper(result[i-j])
			}
		}
		// Заглавными буквами пишется слова перед
		if compare(v, "(cap)") == 0 {
			if i-1 < 0 {
				return
			}
			result[i-1] = capitalise(result[i-1])
		}
		// Заглавными буквами указывается количество слов перед
		if compare(v, "(cap,") == 0 {
			result[i-1] = capitalise(result[i-1])

			le := len(result[i+1])
			numb := result[i+1][:le-1]
			nu, err := strconv.Atoi(numb)
			check(err)

			for j := 1; j <= nu; j++ {
				if i-j < 0 {
					break
				}
				result[i-j] = capitalise(result[i-j])
			}
		}
		// Преобразует "a" в "an", когда следующее слово начинается с гласной или "h".
		if compare(v, "a") == 0 && frune(result[i+1]) == "a" {
			result[i] = "an"
		}
		if compare(v, "a") == 0 && frune(result[i+1]) == "e" {
			result[i] = "an"
		}
		if compare(v, "a") == 0 && frune(result[i+1]) == "i" {
			result[i] = "an"
		}
		if compare(v, "a") == 0 && frune(result[i+1]) == "o" {
			result[i] = "an"
		}
		if compare(v, "a") == 0 && frune(result[i+1]) == "u" {
			result[i] = "an"
		}
		if compare(v, "a") == 0 && frune(result[i+1]) == "h" {
			result[i] = "an"
		}
	}

	// Вызывает remove_tags() и split_white_spaces() и получает новую переменную результата
	notagResult := remove_tags(result)
	result2 := split_white_spaces(notagResult)

	str := ""
	for _, word := range result2 {
		str = str + word + " "
	}
	// Удалить пробелы из строки
	str = remove_spaces(str)

	word := ""
	for i, char := range str {
		if i == len(str)-1 {
			if char == '.' || char == ',' || char == '!' || char == '?' || char == ';' || char == ':' {
				if str[i-1] == ' ' {
					word = word[:len(word)-1]
					word = word + string(char)
				} else {
					word = word + string(char)
				}
			} else {
				word = word + string(char)
			}
		} else if char == '.' || char == ',' || char == '!' || char == '?' || char == ';' || char == ':' {
			if str[i-1] == ' ' {
				word = word[:len(word)-1]
				word = word + string(char)
			} else {
				word = word + string(char)
			}
			if str[i+1] != ' ' && str[i+1] != '.' && str[i+1] != ',' && str[i+1] != '!' && str[i+1] != '?' && str[i+1] != ';' && str[i+1] != ':' {
				word = word + " "
			}
		} else {
			word = word + string(char)
		}
	}
	// Вызывает кавычки() для удаления дополнительных пробелов
	word = quotes(word)
	if len(os.Args) < 3 {
		fmt.Print(word)
		return
	}
	output := []byte(word)
	OurData := os.Args[2]
	words := os.WriteFile(OurData, output, 0644)
	check(words)
}
