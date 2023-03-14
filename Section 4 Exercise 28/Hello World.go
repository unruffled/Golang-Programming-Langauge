// Every Go program needs package main
// Every Go program needs to have func main()
// Programming control flow
// 1. Sequence
// 2. Loop
// 3. Conditional
// Control flow describes the order in which statements or function calls are evaluated
// Go programs are read from top to bottom.
// When the code exists func main, the program is over.

// Packages
// fmt is an example of a Go package
// Packages are prewritten documentation that can be used.
// Packages are a way for programmers to organize their code.
// Instead of using "folders", "packages" are used to group similar prewritten code.

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world!")
	foo()
	fmt.Println("This is the third Hello world!")

	// A for loop is an example of a loop control folow
	for i := 0; i < 100; i++ {
		// An if statement is an example of a conditional statement
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("This is the final Hello world!")
	bar()
}

func foo() {
	fmt.Println("This is the second Hello world!")
}

func bar() {
	fmt.Println("No, this is really the last Hello world!")
}
