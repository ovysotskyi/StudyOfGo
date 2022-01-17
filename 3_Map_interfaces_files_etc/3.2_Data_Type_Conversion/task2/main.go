package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.Trim(str, "\n\r")
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, ",", ".")

	arr := strings.Split(str, ";")
	num1, _ := strconv.ParseFloat(arr[0], 64)
	num2, _ := strconv.ParseFloat(arr[1], 64)

	result := num1 / num2

	fmt.Printf("%.4f", result)

}
