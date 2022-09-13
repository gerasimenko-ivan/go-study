package main

import (
	"fmt"
	"go-study/helpers/inp"
)

func main() {
	input := inp.ReadLine("Enter 'yes' or 'no': ")

	switch input {
	case "yes", "YES", "y", "Y": // value (OR values!) to compare with 'tag'
		fmt.Println("YES entered")
		// NO NEED to use 'BREAK' in GO
	case "no", "NO", "n", "N":
		fmt.Println("NO entered")
	default:
		fmt.Println("Not 'yes' or 'no' entered")
	}
}
