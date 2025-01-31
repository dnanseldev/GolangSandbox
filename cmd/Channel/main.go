package main

import "fmt"

func main() {
	// Create a new channel
	ch := make(chan int, 1)

	// Send a value to the channel
	ch <- 42

	// Receive the value from the channel
	fmt.Println(<-ch)
}
