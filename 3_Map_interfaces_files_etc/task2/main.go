package main

import (
	"fmt"
)

func main() {
	groupCity := map[int][]string{
		10:   {"город1", "город2", "город3"}, // города с населением 10-99 тыс. человек
		100:  {"город4", "город5", "город6"}, // города с населением 100-999 тыс. человек
		1000: {"город7", "город8", "город9"}, // города с населением 1000 тыс. человек и более
	}

	cityPopulation := map[string]int{
		"город4": 101,
		"город5": 102,
		"город1": 1,
		"город":  100,
		"город9": 150,
	}

	for key, _ := range cityPopulation {
		isArr := false
		for _, item := range groupCity[100] {
			if key == item {
				isArr = true
				break
			}
		}

		if !isArr {
			delete(cityPopulation, key)
		}
	}

	for _, item := range groupCity[100] {
		if _, isMap := cityPopulation[item]; !isMap {
			cityPopulation[item] = 100
		}
	}

	fmt.Print(cityPopulation)
}
