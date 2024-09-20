package main

import (
	"fmt"
)

func main() {

	//Allocate new capacity
	var s []int
	s = make([]int, 5, 5)
	x := append(s, 1, 2, 3)
	x[0] = 1337
	s[0] = 6800
	fmt.Println(s, x) //[6800 0 0 0 0] [1337 0 0 0 0 1 2 3]

	//Doesn't allocat new capacity and return reference
	var b []int
	b = make([]int, 5, 150)
	y := append(b, 1, 2, 3)
	y[0] = 1337
	b[0] = 6800
	fmt.Println(b, y) //[6800 0 0 0 0] [6800 0 0 0 0 1 2 3]
	//notice that 1337 is overwritten

}
