package main

import "fmt"

func main() {
	fmt.Print(sumInt(3, 2, 0))
}

func sumInt(numbers ...int) (int, int) {
	var counter, sum int

	for _, number := range numbers {
		counter++
		sum = sum + number
	}

	return counter, sum
}
