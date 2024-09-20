//Using wait-group to use more than 1 function to write to the same channel. Using a waitGroup to close the channel once the two writing functions signal wg.Done()

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go func() {
		//wg.Add(1)
		defer wg.Done()
		ch <- 1
	}()
	wg.Add(1)
	go func() {
		//wg.Add(1)
		defer wg.Done()
		ch <- 2
	}()
	//close the channel after both writing functions have finished.
	go func() {
		wg.Wait() // Wait for all the goroutines to finish
		close(ch) // Close the channel once all done
	}()

	for i := range ch {
		fmt.Println(i) //prints 1 and 2 in the order they were sent to the channel.
	}
	// The for i := range ch loop reads from the channel until it is closed. Once the channel is closed, the loop terminates automatically.
}
