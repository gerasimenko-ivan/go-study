package main

import "fmt"

func main() {
	// var - keyword
	// x - variable name
	// int - variable type
	// = 10 - variable value
	var x int = 10

	// NO TYPE
	var y = 11 // type is based on the type of the right hand side value "infer the type"

	fmt.Printf("x = %d, y = %d", x, y)
}
