package main

import (
	"fmt"
	"time"
)

func main() {
	go say("Hello")
	fmt.Println("Main")

	// anonymous function
	go func(msg string) {
		fmt.Println(msg)
	}("Hello from anonymous function")

	time.Sleep(2 * time.Second)
	fmt.Println("End")
}

func say(meg string) {
	time.Sleep(1 * time.Second)
	fmt.Println(meg)
}
