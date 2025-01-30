package main

import (
	"fmt"
	errors_examples "laboratory/internal/errors_examples"
	"regexp"
)

func main() {

	r := errors_examples.Must(regexp.Compile("123"))
	fmt.Println(r)
}
