package main

import "fmt"

func main() {
	var a, b, c, aPow, bPow, cPow int
	fmt.Scan(&a, &b, &c)

	if c > a && c > b {
		aPow = a * a
		bPow = b * b
		cPow = c * c
	}

	if (aPow + bPow) == cPow {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}
}
