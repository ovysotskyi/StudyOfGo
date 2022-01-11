package main

import "fmt"

func main() {
	fmt.Print(fibonacci(3))
}

func fibonacci(x int) int {
	fibNew, counter, fib, fibOld := 0, 1, 1, 1

	for {
		counter++
		fibNew = fibOld + fibNew
		fibOld = fib
		fib = fibNew

		if counter == x {
			break
		} else if x == 1 {
			fibNew = 1
			break
		}

	}

	return fibNew
}
