//2. Implement a greeting message using a channel

package main

import (
	"fmt"
)

func main() {
	greetingChannel := make(chan string)

	go func() {
		greeting := "Hello, Go!"
		greetingChannel <- greeting 
	}()

	message := <-greetingChannel

	fmt.Println(message)
}
