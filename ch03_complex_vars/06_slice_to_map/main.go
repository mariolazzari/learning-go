package main

import "fmt"

func main() {
	slice := []string{"apple", "banana", "orange", "date"}
	result := convertToMap(slice)

	fmt.Println(result)
}

func convertToMap(slice []string) map[string]string {
	sliceMap := make(map[string]string)

	for k, v := range slice {
		sliceMap[slice[k]] = v
	}

	return sliceMap
}
