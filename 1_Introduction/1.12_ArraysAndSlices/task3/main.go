package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	// здесь ваш код
	var aMax int
	for _, value := range array {
		if value > aMax || (aMax == 0 && value < 0) {
			aMax = value
		}
	}

	fmt.Print(aMax)
}
