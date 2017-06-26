// Go supports anonymous functions, which can form closures. Anonymous functions
// are useful when you want to defina a function inline without having to name
// it.

package main

import "fmt"

// This function intSeq return another function, wich we define anonymously in
// the body of intSeq. The returned function closes over the variables i to form
// a closure
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	// We call intSeq, assigning the result (a function) to nextInt. This
	// function value captures its own i value, which will be updated each time
	// we call nextInt.
	nextInt := intSeq()

	// See the effect of the closure by calling nextInt a few times.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// To confirm that the state is unique to that particular function, create
	// and test a new one.
	newInts := intSeq()
	fmt.Println(newInts())
}
