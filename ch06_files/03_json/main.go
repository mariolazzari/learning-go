package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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

	posts := postFromJson(content)
	for _, post := range posts {
		fmt.Println(post.Title)
	}

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type Post struct {
	UserID    int
	ID        int
	Title     string
	Completed bool
}

func postFromJson(content string) []Post {
	posts := make([]Post, 0, 10)
	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	checkError(err)

	var post Post
	for decoder.More() {
		err := decoder.Decode(&post)
		checkError(err)
		posts = append(posts, post)

	}

	return posts
}
