// Key Points

// fallthrough ignores the next case condition and always executes the next case block.
// Changing i inside a case does not re-evaluate the switch conditions.
// fallthrough can only move execution downward (it cannot jump to a previous case or the default case unless it happens to be next).
package main

import "fmt"

func main() {
	i := 2
	fmt.Println("Switch for i =", i, "goes to:")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
		i = 4 // This does not affect the switch flow.
		fallthrough
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5, 6:
		fmt.Println("five or six")
	default:
		fmt.Println("default")
	}
}
