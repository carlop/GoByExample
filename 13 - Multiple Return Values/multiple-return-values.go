// Go has built-in suppornt for multiple return values. This feature is used
// ofthen in idiomatic Go, for example to return both result and error values
// from a function.

package main

import (
	"fmt"
)

// The (int, int) int this function signature hows that the function return 2
// ints.
func vals() (int, int) {
	return 3, 7
}

func main() {

	// Here we use the 2 different return values from the call with multiple
	// assignment.
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// If you only want a subset of the returned values, use the black
	// identifier _.
	_, c := vals()
	fmt.Println(c)
}
