package main

import (
	"fmt"
	concurrency_examples "laboratory/internal/concurrency_examples"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.NumCPU())

	/*Sequential execution*/
	concurrency_examples.Speak("Daniel", "Why you don't speak to me", 3)
	concurrency_examples.Speak("Hosana", "I'm busy", 1)

	/*Concurrent execution*/
	go concurrency_examples.Speak("Daniel", "Why you don't speak to me", 30)
	go concurrency_examples.Speak("Hosana", "I'm busy", 5)

	time.Sleep(10 * time.Second)
}
