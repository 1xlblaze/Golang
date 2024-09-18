// Returns x,y at the end.
package main

import "fmt"

func main() {
	x, y := split(9)
	fmt.Println(x, y)
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
	//return a,b  <- u can override default return of x,y.
}
