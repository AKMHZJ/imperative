package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 2 {
		fmt.Println("go run main.go [input] [output]")
		os.Exit(1)
	}

	input := arg[0]
	output := arg[1]

	data, err := os.ReadFile(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	containt := strings.Split(string(data), " ")

	i := 0
	for i < len(containt) {
		current := containt[i]
		splittedWithNL := strings.Split(current, "\n")
		if len(splittedWithNL) > 1 {
			containt = Remove(containt, i)
			for j, v := range splittedWithNL {
				containt = AddAt(containt, i, v)
				i++
				if j < len(splittedWithNL)-1 {
					containt = AddAt(containt, i, "\n")
					i++
				}
			}
		} else {
			i++
		}
	}

	i = 0
	for i < len(containt) {
		current := containt[i]
		if len(current) > 1 {
			for len(containt[i]) > 0 && containt[i][0] == '\'' {
				containt = AddAt(containt, i, "'")
				i++
				containt[i] = containt[i][1:]
			}
			j := i
			for len(containt[j]) > 0 && containt[j][len(containt[j])-1] == '\'' {
				containt[j] = containt[j][:len(containt[j])-1]
				i++
				containt = AddAt(containt, i, "'")
			}
			i++
		} else {
			i++
		}
	}

	i = 0
	for i < len(containt) {
		current := containt[i]
		if current == "(bin)" {
			j := i - 1
			for j >= 0 && (containt[j] == "" || containt[j] == "\n") {
				j--
			}
			if j >= 0 {
				containt[j] = Bin(containt[j])
			}
			containt = Remove(containt, i)
		} else if current == "(hex)" {
			j := i - 1
			for j >= 0 && (containt[j] == "" || containt[j] == "\n") {
				j--
			}
			if j >= 0 {
				containt[j] = Hex(containt[j])
			}
			containt = Remove(containt, i)
		} else if current == "(up)" {
			j := i - 1
			for j >= 0 && (containt[j] == "" || containt[j] == "\n") {
				j--
			}
			if j >= 0 {
				containt[j] = strings.ToUpper(containt[j])
			}
			containt = Remove(containt, i)
		} else if current == "(low)" {
			j := i - 1
			for j >= 0 && (containt[j] == "" || containt[j] == "\n") {
				j--
			}
			if j >= 0 {
				containt[j] = strings.ToLower(containt[j])
			}
			containt = Remove(containt, i)
		} else if current == "(cap)" {
			j := i - 1
			for j >= 0 && containt[j] == "" {
				j--
			}
			if j >= 0 {
				containt[j] = Capitalize(containt[j])
			}
			containt = Remove(containt, i)
		} else if current == "(up," {
			if i+1 >= len(containt) {
				i++
				continue
			}
			Next := containt[i+1]
			re := regexp.MustCompile(`^((\+|-)?\d+)\)$`)
			match := re.FindAllStringSubmatchIndex(Next, 1)
			if len(match) > 0 {
				strat := match[0][2]
				end := match[0][3]
				count, _ := strconv.Atoi(Next[strat:end])
				j := i - 1
				for j >= 0 && count > 0 {
					if containt[j] != "" && containt[j] != "\n" {
						containt[j] = strings.ToUpper(containt[j])
						count--
					}
					j--
				}
			} else {
				i++
				continue
			}
			containt = Remove(containt, i)
			containt = Remove(containt, i)
		} else if current == "(low," {
			if i+1 >= len(containt) {
				i++
				continue
			}
			Next := containt[i+1]
			re := regexp.MustCompile(`^((\+|-)?\d+)\)$`)
			match := re.FindAllStringSubmatchIndex(Next, 1)
			if len(match) > 0 {
				strat := match[0][2]
				end := match[0][3]
				count, _ := strconv.Atoi(Next[strat:end])
				j := i - 1
				for j >= 0 && count > 0 {
					if containt[j] != "" && containt[j] != "\n" {
						containt[j] = strings.ToLower(containt[j])
						count--
					}
					j--
				}
			} else {
				i++
				continue
			}
			containt = Remove(containt, i)
			containt = Remove(containt, i)
		} else if current == "(cap," {
			if i+1 >= len(containt) {
				i++
				continue
			}
			Next := containt[i+1]
			re := regexp.MustCompile(`^((\+|-)?\d+)\)$`)
			match := re.FindAllStringSubmatchIndex(Next, 1)
			if len(match) > 0 {
				strat := match[0][2]
				end := match[0][3]
				count, _ := strconv.Atoi(Next[strat:end])
				j := i - 1
				for j >= 0 && count > 0 {
					if containt[j] != "" && containt[j] != "\n" {
						containt[j] = Capitalize(containt[j])
						count--
					}
					j--
				}
			} else {
				i++
				continue
			}
			containt = Remove(containt, i)
			containt = Remove(containt, i)
		} else {
			i++
		}
	}
	i = 0
	for i < len(containt) {
		current := containt[i]

		if strings.ToLower(current) == "a" {
			j := i + 1
			for j < len(containt) && containt[j] == "" {
				j++
			}
			if j < len(containt) && isVowel(containt[j][0]) {
				containt[i] = containt[i] + "n"
			}
		}
		i++
	}

	isFirstQoute := true
	i = 0
	for i < len(containt) {
		current := containt[i]
		if current == "'" && isFirstQoute {
			j := i + 1
			for j < len(containt) && (containt[j] == "" || containt[j] == "\n") {
				containt = Remove(containt, j)
			}
			i++
			if i < len(containt) {
				if containt[i] != "'" {
					isFirstQoute = false
				}
				containt[i] = "'" + containt[i]
				containt = Remove(containt, i-1)
			}
		} else if current == "'" && !isFirstQoute {
			isFirstQoute = true
			j := i - 1
			for i >= 0 && (containt[j] == " " || containt[j] == "\n") {
				containt = Remove(containt, j)
				i--
				j--
			}
			if i > 0 {

				containt[i-1] = containt[i-1] + "'"
				containt = Remove(containt, i)
				i--
			}
			i++
		} else {

			current := containt[i]
			re := regexp.MustCompile(`^([.,!?:;]+)`)
			match := re.FindAllStringSubmatchIndex(current, 1)
			if len(match) == 0 {
				i++
				continue
			}
			start := match[0][2]
			end := match[0][3]
			containt[i] = current[end:]
			j := i - 1
			for j >= 0 {
				current := containt[j]
				if current == "" {
					containt = Remove(containt, j)
					j--
					i--
				} else {
					break
				}
			}
			if j < 0 {
				containt = AddAt(containt, 0, current[start:end])
				j++
				i++
			} else {
				containt[j] = containt[j] + current[start:end]
			}
			j = j + 1
			for j < len(containt) && containt[j] == "" {
				containt = Remove(containt, j)
			}
		}
	}

	result := ""
	i = 0
	for i < len(containt) {
		current := containt[i]
		if i == len(containt)-1 {
			result += current
			i++
		} else if i+1 < len(containt) && containt[i+1] == "\n" {
			result += current
			i++
			result += "\n"
			i++
			if i < len(containt) {
				result += containt[i] + " "
				i++
			}

		} else if containt[i] == "\n" {
			result += "\n"
			i++
		} else {
			result += current + " "
			i++
		}
	}
	// result += "\n"
	errCreate := os.WriteFile(output, []byte(result), 0o644)
	if errCreate != nil {
		fmt.Println("Error")
		os.Exit(1)
	}
	// fmt.Println(result)
}

func Remove(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}

func Capitalize(s string) string {
	s1 := ""
	bol := true
	for _, valeu := range s {
		if (valeu >= 'a' && valeu <= 'z') || (valeu >= 'A' && valeu <= 'Z') || (valeu >= '0' && valeu <= '9') && valeu != '(' || valeu != ')' || valeu == 39 {
			if bol {
				if valeu >= 'a' && valeu <= 'z' {
					s1 += string(valeu - 32)
				} else {
					s1 += string(valeu)
				}
				bol = false
			} else {
				if valeu >= 'A' && valeu <= 'Z' {
					s1 += string(valeu + 32)
				} else {
					s1 += string(valeu)
				}
			}
		} else {
			s1 += string(valeu)
			bol = true
		}
	}
	return s1
}

func Bin(binStr string) string {
	dec, err := strconv.ParseInt(binStr, 2, 64)
	if err == nil {
		return fmt.Sprint(dec)
	}
	return binStr
}

func Hex(hexStr string) string {
	dec, err := strconv.ParseInt(hexStr, 16, 64)
	if err == nil {
		return fmt.Sprint(dec)
	}
	return hexStr
}

func AddAt(s []string, i int, value string) []string {
	if i == len(s) {
		return append(s, value)
	}
	new := append(s[:i+1], s[i:]...)
	new[i] = value
	return new
}

func isVowel(r byte) bool {
	switch unicode.ToLower(rune(r)) {
	case 'a', 'e', 'i', 'o', 'u', 'h':
		return true
	default:
		return false
	}
}
