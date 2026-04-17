package main

import (
	"io"
	"log"
	"os"
)

func main() {
	size, err := fileLen("./main.go")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("File size:", size)

	size, err = fileLen("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(size)
}

func fileLen(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	n, err := io.Copy(io.Discard, file)
	if err != nil {
		return 0, err
	}

	return int(n), nil
}
