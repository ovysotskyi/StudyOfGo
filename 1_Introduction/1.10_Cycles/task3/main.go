package main

import "fmt"

func main() {
	var a, sum int

	fmt.Scan(&a)

	for i := 0; i < a; i++ {
		var num int
		fmt.Scan(&num)

		if (num > 9 && num < 100) && (num%8 == 0) {
			sum = sum + num
		}
	}

	fmt.Print(sum)
}
