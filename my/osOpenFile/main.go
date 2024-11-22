package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("sample.txt", os.O_RDWR|os.O_CREATE, 0644)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	// Получаем информацию о файле
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("File Name: %s\n", fileInfo.Name())
	fmt.Printf("Size: %d bytes\n", fileInfo.Size())
	fmt.Printf("Permissions: %s\n", fileInfo.Mode())
	fmt.Printf("Last Modified: %s\n", fileInfo.ModTime())

}
