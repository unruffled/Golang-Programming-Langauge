package main

import (
	"fmt"
)

func main() {
	// Short declaration operator
	// := is the short declaration operator used to assign a value to an identifier
	x, y, z := 42, "James Bond", true

	fmt.Println(x, y, z)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
