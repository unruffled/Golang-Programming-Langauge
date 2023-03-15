// Go is all about type
// Go is a static programming language that allows one to create their own type

package main

import (
	"fmt"
)

var a int

type hotdog int

// Variable b is type hotdog
var b hotdog

func main() {
	a = 42
	b = 43

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// Becasue a is of type int, b can be convered from type hotdog to int with the following conversion
	// Conversion = Type "(Expression that can be converted, ...,)"
	a = int(b)
}
