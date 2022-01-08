package main

import "fmt"

func main() {
	for ;; {
		var number int
		fmt.Scan(&number)
	
		if number < 10 {
			continue
		} else if number > 100 {
			break
		} else {
			fmt.Println(number)
		}
	}
}
