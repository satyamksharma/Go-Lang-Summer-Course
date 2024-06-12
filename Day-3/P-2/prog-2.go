/*Exercise 2: Mutex with Multiple Resources
Objective: Simulate a scenario where multiple resources are protected by different mutexes.
Instructions:
1.	Create a program with two counters and two corresponding mutexes.
2.	Start 5 goroutines that increment the first counter and 5 goroutines that increment the second counter.
3.	Each goroutine should increment its respective counter 1000 times.
4.	Use the mutexes to protect access to the counters.
5.	Print the final values of both counters.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	counter1 := 0
	counter2 := 0

	var mutex1 sync.Mutex
	var mutex2 sync.Mutex

	var wg sync.WaitGroup

	numGoroutines := 5

	incrementsPerGoroutine := 1000

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1) 

		go func() {
			defer wg.Done() 

			for j := 0; j < incrementsPerGoroutine; j++ {
				mutex1.Lock()   
				counter1++      
				mutex1.Unlock() 
			}
		}()
	}

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1) 

		go func() {
			defer wg.Done() 

			for j := 0; j < incrementsPerGoroutine; j++ {
				mutex2.Lock()   
				counter2++      
				mutex2.Unlock() 
			}
		}()
	}

	wg.Wait()

	fmt.Println("Final counter1 value:", counter1)
	fmt.Println("Final counter2 value:", counter2)
}
