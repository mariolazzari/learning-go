package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

const (
	MAX   = 100_000
	CHUNK = 1000
)

func buildSqrtMap() map[int]float64 {
	result := make(map[int]float64, MAX)

	for i := range MAX {
		result[i] = math.Sqrt(float64(i))
	}

	return result
}

func main() {
	start := time.Now()
	// The map is generated only once and then cached.
	sqrtMap := sync.OnceValue(buildSqrtMap)

	// First call initializes the cache.
	cache := sqrtMap()

	fmt.Println(time.Since(start))

	// Read every chunk value.
	for i := 0; i < MAX; i += CHUNK {
		fmt.Printf("sqrt(%d) = %.4f\n", i, cache[i])
	}
}
