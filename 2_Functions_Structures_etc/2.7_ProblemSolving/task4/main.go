package main

import (
	"fmt"
	"strconv"
	"strings"
)

func pow(str string) (string, error) {
	var newStr string
	strArr := strings.Split(str, "")
	for _, item := range strArr {
		number, err := strconv.Atoi(item)
		if err != nil {
			return "", err
		}

		number *= number
		newStr = newStr + strconv.Itoa(number)

	}

	return newStr, nil
}

func main() {
	var str string
	fmt.Scan(&str)

	max, err := pow(str)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(max)
	}
}
