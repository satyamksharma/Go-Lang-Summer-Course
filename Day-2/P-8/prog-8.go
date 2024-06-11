// 8. Write a program to count the number of characters in a given string

package main

import (
	"fmt"
)

func countCharacters(s string) int {
	return len(s)
}

func main() {
	// Define a string
	str := "Go Summer Course"

	// Count the number of characters in the string
	charCount := countCharacters(str)

	// Print the result
	fmt.Printf("The no. of chars in \"%s\" is: %d\n", str, charCount)
}
