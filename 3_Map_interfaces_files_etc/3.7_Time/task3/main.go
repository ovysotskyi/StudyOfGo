package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func timeParse(str string) time.Time {
	str = strings.TrimSpace(str)
	time, _ := time.Parse("02.01.2006 15:04:05", str)
	return time
}

func main() {
	buf, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	times := strings.Split(buf, ",")

	time1 := timeParse(times[0])
	time2 := timeParse(times[1])

	if time1.After(time2) {
		fmt.Println(time1.Sub(time2))
		return
	} else {
		fmt.Println(time2.Sub(time1))
		return
	}

}
