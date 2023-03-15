package main

import (
	"fmt"
)

var x = 42
var y = "James Bond"
var z = true

func main() {
	// Sprintf formats according to a format specifier and returns the resulting string
	// %v is the value in a default format
	// The double quotes format the values as a string
	// The \t inserts a tab between the string values
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}
