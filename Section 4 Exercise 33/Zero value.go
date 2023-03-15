package main

import (
	"fmt"
)

var y string

func main() {
	// Declare a variable to be of a certain type and assign a value of that type to the variable
	fmt.Println("Hello world!")
	fmt.Println("Printing the value of y: ", y)
	fmt.Printf("%T\n", y)
	fmt.Println("Hello universe!")

	y = "Bond, James Bond"

	fmt.Println("Printing the value of y: ", y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)

	// fmt.Print() will print without a line
	// fmt.Printf() needs a format string
	s := fmt.Sprintf("%#x\n", y)
	fmt.Println(s)
}
