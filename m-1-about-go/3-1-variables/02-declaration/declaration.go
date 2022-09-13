package main

import "fmt"

func main() {
	// full declaration
	var x int = 10
	// NO TYPE
	var y = 11 // type is based on the type of the right hand side value "infer the type"
	// split declaration & initialization
	var z int
	z = 12
	// short declaration
	j := 13

	fmt.Printf("x = %d, y = %d, z = %d, j = %d", x, y, z, j)
}
