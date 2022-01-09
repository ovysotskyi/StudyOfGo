package main

import "fmt"

func main() {
	var number, newNumber int 
	fmt.Scan(&number)

	for i := 0; i < 3; i++ {
		newNumber = newNumber * 10 + number % 10
		number = number / 10
	}

	fmt.Print(newNumber)
}
