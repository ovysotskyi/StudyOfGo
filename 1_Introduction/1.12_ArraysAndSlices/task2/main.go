package main

import "fmt"
func main() {
	var size, number int
	fmt.Scan(&size)

 	var slice []int

	for i := 0; i < size; i++ {
		fmt.Scan(&number)
		slice = append(slice, number)
	}

	fmt.Print(slice[3])
}
