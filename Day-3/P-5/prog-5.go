/*
Exercise 5: Bounded Buffer
Objective: Implement a producer-consumer problem using a semaphore.
Instructions:
1.	Create a bounded buffer with a fixed size as 5
2.	Implement a producer goroutine that adds items to the buffer.
3.	Implement a consumer goroutine that removes items from the buffer.
4.	Use semaphores to ensure that the buffer does not overflow or underflow.
5.	Use a mutex to protect access to the buffer.
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

const bufferSize = 5

type BoundedBuffer struct {
	buffer []int
	mutex  sync.Mutex
	empty  chan struct{}
	full   chan struct{}
}

func NewBoundedBuffer(size int) *BoundedBuffer {
	return &BoundedBuffer{
		buffer: make([]int, 0, size),
		empty:  make(chan struct{}, size),
		full:   make(chan struct{}, size),
	}
}

func (bb *BoundedBuffer) Produce(item int, wg *sync.WaitGroup) {
	defer wg.Done()

	bb.full <- struct{}{}

	bb.mutex.Lock()
	bb.buffer = append(bb.buffer, item)
	fmt.Printf("Produced: %d\n", item)
	bb.mutex.Unlock()

	<-bb.empty
}

func (bb *BoundedBuffer) Consume(wg *sync.WaitGroup) {
	defer wg.Done()

	bb.empty <- struct{}{}

	bb.mutex.Lock()
	item := bb.buffer[0]
	bb.buffer = bb.buffer[1:]
	fmt.Printf("Consumed: %d\n", item)
	bb.mutex.Unlock()

	<-bb.full
}

func main() {
	bb := NewBoundedBuffer(bufferSize)
	var wg sync.WaitGroup

	for i := 0; i < bufferSize; i++ {
		bb.empty <- struct{}{}
	}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go bb.Produce(i, &wg)
		time.Sleep(100 * time.Millisecond) 
	}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go bb.Consume(&wg)
		time.Sleep(200 * time.Millisecond) 
	}

	wg.Wait()
	fmt.Println("All operations have completed.")
}
