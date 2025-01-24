package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "Hello from Go"
	file, err := os.Create("./file.txt")
	checkError(err)
	defer file.Close()

	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("File written with %v chars\n", length)
	defer readFile("./file.txt")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(fileName string) {
	data, err := os.ReadFile(fileName)
	checkError(err)
	fmt.Printf("Text read from file %v: %v", fileName, string(data))
}
