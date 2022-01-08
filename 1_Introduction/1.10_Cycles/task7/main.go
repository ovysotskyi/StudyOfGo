package main

import "fmt"
func main() {
	var x, p, y int
	year := 0
	fmt.Scan(&x, &p, &y)

	for ;; {
		x = x + x * p/100
		year++
		if x >= y {
			fmt.Println(year)
			break
		}
		
	}
	
}
