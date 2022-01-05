package main

import (
	"fmt"
)

func main() {

	var a int
	fmt.Scan(&a)

	hours := a / 30
	minutes := 2 * (a % 30)
	fmt.Println("It is", hours, "hours ", minutes, "minutes.")
}
