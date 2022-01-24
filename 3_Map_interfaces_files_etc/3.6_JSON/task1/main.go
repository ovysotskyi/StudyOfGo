package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Rating struct {
	Average float32
}

type Student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int
}

type Group struct {
	ID       int
	Number   string
	Year     int
	Students []Student
}

func main() {
	var group Group
	file, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(file, &group); err != nil {
		fmt.Println(err)
		return
	}

	var summ float32
	for _, item := range group.Students {
		summ += float32(len(item.Rating))
	}

	var rating Rating
	rating.Average = summ / float32(len(group.Students))

	data, err := json.MarshalIndent(&rating, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", data)
}
