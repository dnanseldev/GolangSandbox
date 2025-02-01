package concurrency_examples

import (
	"fmt"
	"time"
)

func R1(ch chan int) {
	time.Sleep(time.Second)
	ch <- 100 //blocking operation
	fmt.Println("R1 done only after ch is read")
}
