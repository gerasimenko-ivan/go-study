package main

import (
	"fmt"
	"go-study/helpers/inp"
)

func main() {
	input := inp.ReadLine("Enter yes or no: ")

	switch input {
	case "yes":
		fmt.Println("YES entered")
	case "no":
		fmt.Println("NO entered")
	default:
		fmt.Println("Not 'yes' or 'no' entered")
	}
}
