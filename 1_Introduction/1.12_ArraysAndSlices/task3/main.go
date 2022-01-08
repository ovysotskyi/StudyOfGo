package main
import "fmt"

func main()  {
	array := [5]int{}
	var a int
	for i:=0; i < 5; i++{
		fmt.Scan(&a)
		array[i] = a
	}
    // здесь ваш код
	var aMax int
	for _, value := range array {
		if value < 0 {
			if -value > aMax {
				aMax = value
			}
		} else {
			if value > aMax {
				aMax = value
			}
		}

		
	}

	fmt.Print(aMax)
}