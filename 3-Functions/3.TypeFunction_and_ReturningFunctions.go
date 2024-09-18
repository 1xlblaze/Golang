package main

var x = 0

func main() {
	//local x
	x := 0
	//y := 0
	func0 := func() int { x++; return x } //func can be assigned to the variable
	func1 := incrementGlobalX             //without ()
	func2 := wrapper()
	func3 := wrapper()
	func4 := incrementy

	println(func0(), " : func0 (local x)")
	println(func1(), " : func1 (global x)")
	println(func2(), " : func2 (per func scope x1)") //
	//Functions that are returned from another functions has its own scope per returned function
	println(func3(), " : func3 (per func scope x2)")
	println(func4(), " : func4 (per func scope )")
	println("Second Increment")
	println(func0(), " : func0 (local x)")
	println(func1(), " : func1 (global x)")
	println(func2(), " : func2 (per func scope x1)")
	println(func3(), " : func3 (per func scope x2)")
	println(func4(), " : func3 (per func scope )")
}

func incrementGlobalX() int {
	x++
	return x
}

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func incrementy() int {
	y := 0
	y++
	return y
}
