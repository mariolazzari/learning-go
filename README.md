# Learning Go (2nd edition)

## Setup

```sh
go mod init github.com/mariolazzari/learning-go/01-setup/hello-world
go run .
go build
```

```go
package main

import "fmt"

func main() {
	fmt.Println("Ciao Mario!")
}
```

```makefile
.DEFAULT_GOAL := build

BINARY_NAME := app

.PHONY: fmt vet build clean run

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o $(BINARY_NAME) .

run: build
	./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)
```

## Types

### Types ex1

- Write a program that declares an integer variable called i with the value 20. 
- Assign i to a floating-point variable named f. 
- Print out i and f.

```go
package main

import "fmt"

func main() {
	i := 20
	var f float64 = float64(i)

	fmt.Printf("i = %d, f = %.3f\n", i, f)
}
```

### Types ex2

- Write a program that declares a constant called value that can be assigned to both an integer and a floating-point variable. 
- Assign it to an integer called i and a floating-point variable called f. 
- Print out i and f.

```go
package main

import "fmt"

func main() {
	const value = 1
	i := value
	var f float64 = value // inferred

	fmt.Printf("i = %d, f = %.3f\n", i, f)
}
```

### Types ex 3

- Write a program with three variables, one named b of type byte, one named smallI of type int32, and one named bigI of type uint64. 
- Assign each variable the maximum legal value for its type; 
- then add 1 to each variable. 
- Print out their values.

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var b byte = math.MaxUint8
	b++

	var smallI int32 = math.MaxInt32
	smallI++

	var bigI uint64 = math.MaxUint64
	bigI++

	fmt.Printf("b = %d\n", b)
	fmt.Printf("smallI = %d\n", smallI)
	fmt.Printf("bigI = %d\n", bigI)
}
```

## Composite types

### Composite types ex1

- Write a program that defines a variable named greetings of type slice of strings with the following values: "Hello", "Hola", "नमस्कार", "こんにちは", and "Привіт". 
- Create a subslice containing the first two values; 
- a second subslice with the second, third, and fourth values; 
- and a third subslice with the fourth and fifth values. 
- Print out all four slices.

```go
package main

import "fmt"

func main() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	first2 := greetings[0:2]
	firstMid := greetings[1:4]
	last2 := greetings[3:]

	fmt.Println("First 2 slice:", first2)
	fmt.Println("Middle  slice:", firstMid)
	fmt.Println("Last 2 slice:", last2)
}
```

### Composite types ex2

- Write a program that defines a string variable called message with the value "Hi 😂 and 😭" 
- and prints the fourth rune in it as a character, not a number.

```go
package main

import "fmt"

func main() {
	msg := "Hi 😂 and 😭"
	// convert to rune slice
	runes := []rune(msg)
	// print the 4th rune
	fmt.Printf("%c\n", runes[3])
	fmt.Printf("%c\n", runes[9])
}
```

### Composite types ex3

- Write a program that defines a struct called Employee with three fields: firstName, lastName, and id. 
- The first two fields are of type string, and the last field (id) is of type int. 
- Create three instances of this struct using whatever values you’d like. 
- Initialize the first one using the struct literal style without names, the second using the struct literal style with names, and the third with a var declaration. 
- Use dot notation to populate the fields in the third struct. 
- Print out all three structs.
