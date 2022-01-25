package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const now = 1589570165

func main() {
	buf, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	buf = strings.TrimSpace(buf)
	buf = strings.ReplaceAll(buf, " ", "")
	buf = strings.ReplaceAll(buf, "мин.", "m")
	buf = strings.ReplaceAll(buf, "сек.", "s")

	d, _ := time.ParseDuration(buf)

	newDate := time.Unix(int64(now+d.Seconds()), 0).UTC()

	fmt.Println(newDate.Format("Mon Jan 02 15:04:05 UTC 2006"))

}
