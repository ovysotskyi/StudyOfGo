package main

import "fmt"

func main() {
	var num1, num2, num3 int
	fmt.Scan(&num1, &num2)

	for i := 0; i < 4; i++ {
		if num1 >= 0 && num1 < 10000 {
			num1_temp := num1 % 10
			num1 = num1 / 10

			num2_new := num2
			for j := 0; j < 4; j++ {
				num2_temp := num2_new % 10
				if num1_temp == num2_temp {
					num3 = (num3 * 10) + num1_temp
				}

				num2_new = num2_new / 10
				if num2_new < 1 {
					break
				}
			}

			if num1 < 1 {
				break
			}
		}
	}

	for i := 0; i < 4; i++ {
		fmt.Print(num3%10, " ")
		num3 = num3 / 10
		if num3 < 1 {
			break
		}
	}
}
