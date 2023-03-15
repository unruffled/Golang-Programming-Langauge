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
// From the package fmt, we're using the function Println with the dot notation, fmt.Println()
// It's not a good programming practice to use unused characters in your code.
// Code pollution is leaving unused variables in code.

// Identifiers
// Identifiers name program entities such as variables and types.
// An identifier is a sequence of one or more letters or a sequence that must begin with a letter.
// Some identifiers are predeclared (e.g., bool, byte, int, int8, etc.)

// Keywords
// Go uses keywords that are reserached and may not be used as identifiers (e.g., break, default, func, go, goto, etc.)

// Short declaration operatior :=
// The short declaration operator allows us to write code

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

	// The short declaration operation declares and assigns a value to an identifier.
	x := 42
	// A statement is a small standalone element of a program.
	x = 92
	fmt.Println("The value of x is: ", x)
	bar()
}

func foo() {
	fmt.Println("This is the second Hello world!")
}

func bar() {
	fmt.Println("No, this is really the last Hello world!")
	fmt.Println(1, 2, 3, true, false)
	n, e := fmt.Println("Let's see if I can increase the bytes...", 42, true)
	fmt.Println(n)
	fmt.Println(e)
}
