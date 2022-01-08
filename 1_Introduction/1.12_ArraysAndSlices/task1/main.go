package main

import "fmt"
func main() {
	var workArray [10]uint8
	var number uint8
	var index1, index2 int

	for index := range workArray {
		fmt.Scan(&number)

		workArray[index] = number
	}

	for i := 0; i < 3; i++ {
		fmt.Scan(&index1, &index2)

		workArray[index1], workArray[index2] = workArray[index2], workArray[index1]
	}

	for index := range workArray {
		fmt.Print(workArray[index], " ")
	}
}
