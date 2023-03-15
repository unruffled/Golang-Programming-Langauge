package main

import (
	"fmt"
)

// var is used to declare a variable with package level scope

var x int
var y string
var z bool

func main() {
	fmt.Println(x, y, z)
	// The compiler assigned values to assigned variables are zeo values
	fmt.Println("Hello world!")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
