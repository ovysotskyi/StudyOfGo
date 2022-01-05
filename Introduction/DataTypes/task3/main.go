package main

import (
	"fmt"
	"math"
)

func main() {

	var a float64
	fmt.Scan(&a)

	// b := a * a
	//or
	b := math.Pow(a, 2)
	fmt.Println(b)
}
