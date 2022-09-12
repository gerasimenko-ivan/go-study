package main

import (
	"fmt"
	"go-study/helpers/inp"
	"strings"
)

func main() {
	input := inp.ReadLine("Enter 'yes' or 'no' (or 'y'/'n'): ")

	switch {
	case input == "yes" || input == "y":
		fmt.Println("YES entered")
	case input == "no" || input == "n":
		fmt.Println("NO entered")
	case strings.Contains(input, "yes") || strings.Contains(input, "no"):
		fmt.Println("May be YES, may be NO entered")
		fallthrough // let condition below to be executed
	default:
		fmt.Println("Not 'yes' or 'no' entered")
	}
}
