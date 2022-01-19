package main

import (
	"fmt"
)

type Battery struct {
	chargeSize string
}

func (b Battery) String() string {
	size := "["

	units := 0
	for _, value := range b.chargeSize {
		if value == rune('1') {
			units++
		}
	}

	for i := 0; i < 10-units; i++ {
		size += " "
	}
	for i := 0; i < units; i++ {
		size += "X"
	}

	size += "]"

	return size
}

func main() {
	var batteryForTest Battery

	fmt.Scan(&batteryForTest.chargeSize)

	fmt.Print(batteryForTest)
}
