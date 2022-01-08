package main

import (
	"fmt"
)

func main() {

	var num int
	fmt.Scan(&num)

	if num > 0 && num < 10 {
		fmt.Print(num)
	} else if num > 9 && num < 100 {
		fmt.Print(num / 10)
	} else if num > 99 && num < 1000 {
		fmt.Print(num / 100)
	} else if num > 999 && num < 10000 {
		fmt.Print(num / 1000)
	} else if num == 10000 {
		fmt.Print(num / 10000)
	}
}
