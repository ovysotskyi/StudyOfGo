package main

import (
	"fmt"
	"time"
)

func work(n int) int {
	if n > 3 {
		time.Sleep(time.Millisecond * 500)
		return n + 1
	} else {
		time.Sleep(time.Millisecond * 500)
		return n - 1
	}
}

func main() {
	coll := make(map[int]int)
	var arr [10]int
	for index := range arr {
		fmt.Scan(&arr[index])
	}

	for _, item := range arr {
		if _, isMap := coll[item]; !isMap {
			coll[item] = work(item)
		}

		fmt.Printf("%d ", coll[item])
	}
}
