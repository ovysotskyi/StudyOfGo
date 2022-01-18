package main

import (
	"fmt"
)

func main() {
	var test uint
	test = 727078

	fn := func(num uint) uint {
		var mod, newMod uint
		for num != 0 {
			temp := num % 10
			if temp%2 == 0 && temp != 0 {
				mod = (mod * 10) + temp
			}

			num = num / 10
		}

		for mod != 0 {
			temp := mod % 10
			newMod = (newMod * 10) + temp

			mod = mod / 10
		}

		if newMod == 0 {
			return 100
		}

		return newMod
	}

	fmt.Print(fn(test))
}
