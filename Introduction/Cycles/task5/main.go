package main

import "fmt"

func main() {
	var n, c, d int
	fmt.Scan(&n, &c, &d)
	max := 10000
	min := 1
	if (min <= n && n <= max) && (min <= c && c <= max) && (min <= d && d <= max) && (c != d) {
		for i := 0; i <= n; i++ {
			if i%c == 0 && i%d != 0 && c != d {
				fmt.Println(i)
				break
			}
		}
	}
}
