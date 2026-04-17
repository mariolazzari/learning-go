package main

import (
	"fmt"
)

func main() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Mario")) // should print Hello Mario
	fmt.Println(helloPrefix("Maria")) // should print Hello Maria
}

func prefixer(prefix string) func(string) string {
	return func(name string) string {
		return prefix + " " + name
	}
}
