package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	if number > 0 && number < 100 {
		if number != 11 && number%10 == 1 {
			fmt.Print(number, " korova")
		} else if (number%10 == 2 || number%10 == 3 || number%10 == 4) && (number < 10 || number > 20) {
			fmt.Print(number, " korovy")
		} else {
			fmt.Print(number, " korov")
		}
	}
}
