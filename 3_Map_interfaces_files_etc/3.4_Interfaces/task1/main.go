package main

import (
	"fmt"
)

func main() {
	var value1, value2, operation interface{}

	var flt1 float64
	var flt2 int
	flt1 = 12.7654
	flt2 = 1
	op := "-"

	value1, value2 = flt1, flt2
	operation = op

	if _, ok := value1.(float64); !ok {
		fmt.Printf("value=%v: %T", value1, value1)
		return
	}

	if _, ok := value2.(float64); !ok {
		fmt.Printf("value=%v: %T", value2, value2)
		return
	}

	if _, ok := operation.(string); !ok {
		fmt.Printf("неизвестная операция")
		return
	}

	switch operation {
	case "+":
		fmt.Printf("%.4f", value1.(float64)+value2.(float64))
	case "-":
		fmt.Printf("%.4f", value1.(float64)-value2.(float64))
	case "/":
		fmt.Printf("%.4f", value1.(float64)/value2.(float64))
	case "*":
		fmt.Printf("%.4f", value1.(float64)*value2.(float64))
	default:
		fmt.Printf("неизвестная операция")
	}
}
