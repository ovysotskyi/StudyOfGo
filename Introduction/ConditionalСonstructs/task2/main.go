package main

import (
	"fmt"
)

func main() {

	var num, a, b, c int
	fmt.Scan(&num)

	c = num % 10
	b = (num - c) / 10 % 10
	a = num / 100

	if a != b && a != c && b != c {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}

}
