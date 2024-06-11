// 3. Write code to demonstrate go routines

package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
		time.Sleep(5 * time.Second ) 
	}
}

func main() {
	go printMessage("Hello from goroutine 1")
	go printMessage("Hello from goroutine 2")
	go printMessage("Hello from goroutine 3")

	time.Sleep(time.Second * 3)
	
	fmt.Println("All goroutines finished")
}
