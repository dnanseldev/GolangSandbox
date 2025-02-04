package main

import (
	"fmt"
	"laboratory/internal/concurrency_examples"
	"log"
	"time"
)

func _main() {
	ch := make(chan int, 10000)

	start := time.Now()
	go concurrency_examples.PrimeNumbers(cap(ch), ch)

	for prime := range ch {
		fmt.Printf("%d ", prime)
	}
	fmt.Println("\nEnd")
	elapsed := time.Since(start)

	log.Printf("Prime numbers %s", elapsed)
}
