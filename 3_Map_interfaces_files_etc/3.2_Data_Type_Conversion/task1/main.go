package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func strToNumber(str string) (int64, error) {
	runes := []rune(str)
	var newStr string

	for _, item := range runes {
		if unicode.IsDigit(item) {
			newStr += string(item)
		}
	}

	return strconv.ParseInt(newStr, 10, 64)
}

func adding(str1, str2 string) int64 {
	num1, _ := strToNumber(str1)
	num2, _ := strToNumber(str2)

	return num1 + num2
}

func main() {
	var str1, str2 string
	fmt.Scan(&str1, &str2)

	fmt.Print(adding(str1, str2))
}
