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
