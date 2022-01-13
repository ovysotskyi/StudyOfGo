package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	var password string
	fmt.Scan(&password)

	passwordRunes := []rune(password)

	if utf8.RuneCountInString(password) >= 5 {
		letters := 0
		digits := 0

		for _, item := range passwordRunes {
			if unicode.Is(unicode.Latin, item) && unicode.IsLetter(item) {
				letters++
			} else if unicode.IsDigit(item) {
				digits++
			} else {
				letters = 0
				digits = 0
				break
			}
		}

		if letters > 0 && digits > 0 {
			fmt.Print("Ok")
		} else {
			fmt.Println("Wrong password")
		}

	} else {
		fmt.Println("Wrong password")
	}

}
