package concurrency_examples

import (
	"fmt"
	"time"
)

func ExChannel() {
	// Create a new channel
	ch := make(chan int, 1)

	// Send a value to the channel
	ch <- 42

	// Receive the value from the channel
	fmt.Println(<-ch)
}

func TwoTreeFourTimes(base int, ch chan int) {
	time.Sleep(time.Second)
	ch <- base * 2

	time.Sleep(time.Second)
	ch <- base * 3

	time.Sleep(3 * time.Second)
	ch <- base * 4

}
