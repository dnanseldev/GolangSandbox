package concurrency_examples

import "time"

func PrintNumbers(cn chan<- int) {
	for i := 0; i < 10; i++ {
		cn <- i
		time.Sleep(time.Millisecond * 150)
	}
	close(cn)
}
