package main

import (
	"fmt"
)

// The var keyword can be used outside the scope of a function body
var z = 99

// Declares a variable with identfier t of type int and assigns the value of 0 value of type into to t
var t int

func main() {
	// Short declaration operator :=
	// := x and 42 are operands
	// Declare a varaible and assign a value of a certain type
	// A short declaration operation needs to be within a function body
	x := 42
	fmt.Println(x)

	// Keywords
	// var is a keyword
	// var is used to declare a variable outside the function body
	// It is best practice to limit the scope of a variable and use a short declaration operator
	var y = 43
	fmt.Println(y)
	fmt.Println(z)
	printz()
}

func printz() {
	fmt.Println(z)
	fmt.Println("Print the value of t", t)
}
