package main

import "fmt"

func main() {
	const x = 123
	const (
		y = 1.23
		z = "hi"
	)
	fmt.Printf("x = %d, y = %f, z = %s", x, y, z)
}
