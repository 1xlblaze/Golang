// Returns x,y at the end.
package main

import "fmt"

func main() {
	x, y := split(9)
	avg := average(10, 20, 30, 40, 50)
	println("Average:", avg)
	fmt.Println(x, y)
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
	//return a,b  <- u can override default return of x,y.
} //named returns

// x value here has no use for average. just illustrates the idea of having arguments
// then a variable number of arguments
func average(x int, values ...int) float64 { //variadic function
	//print values
	fmt.Println("Single argument value: ", x)
	fmt.Println("Variable argument values: ", values)

	//calculate average
	total := 0
	for _, value := range values {
		total += value
	}

	return float64(total) / float64(len(values))
}
