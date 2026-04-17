package main

import (
	"errors"
	"log"
)

func main() {
	res, err := div(4, 2)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)

	_, err = div(4, 0)
	if err != nil {
		log.Fatal(err)
	}
}

func div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
