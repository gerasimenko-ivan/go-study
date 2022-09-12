package main

import (
	"fmt"
	"go-study/helpers/inp"
)

func main() {
	input := inp.ReadLine("Enter 'yes' or 'no' (or 'y'/'n'): ")

	switch {
	case input == "yes" || input == "y":
		fmt.Println("YES entered")
	case input == "no" || input == "n":
		fmt.Println("NO entered")
	default:
		fmt.Println("Not 'yes' or 'no' entered")
	}
}
