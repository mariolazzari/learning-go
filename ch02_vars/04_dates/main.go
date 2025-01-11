package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Now is", time.Now())

	t := time.Date(1975, time.March, 28, 0, 0, 0, 0, time.UTC)
	fmt.Println("My time:", t)
	fmt.Println("My formatted date:", t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Fri Mar 28 00:00:00 1975")
	fmt.Printf("Parsed time type %T\n", parsedTime)

}
