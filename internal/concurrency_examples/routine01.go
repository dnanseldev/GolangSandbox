package concurrency_examples

import (
	"fmt"
	"time"
)

func Speak(name, text string, amount int) {
	for i := 0; i < amount; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteration: %d)\n", name, text, i+1)
	}
}
