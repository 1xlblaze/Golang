package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	counter int32
	mu      sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done() // delete

	// Lock the mutex to ensure exclusive access to counter
	mu.Lock()
	//counter++
	atomic.AddInt32(&counter, 1) //atomic operation ensure thread safe execution
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup // wait for go routines to finish
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1) //add
		go increment(&wg)

	}
	wg.Wait()
	fmt.Println((int(time.Since(start) / time.Nanosecond)))

	//wg.Wait()                              //wait for all the go routines to finish
	fmt.Println("Final Counter:", counter) // Output: Final Counter: 1000
}
