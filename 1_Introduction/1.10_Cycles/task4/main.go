package main

import "fmt"

func main() {
	var count, maxNum int

	for {
		var num int

		fmt.Scan(&num)

		if num == 0 {
			break
		} else if num > maxNum {
			count = 1
			maxNum = num
		} else if num == maxNum {
			count++
		}
	}

	fmt.Print(count)
}

// or
// package main

// import "fmt"

// func main() {
// 	var count, maxNum, num int

// 	for fmt.Scan(&num); num != 0; fmt.Scan(&num) {
// 		if num > maxNum {
// 			count = 1
// 			maxNum = num
// 		} else if num == maxNum {
// 			count++
// 		}
// 	}

// 	fmt.Print(count)
// }
