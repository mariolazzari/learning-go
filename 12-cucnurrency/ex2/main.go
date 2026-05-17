package main

import (
	"fmt"
	"sync"
)

func run() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	var wg sync.WaitGroup

	// Producer function
	producer := func(ch chan int, start int) {
		defer wg.Done()

		for i := range 10 {
			ch <- start + i
		}
	}

	wg.Add(2)

	go producer(ch1, 0)
	go producer(ch2, 100)

	// Closer goroutine
	go func() {
		wg.Wait()
		close(ch1)
		close(ch2)
	}()

	// Consumer loop (fan-in via select)
	closed1 := false
	closed2 := false

	for {
		if closed1 && closed2 {
			break
		}

		select {
		case v, ok := <-ch1:
			if !ok {
				closed1 = true
				ch1 = nil
				continue
			}
			fmt.Println("received from goroutine-1:", v)

		case v, ok := <-ch2:
			if !ok {
				closed2 = true
				ch2 = nil
				continue
			}
			fmt.Println("received from goroutine-2:", v)
		}
	}
}

func main() {
	run()
}
