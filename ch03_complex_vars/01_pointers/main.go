package main

import "fmt"

func main() {
	var p *int
	fmt.Println("Value of  p:", p)
	// fmt.Println("Value of *p:", *p) -> error

	anInt := 42
	p = &anInt
	fmt.Println("Value of  p:", p)
	fmt.Println("Value of *p:", *p)

	value1 := 42.23
	pointer1 := &value1
	fmt.Println("\nValue 1:", *pointer1)

	// update var through its pointer
	*pointer1 /= 2
	fmt.Println("*pointer1 = ", *pointer1)
	fmt.Println("*value1   = ", *&value1)

}
