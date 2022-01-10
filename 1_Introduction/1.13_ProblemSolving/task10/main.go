package main

import "fmt"

func main() {
	var a, b, number, count int
	fmt.Scan(&a, &b)

	if a <= b {
		for i := a; i <= b; i++ {
			if i%7 == 0 {
				number = i
				count++
			}
		}

		if count > 0 || (a == 0 && b == 0) {
			fmt.Print(number)
		} else {
			fmt.Print("NO")
		}
	}
}
