package main

import (
	"errors"
	"fmt"
	"math"
)

func hypotenuse(a, b int) (int, error) {
	if a != 0 || b != 0 {
		result := a*a + b*b
		return int(math.Sqrt(float64(result))), nil
	}

	return 0, errors.New("err")
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	result, err := hypotenuse(a, b)
	if err != nil {
		fmt.Print("ошибка")
	} else {
		fmt.Print(result)
	}
}
