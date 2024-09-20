// Using only channels without waitGroup
package main

import (
	"fmt"
	"strconv"
)

func main() {
	c := make(chan string)
	done := make(chan bool)

	n := 2

	for i := 0; i < n; i++ {
		go func(i int) {
			for t := i * 10; t < i*10+10; t++ {
				c <- "From " + strconv.Itoa(i) + " : " + strconv.Itoa(t)
			}
			done <- true
		}(i)
	}

	go func() {
		//receive the n-dones from the go-routines
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	for x := range c {
		fmt.Println(x)
	}
}
