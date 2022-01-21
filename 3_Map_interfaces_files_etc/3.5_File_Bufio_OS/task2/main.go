package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil // Проигнорируем директории
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}

	r := csv.NewReader(file)
	data, err := r.ReadAll()
	if err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}

	if len(data) > 1 {
		fmt.Println(data[4][2])
	}

	defer file.Close()

	return nil
}

func main() {
	root := `3_Map_interfaces_files_etc\3.5_File_Bufio_OS\task2\root`
	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}
}
