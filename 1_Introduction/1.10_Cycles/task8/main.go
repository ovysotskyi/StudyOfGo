package main

import "fmt"

func main() {
	var num1, num2, num3 int
	fmt.Scan(&num1, &num2)

	for num1 != 0 {
		if num1 >= 0 && num1 < 10000 {
			num1_temp := num1 % 10
			num1 = num1 / 10

			num2_new := num2
			for num2_new != 0 {
				num2_temp := num2_new % 10
				if num1_temp == num2_temp {
					num3 = (num3 * 10) + num1_temp
				}

				num2_new = num2_new / 10
			}
		}
	}

	for num3 != 0 {
		fmt.Print(num3%10, " ")
		num3 = num3 / 10
	}
}
