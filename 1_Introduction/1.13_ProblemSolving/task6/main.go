package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a > 0 && b > 0 {

		c := (float64(a) + float64(b)) / 2

		fmt.Print(c)
	}
}
