package concurrency_examples

import (
	"fmt"
)

func Routine(ch chan int) {

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	fmt.Println("It executed!")
	ch <- 5

}

func isPrimeNumber(number int) bool {
	if number <= 1 {
		return false
	}

	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func PrimeNumbers(nbrs int, ch chan int) {

	init := 2
	for i := 0; i < nbrs; i++ {

		for prime := init; ; prime++ {
			if isPrimeNumber(prime) {
				ch <- prime
				init = prime + 1
				//time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}
	close(ch)

}
