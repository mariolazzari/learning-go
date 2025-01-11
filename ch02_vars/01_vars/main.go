package main

import "fmt"

// constants
const aConst = 64

func main() {
	var aString string = "Go string"
	fmt.Println(aString)
	fmt.Printf("The var type is %T\n", aString)

	var anInt int = 42
	fmt.Println(anInt)
	fmt.Printf("The var type is %T\n", anInt)

	// default
	var defInt int
	fmt.Println(defInt)
	fmt.Printf("The var type is %T\n", defInt)

	// inferred type
	var anotherString = "Another string"
	fmt.Println(anotherString)
	fmt.Printf("The var type is %T\n", anotherString)

	var anotherInt = 75
	fmt.Println(anotherInt)
	fmt.Printf("The var type is %T\n", anotherInt)

	// shortcut (only inside functions)
	myString := "Shortcut"
	fmt.Println(myString)
	fmt.Printf("The var type is %T\n", myString)

	fmt.Println(aConst)
	fmt.Printf("The var type is %T\n", aConst)

}
