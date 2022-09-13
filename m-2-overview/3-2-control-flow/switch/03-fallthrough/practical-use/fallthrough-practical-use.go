package main

import (
	"fmt"
	"go-study/helpers/inp"
	"strconv"
)

func main() {
	// TODO: still need more real-life example
	readLine := inp.ReadLine("Enter a positive number (1-10):")
	inputNumber, _ := strconv.Atoi(readLine)
	switch {
	case inputNumber < 3:
		fmt.Println("number < 3")
		fallthrough
	case inputNumber < 5:
		fmt.Println("number < 5")
		fallthrough
	case inputNumber < 7:
		fmt.Println("number < 7")
		fallthrough
	default:
		fmt.Println("default")
	}
}
