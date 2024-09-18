// take input like cin >> in c++
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var x int = 1337
	var y string = "string value"

	_, err := fmt.Scan(&x, &y)
	fmt.Println("You Entered x:", x, " and y: ", y, " Error: ", err)

	// take input like scanf in C
	_, err = fmt.Scanf("%d %s", &x, &y)
	fmt.Println("You Entered x:", x, " and y: ", y, " Error: ", err)

	// take input with white spaces
	var z string = "string"

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	z = scanner.Text()

	fmt.Println("You Entered z:", z)
}
