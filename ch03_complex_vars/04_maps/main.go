package main

import (
	"fmt"
	"sort"
)

func main() {
	// init map with make
	states := make(map[string]string)
	fmt.Println(states)

	// set map element
	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"
	fmt.Println(states)

	// get map element
	california := states["CA"]
	fmt.Println(california)
	california = states["Ca"]
	fmt.Println("Ca =", california)

	// delete map element
	delete(states, "CA")
	fmt.Println(states)

	states["NY"] = "New York"

	// loop
	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println(keys)

	// iterator loop
	for i := range keys {
		fmt.Println(states[keys[i]])
	}

}
