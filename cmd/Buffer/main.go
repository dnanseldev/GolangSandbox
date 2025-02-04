package main

import (
	"fmt"
	"laboratory/internal/concurrency_examples"
)

func main() {

	cn := make(chan int)

	go concurrency_examples.PrintNumbers(cn)

	for v := range cn {
		fmt.Printf("read from channel: %d\n", v)
	}
}
