package main

import "fmt"

func main() {
	a := 4
	b := 2

	x1 := &a
	x2 := &b

	test(x1, x2)
}

func test(x1 *int, x2 *int) {
	*x1, *x2 = *x2, *x1

	fmt.Println(*x1, *x2)
}
