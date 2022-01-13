package main

import (
	"fmt"
)

func main() {
	var text string
	fmt.Scan(&text)

	var str []rune
	for index, rune := range []rune(text) {
		if index%2 != 0 {
			str = append(str, rune)
		}
	}

	fmt.Print(string(str))
}
