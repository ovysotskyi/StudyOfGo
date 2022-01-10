package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	fibNew, counter, fib, fibOld := 0, 1, 1, 1

	for {
		counter++
		fibNew = fibOld + fibNew
		fibOld = fib
		fib = fibNew

		if number == fib {
			fmt.Print(counter)
			break
		} else if number < fib {
			fmt.Print(-1)
			break
		}

	}
}
