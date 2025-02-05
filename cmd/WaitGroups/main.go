package main

import (
	"laboratory/internal/concurrency_examples"
	"sync"
	"time"
)

func main() {

	//TODO: Implement a way to wait a function finishes its execution
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			concurrency_examples.DelayToExecute(time.Second * 2)
		}()
	}

	wg.Wait()
}
