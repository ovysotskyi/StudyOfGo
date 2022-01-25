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

	date, err := time.Parse(time.RFC3339, str)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(date)

	fmt.Println(date.Format(time.UnixDate))
}
