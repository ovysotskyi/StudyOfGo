package main

import (
	"fmt"
	"strings"
)

func main() {
	var text string
	fmt.Scan(&text)

	for _, letter := range text {
		if strings.Count(text, string(letter)) == 1 {
			fmt.Print(string(letter))
		}
	}

}
