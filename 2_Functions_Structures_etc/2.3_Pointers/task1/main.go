package main

import "fmt"

func main() {
	a := 2
	b := 2

	x1 := &a
	x2 := &b

	test(x1, x2)
}

func test(x1 *int, x2 *int) {
	y := *x1 * *x2

	fmt.Print(y)
}
