package main

import "fmt"

func main() {
	var size, number int
	fmt.Scan(&size)

	var arr []int
	for i := 0; i < size; i++ {
		fmt.Scan(&number)
		arr = append(arr, number)
	}

	for index, number := range arr {
		if index % 2 == 0 {
			fmt.Print(number, " ")
		}
	}
}
