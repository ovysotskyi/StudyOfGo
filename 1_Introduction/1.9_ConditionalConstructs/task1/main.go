package main

import "fmt"

func main() {
	var a int

	fmt.Scan(&a)

	// if a < 0 {
	// 	fmt.Println("Число отрицательное")
	// } else if a > 0 {
	// 	fmt.Println("Число положительное")
	// } else {
	// 	fmt.Println("Ноль")
	// }
	// or
	switch {
	case a < 0:
		fmt.Println("Число отрицательное")
	case a > 0:
		fmt.Println("Число положительное")
	case a == 0:
		fmt.Println("Ноль")
	}
}
