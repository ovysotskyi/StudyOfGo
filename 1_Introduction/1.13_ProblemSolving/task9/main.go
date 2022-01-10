package main

import "fmt"

func main() {
	// var number, summ, dRoot int
	// fmt.Scan(&number)

	// if number > 0 && number < 10e6 {
	// 	for {
	// 		summ = summ + (number % 10)
	// 		number = number / 10

	// 		if number%10 == 0 {
	// 			break
	// 		}
	// 	}

	// 	for {
	// 		dRoot = dRoot + (summ % 10)
	// 		summ = summ / 10

	// 		if summ%10 == 0 {
	// 			break
	// 		}
	// 	}

	// 	fmt.Print(dRoot)
	// }
	//or https://codeforces.com/blog/entry/18286

	var number, dRoot int
	fmt.Scan(&number)

	dRoot = 1 + ((number - 1) % 9)

	fmt.Print(dRoot)

}
