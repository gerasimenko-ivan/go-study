package main

import (
	"fmt"
	"go-study/helpers/inp"
)

func main() {
	input := inp.ReadLine("Enter 'yes' or 'no' (or 'y'/'n'): ")

	switch { // No tag
	case input == "yes" || input == "y": // condition, could be complex (not like 'tag == case-value')
		fmt.Println("YES entered")
	case input == "no" || input == "n":
		fmt.Println("NO entered")
	default:
		fmt.Println("Not 'yes' or 'no' entered")
	}
}
