package main

import "fmt"

func main() {
	var number, summ int 
	fmt.Scan(&number)

	for i := 0; i < 3; i++ {
		summ += number % 10
		number = number / 10
	}

	fmt.Print(summ)
}
