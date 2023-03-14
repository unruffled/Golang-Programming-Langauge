package main

import (
	"fmt"
)

var y = 42
var z = "Shaken, not stirred!"

func main() {
	fmt.Println("Hello world!", y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)

	// z = 43
	// Because z is string type, it cannot be assigned an int
	// A variable is declared to hold a value of a certain type
	// Go is a not a dynamic programming language
	// A variable of a certain type can only hold values of that type

	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
