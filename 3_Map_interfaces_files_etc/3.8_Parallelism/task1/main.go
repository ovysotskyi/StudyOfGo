package main

import (
	"fmt"
)

func task(myChan chan int, num int) {
	myChan <- num + 1
}

func main() {
	myChan := make(chan int)
	var num int
	fmt.Scan(&num)

	go task(myChan, num)

	fmt.Print(<-myChan)

}
