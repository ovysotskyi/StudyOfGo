package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

func convertStr(str string) (string, error) {
	length := utf8.RuneCountInString(str)

	if length > 0 && length <= 1000 {
		newStr := strings.Split(str, "")
		result := strings.Join(newStr, "*")
		// or
		// for index, item := range str {
		// 	newStr = newStr + string(item)
		// 	if index != length-1 {
		// 		newStr = newStr + "*"
		// 	}
		// }

		return result, nil
	}

	return "", errors.New("err")
}

func main() {
	var str string
	fmt.Scan(&str)

	newStr, err := convertStr(str)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(newStr)
	}
}
