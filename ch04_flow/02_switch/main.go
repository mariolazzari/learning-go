package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(7) + 1

	var result string

	switch dow {
	case 1:
		result = "Sunday"
	case 2:
		result = "Monday"
	case 3:
		result = "Tuesday"
	case 4:
		result = "Wednesday"
	case 5:
		result = "Tuesday"
	case 6:
		result = "Friday"
	case 7:
		result = "Saturday"
		// fallthrough
	default:
		result = "Another day"
	}

	fmt.Println(result)
}
