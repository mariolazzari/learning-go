package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/todos"

func main() {
	res, err := http.Get(url)
	checkError(err)

	fmt.Printf("response type: %T\n", res)
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	checkError(err)

	content := string(bytes)
	fmt.Println(content)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
