package main

import "fmt"

func UpdateSlice(strings []string, str string) {
	size := len(strings)
	if size == 0 {
		strings = append(strings, str)
	} else {
		strings[size-1] = str
	}
	fmt.Println("inside UpdateSlice:", strings)
}

func GrowSlice(strings []string, str string) {
	strings = append(strings, str)
	fmt.Println("inside GrowSlice:", strings)
}

func main() {
	// --- UpdateSlice (change visible)
	names := []string{"Mario", "Maria"}
	fmt.Println("before UpdateSlice:", names)
	UpdateSlice(names, "Mariarosa")
	fmt.Println("after  UpdateSlice:", names)

	// --- GrowSlice (change not visible)
	names = []string{"Mario", "Maria"}
	fmt.Println("\nbefore GrowSlice:", names)
	GrowSlice(names, "Mariarosa")
	fmt.Println("after  GrowSlice:", names)

	// empty slice
	names = []string{}
	fmt.Println("\nbefore UpdateSlice (empty):", names)
	UpdateSlice(names, "Mariarosa")
	fmt.Println("after  UpdateSlice (empty):", names)
}
