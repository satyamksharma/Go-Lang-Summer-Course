/*
Exercise 1: Simple Mutex
Objective: Implement a simple counter with a mutex to prevent race conditions.
Instructions:
1.	Create a program that has a counter initialized to 0.
2.	Start 10 goroutines, each incrementing the counter 1000 times.
3.	Use a mutex to ensure that the increments do not cause race conditions.
4.	Print the final value of the counter.
*/
package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

func increment() {
	for i := 0; i < 1000; i++ {
		mutex.Lock()
		counter++
		mutex.Unlock()
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go increment()
	}

	var wg sync.WaitGroup
	wg.Add(10)
	wg.Wait()

	fmt.Println("Final counter value:", counter)
}
