package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// this example depicts invalid use of waitgroup
// waitgroup if used like this in the example (parallelism would be compromised and would be used like sequential executions)
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
		wg.Wait()

	}
	//wg.Wait()
	fmt.Println((int(time.Since(start) / time.Nanosecond)))
	//: Calls wg.Wait() inside the loop, which blocks the execution after every goroutine and leads to serialized execution,
	//negating the benefit of using goroutines in the first place.
	//wg.Wait()                              //wait for all the go routines to finish
	fmt.Println("Final Counter:", counter) // Output: Final Counter: 1000
}
