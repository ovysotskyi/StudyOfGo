package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewScanner(os.Stdin)

	var sum int

	for {
		in.Scan()
		num, err := strconv.Atoi(in.Text())
		if err != nil {
			break
		}

		sum += num
	}

	io.WriteString(os.Stdout, strconv.Itoa(sum))
}
