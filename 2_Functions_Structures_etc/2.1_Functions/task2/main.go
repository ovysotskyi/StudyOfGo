package main

import "fmt"

func main() {
	var number, numRem, digit, newNumber int64
	// var digit int
	fmt.Scan(&number, &digit)

	for {
		numRem = number % 10
		if numRem != digit {
			newNumber = (newNumber * 10) + numRem
		}

		number = number / 10
		if number == 0 {
			break
		}
	}

	for {
		numRem := newNumber % 10
		fmt.Printf("%d", numRem)

		newNumber = newNumber / 10
		if newNumber == 0 {
			break
		}
	}
}
