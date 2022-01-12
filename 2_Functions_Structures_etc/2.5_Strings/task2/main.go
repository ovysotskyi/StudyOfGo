package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.Replace(text, " ", "", -1)

	str := []rune(text)
	var rihtRunes []rune

	//-2 because len - rune(\n); rune(\n) = 2 bite
	strSize := (utf8.RuneCountInString(text) - 2)

	if strSize%2 == 0 {
		right := str[strSize/2 : strSize]
		for i := len(right) - 1; i >= 0; i-- {
			rihtRunes = append(rihtRunes, right[i])
		}
	} else {
		right := str[strSize/2+1 : strSize]
		for i := len(right) - 1; i >= 0; i-- {
			rihtRunes = append(rihtRunes, right[i])
		}
	}

	if string(str[:strSize/2]) == string(rihtRunes) {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
