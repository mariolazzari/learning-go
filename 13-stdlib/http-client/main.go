package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	// client
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	// request
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		log.Fatalf("request: %v", err)
	}
	req.Header.Add("X-My-Client", "Learing Go")

	// response
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("response: %v", err)
	}
	defer res.Body.Close()

	// response status
	if res.StatusCode != http.StatusOK {
		log.Fatalf("unexpected status: %s", res.Status)
	}

	// decode body
	var post Post
	err = json.NewDecoder(res.Body).Decode(&post)
	if err != nil {
		log.Fatalf("post decode: %v", err)
	}

	fmt.Printf("Post title: %s\n\n", post.Title)
	fmt.Printf("Post: %+v\n", post)
}
