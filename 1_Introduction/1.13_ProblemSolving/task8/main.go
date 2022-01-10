package main

import "fmt"

func main() {
	var size, number, numMin, count int
	fmt.Scan(&size)

	for i := 0; i < size; i++ {
		fmt.Scan(&number)

		if numMin == 0 {
			numMin = number
		}

		if number < numMin {
			numMin = number
			count = 1
		} else if number == numMin {
			count++
		}
	}

	fmt.Print(count)
}
