// 7. Using go routine send name, srn and cgpa

package main

import (
	"fmt"
	"time"
)

func sendName(ch chan string, name string) {
	time.Sleep(5 * time.Millisecond)
	ch <- name
	close(ch)
}

func sendSRN(ch chan string, srn string) {
	time.Sleep(5 * time.Millisecond)
	ch <- srn
	close(ch)
}

func sendCGPA(ch chan float64, cgpa float64) {
	time.Sleep(5 * time.Millisecond)
	ch <- cgpa
	close(ch)
}

func main() {
	// Create channels for name, SRN, and CGPA
	nameChan := make(chan string)
	srnChan := make(chan string)
	cgpaChan := make(chan float64)

	// Start goroutines to send name, SRN, and CGPA
	go sendName(nameChan, "Satyam Kumar")
	go sendSRN(srnChan, "PES2UG21CS486")
	go sendCGPA(cgpaChan, 8.35)

	// Receive and print the values from the channels
	name := <-nameChan
	srn := <-srnChan
	cgpa := <-cgpaChan

	fmt.Println("Name:", name)
	fmt.Println("SRN:", srn)
	fmt.Println("CGPA:", cgpa)
}
