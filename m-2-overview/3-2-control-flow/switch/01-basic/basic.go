package main

import (
	"fmt"
	"go-study/helpers/inp"
)

func main() {
	input := inp.ReadLine("Enter 'yes' or 'no': ")

	switch input {
	case "yes": // value to compare with 'tag'
		fmt.Println("YES entered")
	case "no":
		fmt.Println("NO entered")
	default:
		fmt.Println("Not 'yes' or 'no' entered")
	}
}
