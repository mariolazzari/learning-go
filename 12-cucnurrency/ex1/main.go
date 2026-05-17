package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	numsCh := make(chan int)

	// Producer function
	producer := func(start int) {
		defer wg.Done()
		for i := range 10 {
			numsCh <- start + i
		}
	}

	wg.Add(2)
	go producer(1)
	go producer(11)

	// closer
	go func() {
		wg.Wait()
		close(numsCh)
	}()

	// CONSUMER
	var consumerWg sync.WaitGroup
	consumerWg.Go(func() {
		for v := range numsCh {
			fmt.Println("received:", v)
		}
	})
	consumerWg.Wait()
}
