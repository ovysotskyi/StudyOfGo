package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	buf, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str := strings.TrimSpace(buf)
	date, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		fmt.Println(err)
	}

	if date.Hour() == 13 || date.Hour() > 13 {
		date = date.AddDate(0, 0, 1)
		fmt.Println(date.Format("2006-01-02 15:04:05"))
	} else {
		fmt.Println(date.Format("2006-01-02 15:04:05"))

	}
}
