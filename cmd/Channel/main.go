package main

import (
	"fmt"
	"laboratory/internal/concurrency_examples"
)

func main() {
	//concurrency_examples.ExChannel()

	ch := make(chan int)
	go concurrency_examples.TwoTreeFourTimes(2, ch)
	fmt.Println("It passed here")

	a, b := <-ch, <-ch
	fmt.Println(a, b)
	fmt.Println("It stoped here")

	fmt.Println(<-ch)
	fmt.Println("It should lock")
	fmt.Println(<-ch)
}
