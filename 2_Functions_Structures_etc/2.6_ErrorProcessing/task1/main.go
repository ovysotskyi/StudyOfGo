package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b != 0 {
		return a / b, nil
	}

	return 0, errors.New("err")
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	result, err := divide(a, b)
	if err != nil {
		fmt.Print("ошибка")
	} else {
		fmt.Print(result)
	}
}
