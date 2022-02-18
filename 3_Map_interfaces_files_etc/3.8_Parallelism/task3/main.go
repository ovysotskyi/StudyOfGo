package main

import (
	"fmt"
	"time"
)

func removeDuplicates(inputStream chan string, outputStream chan string) {
	var tmp string
	for item := range inputStream {
		if tmp != item {
			tmp = item
			outputStream <- item
		}
	}

	close(outputStream)
}
func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)

	go removeDuplicates(inputStream, outputStream)
	go func() {
		for item := range outputStream {
			fmt.Println("chan", item)
		}
	}()

	for _, item := range "111221133331444444" {
		time.Sleep(500 * time.Microsecond)
		inputStream <- string(item)
	}

	close(inputStream)

}
