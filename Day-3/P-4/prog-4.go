/*Exercise 4: Simple Semaphore
Objective: Implement a simple semaphore to limit the number of concurrent goroutines.
Instructions:
1.	Create a semaphore with a limit of 3.
2.	Start 10 goroutines that each acquire the semaphore, perform some and then release the semaphore.
3.	Use channels to implement the semaphore.
4.	Ensure that no more than 3 goroutines are executing the work section at the same time.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	semaphore := make(chan struct{}, 3)

	var wg sync.WaitGroup

	numGoroutines := 10

	work := func(id int) {
		defer wg.Done()

		semaphore <- struct{}{}

		fmt.Printf("Goroutine %d is working\n", id)
		time.Sleep(1 * time.Second) 

		fmt.Printf("Goroutine %d is done\n", id)
		<-semaphore
	}

	for i := 1; i <= numGoroutines; i++ {
		wg.Add(1)
		go work(i)
	}

	wg.Wait()

	fmt.Println("All goroutines have finished.")
}
