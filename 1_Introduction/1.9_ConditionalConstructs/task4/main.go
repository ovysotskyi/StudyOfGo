package main

import (
	"fmt"
)

func main() {

	var num int
	fmt.Scan(&num)

	num1 := num / 1000
	num2 := num % 1000

	sum1 := (num1 % 10) + (num1 / 10 % 10) + (num1 / 100)
	sum2 := (num2 % 10) + (num2 / 10 % 10) + (num2 / 100)

	if sum1 == sum2 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
