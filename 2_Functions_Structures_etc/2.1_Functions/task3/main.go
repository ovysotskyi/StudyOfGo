package main

import "fmt"

func main() {
	fmt.Print(vote(0, 1, 1))
}

func vote(x int, y int, z int) int {
	var number, counter int
	arr := [3]int{x, y, z}

	for _, item := range arr {
		var itemCounter int
		for _, numItem := range arr {
			if item == numItem {
				itemCounter++
			}
		}

		if itemCounter > counter {
			counter = itemCounter
			number = item
		}
	}

	return number
}

//or
// func vote(x int, y int, z int) int {
//     sum := x + y + z
//     if sum < 2 {
//         return 0
//     }
//     return 1
// }
