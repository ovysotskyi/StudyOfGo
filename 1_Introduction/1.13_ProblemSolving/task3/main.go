package main

import "fmt"

func main() {
	var seconds, hours, minutes int
	fmt.Scan(&seconds)

	if seconds > 0 && seconds < 86399 {
		hours = seconds / 3600
		minutes = (seconds - (hours * 3600)) / 60

		fmt.Printf("It is %d hours %d minutes.", hours, minutes)
	}
}
