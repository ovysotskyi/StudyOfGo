package main

import (
	"fmt"
	"strings"
)

func main() {
	var text, text2 string

	fmt.Scan(&text, &text2)

	fmt.Print(strings.Index(text, text2))

	// or
	// result := -1

	// //-2 because len - rune(\n); rune(\n) = 2 bite
	// len1 := utf8.RuneCountInString(text)
	// len2 := utf8.RuneCountInString(text2)

	// for index := range text {
	// 	if index+len2 > len1 {
	// 		break
	// 	}

	// 	str := text[index : index+len2]
	// 	if str == text2[:len2] {
	// 		result = index
	// 		break
	// 	}
	// }

	// fmt.Print(result)
}
