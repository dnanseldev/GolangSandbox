package main

import (
	"fmt"
	"laboratory/internal/concurrency_examples"
)

func main() {
	c := make(chan int) //channel with no buffer
	go concurrency_examples.R1(c)

	fmt.Println(<-c) //blocking operation
	fmt.Println("it is read")
	fmt.Println(<-c)   //deadlock
	fmt.Println("End") //never reached
}
