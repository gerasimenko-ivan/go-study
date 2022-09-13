package main

import "fmt"

func main() {
	// zero-value is set to uninitialized value
	var i int
	fmt.Printf("i = '%d'\n", i) // i = '0'

	var f float64
	fmt.Printf("f = '%f'\n", f) // f = '0.000000'

	var s string
	fmt.Printf("s = '%s'\n", s) // s = ''

	var b bool
	fmt.Printf("b = '%t'\n", b) // b = 'false'
}
