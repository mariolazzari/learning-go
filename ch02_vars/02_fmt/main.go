package main

import "fmt"

func main() {
	// print
	str1 := "Prima string"
	str2 := "Seconda string"
	str3 := "Terza string"
	fmt.Println(str1, str2, str3)

	anInt := 42
	stringLen, err := fmt.Println("anInt", anInt)
	if err == nil {
		fmt.Println("String length:", stringLen)
	}

	// formatting strings
	fmt.Printf("Value is %v\n", anInt)
	fmt.Printf("Data type is %T\n", anInt)

}
