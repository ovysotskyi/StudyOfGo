package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	root := `3_Map_interfaces_files_etc\3.5_File_Bufio_OS\task3\task.data`
	file, err := os.Open(root)
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}
	defer file.Close()

	index := 0
	reader := bufio.NewReader(file)
	for err != io.EOF {
		str, err := reader.ReadString(';')
		if err != nil {
			panic(err)
		}
		index++
		if str == "0;" {
			fmt.Println(index)
			return
		}
	}

	//or
	//reader := bufio.NewReader(file)
	// var str []string
	// for {
	// 	line, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			break
	// 		} else {
	// 			fmt.Println(err)
	// 			return
	// 		}
	// 	}
	// 	str = strings.Split(line, ";")
	// }

	// for index, item := range str {
	// 	if item == "0" {
	// 		fmt.Print(index + 1)
	// 	}
	// }
}
