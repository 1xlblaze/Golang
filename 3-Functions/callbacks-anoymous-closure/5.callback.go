package main

import "fmt"

// callback function : function passed as an argument to the other function
func visit(g []int, call func()) { //callback
	for _, v := range g {
		println("inside visit", v)
		call()
		//anonymous function calling
	}
}

func main() {
	g := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	visit(g, func() {
		for _, v := range g { //function is defined inside function therefore it can acess the value of the outer function : closure
			println(v)
		}
		fmt.Println("inside main")
	}) //anonymous function calling

}

// The variable g is defined in main.
// The anonymous function inside visit(g, func() {...}) does not receive g as a parameter, but it can still access it because g is in the scope of the outer function (main).
// Go "closes over" the variable g, meaning that even though the anonymous function doesn’t explicitly take g as a parameter, it can still reference it.

// This works because Go remembers the environment (the variables in scope) where the anonymous function is defined.
// Since g is in the scope of main, and the anonymous function is also defined within main, the function has access to g.
