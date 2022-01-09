package main

import "fmt"

func main() {
	var size, number, counter int
	fmt.Scan(&size)

	var arr []int
	for i := 0; i < size; i++ {
		fmt.Scan(&number)
		arr = append(arr, number)
	}

	for _, number := range arr {
		if number > 0 {
			counter++
		}
	}

	fmt.Print(counter)
}
