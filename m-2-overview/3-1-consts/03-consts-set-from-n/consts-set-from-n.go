package main

import "fmt"

const (
	one = 1 << iota
	two
	four
	eight
)
const (
	hundred = 100 << iota
	twoHundred
	fourHundred
)

func main() {
	fmt.Printf("one = %d, two = %d, four = %d, eight = %d\n", one, two, four, eight)
	fmt.Printf("hundred = %d, twoHundred = %d, fourHundred = %d", hundred, twoHundred, fourHundred)
}
