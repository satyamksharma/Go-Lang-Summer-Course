// 4. Multiplexing - Write code to send data from one channel to another

package main

import (
	"fmt"
	"time"
)

func sendData(src chan int) {
	for i := 1; i <= 5; i++ {
		src <- i
		time.Sleep(time.Second) 
	}
	close(src)
}

func forwardData(src, dst chan int) {
	for data := range src {
		dst <- data
	}
	close(dst)
}

func main() {
	src := make(chan int)
	dst := make(chan int)

	go sendData(src)
	go forwardData(src, dst)

	for data := range dst {
		fmt.Println(data)
	}

	fmt.Println("All data forwarded and printed")
}
