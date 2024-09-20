package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int) //bidirection unbuffered channel

	go func() {
		for i := 0; i < 9; i++ {
			time.Sleep(time.Second)
			c <- i //sending
		}
	}()

	go func() {
		for {
			fmt.Println(<-c) //receiving
		}
	}()

	time.Sleep(time.Second * 15)
}
