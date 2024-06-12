/*Exercise 3: Reader-Writer Problem
Objective: Implement a reader-writer lock using a mutex and condition variables.
Instructions:
1.	Create a shared resource
2.	Implement two types of goroutines: readers and writers.
3.	Readers should read the resource without locking out other readers but should not read when a writer is writing.
4.	Writers should have exclusive access to the resource.
5.	Use a mutex and condition variables to handle synchronization.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

type SharedResource struct {
	data int
	mu   sync.Mutex
	cond *sync.Cond
	readers int
	writing bool
}

func NewSharedResource() *SharedResource {
	sr := &SharedResource{}
	sr.cond = sync.NewCond(&sr.mu)
	return sr
}

func (sr *SharedResource) Reader(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	sr.mu.Lock()
	for sr.writing {
		sr.cond.Wait()
	}
	sr.readers++
	sr.mu.Unlock()

	fmt.Printf("Reader %d is reading the data: %d\n", id, sr.data)
	time.Sleep(100 * time.Millisecond) 

	sr.mu.Lock()
	sr.readers--
	if sr.readers == 0 {
		sr.cond.Broadcast()
	}
	sr.mu.Unlock()
}

func (sr *SharedResource) Writer(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	sr.mu.Lock()
	for sr.readers > 0 || sr.writing {
		sr.cond.Wait()
	}
	sr.writing = true
	sr.mu.Unlock()

	sr.data++
	fmt.Printf("Writer %d is writing the data: %d\n", id, sr.data)
	time.Sleep(100 * time.Millisecond) 

	sr.mu.Lock()
	sr.writing = false
	sr.cond.Broadcast()
	sr.mu.Unlock()
}

func main() {
	sr := NewSharedResource()
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go sr.Reader(i, &wg)
	}

	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go sr.Writer(i, &wg)
	}

	wg.Wait()
}
