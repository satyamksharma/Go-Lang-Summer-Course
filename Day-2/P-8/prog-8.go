// 8. Write a program to count the number of characters in a given string

package main

import (
	"fmt"
)

func countCharacters(s string) int {
	return len(s)
}

func main() {
	str := "Go Summer Course"

	charCount := countCharacters(str)

	fmt.Printf("The no. of chars in \"%s\" is: %d\n", str, charCount)
}
