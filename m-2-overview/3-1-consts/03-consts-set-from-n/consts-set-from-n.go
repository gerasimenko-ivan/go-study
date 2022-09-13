package main

import "fmt"

func main() {
	const (
		one = 1 << iota
		two
		four
		eight
	)
	fmt.Printf("one = %d, two = %d, four = %d, eight = %d\n", one, two, four, eight)

	const (
		hundred = 100 << iota
		twoHundred
		fourHundred
	)
	fmt.Printf("hundred = %d, twoHundred = %d, fourHundred = %d", hundred, twoHundred, fourHundred)
}
