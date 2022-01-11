package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str := []rune(text)

	// -3 because len - 1 - rune(\n); rune(\n) = 2 bite
	if unicode.IsUpper(str[0]) && str[utf8.RuneCountInString(text)-3] == '.' {
		fmt.Print("Right")
	} else {
		fmt.Print("Wrong")

	}
}
