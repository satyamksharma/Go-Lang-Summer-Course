// 6. Create three channels of different types and each channel will send message after 5 milliseconds

package main

import (
	"fmt"
	"time"
)

func sendInt(ch chan int, value int) {
	time.Sleep(5 * time.Millisecond)
	ch <- value
	close(ch)
}

func sendString(ch chan string, value string) {
	time.Sleep(5 * time.Millisecond)
	ch <- value
	close(ch)
}

func sendFloat(ch chan float64, value float64) {
	time.Sleep(5 * time.Millisecond)
	ch <- value
	close(ch)
}

func main() {
	// Create channels of different types
	intChan := make(chan int)
	stringChan := make(chan string)
	floatChan := make(chan float64)

	// Start goroutines to send messages
	go sendInt(intChan, 42)
	go sendString(stringChan, "Hello, Go!")
	go sendFloat(floatChan, 3.14)

	// Receive and print messages from the channels
	fmt.Println("Received from intChannel:", <-intChan)
	fmt.Println("Received from stringChannel:", <-stringChan)
	fmt.Println("Received from floatChannel:", <-floatChan)
}
