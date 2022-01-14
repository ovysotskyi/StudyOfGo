package main

import (
	"fmt"
	"math"
)

func M(p, v float64) float64 {
	return p * v
}

func W(k, p, v float64) float64 {
	return math.Sqrt(k / M(p, v))
}

func T(k, p, v float64) float64 {
	return 6 / W(k, p, v)
}

func main() {
	var k, p, v float64
	fmt.Scan(&k, &p, &v)

	fmt.Print(T(k, p, v))

}
