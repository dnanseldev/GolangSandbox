package main

import (
	"fmt"
	"laboratory/internal/concurrency_examples"
)

func main() {
	t1 := concurrency_examples.Title("https://www.google.com", "https://www.cod3r.com.br")
	t2 := concurrency_examples.Title("https://www.youtube.com", "https://www.amazon.com")

	fmt.Println("First: ", <-t1)
	fmt.Println("Second: ", <-t2)

}
