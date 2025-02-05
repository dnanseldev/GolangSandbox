package concurrency_examples

import (
	"fmt"
	"time"
)

func DelayToExecute(t time.Duration) {
	fmt.Println("Starting method execution")

	time.Sleep(t)

	fmt.Println("Finishing method execution")
}
