package main

import "fmt"

func main() {
	var size, number, count int
	fmt.Scan(&size)

	for i := 0; i < size; i++ {
		fmt.Scan(&number)

		if number == 0 {
			count++
		}
	}

	fmt.Print(count)
}
