package main

import (
	"fmt"
)

func task2(myChan chan string, str string) {
	for i := 0; i < 5; i++ {
		myChan <- str + " "
	}

}

func main() {
	myChan := make(chan string)
	var str string
	fmt.Scan(&str)

	go task2(myChan, str)

	fmt.Print(<-myChan)
}
