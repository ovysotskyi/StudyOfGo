package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func searchMax(str string) (string, error) {
	runes := []rune(str)
	length := utf8.RuneCountInString(str)

	if length > 0 && length <= 1000 {
		max := 0
		for _, item := range runes {
			if unicode.IsDigit(item) {
				if int(item) > max {
					max = int(item)
				}
			}
		}

		return strconv.Itoa(max), nil
	}

	return "", errors.New("err")
}

func main() {
	var str string
	fmt.Scan(&str)

	max, err := searchMax(str)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(max)
	}
}
