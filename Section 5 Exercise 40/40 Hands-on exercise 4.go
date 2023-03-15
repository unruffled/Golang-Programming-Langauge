package main

import (
	"fmt"
)

// Creating a type hotdog of underlying type int
type hotdog int

// Declaring the variable x as type hotdog
var x hotdog

func main() {
	fmt.Println(x)
	// Prints the formatted type of x using %T
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)
}
