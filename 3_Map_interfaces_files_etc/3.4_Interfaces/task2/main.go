package main

import (
	"fmt"
	"strings"
)

type Battery struct {
	charge string
}

func (b Battery) String() string {
	ans := strings.Repeat("X", strings.Count(b.charge, "1"))
	return fmt.Sprintf("[%10v]", ans)
}

// or
// func (b Battery) String() string {
// 	size := "["

// 	units := 0
// 	for _, value := range b.charge {
// 		if value == rune('1') {
// 			units++
// 		}
// 	}

// 	for i := 0; i < 10-units; i++ {
// 		size += " "
// 	}
// 	for i := 0; i < units; i++ {
// 		size += "X"
// 	}

// 	size += "]"

// 	return size
// }

func main() {
	var batteryForTest Battery

	fmt.Scan(&batteryForTest.charge)

	fmt.Print(batteryForTest)
}
