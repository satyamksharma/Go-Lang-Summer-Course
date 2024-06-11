// 8. Write a program to count the number of characters in a given string

package main

import (
	"fmt"
)

type Counter interface {
	Count() int
}

type StringCounter struct {
	Text string
}

func (sc StringCounter) Count() int {
	return len(sc.Text)
}

func main() {
	input := StringCounter{Text: "Hello, Go!"}

	var counter Counter = input

	fmt.Printf("The number of characters in the string is: %d\n", counter.Count())
}
