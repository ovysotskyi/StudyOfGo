package main

import "fmt"

func main() {
	fmt.Print(minimumFromFour())
}

func minimumFromFour() int {
	var number, nMin int

	for i := 0; i < 4; i++ {

		fmt.Scan(&number)

		if nMin == 0 || nMin > number {
			nMin = number
		}
	}

	return nMin
}
