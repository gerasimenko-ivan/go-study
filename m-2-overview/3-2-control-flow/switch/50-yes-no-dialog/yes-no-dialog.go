package main

import (
	"fmt"
	"go-study/helpers/inp"
)

func main() {
	for {
		input := inp.ReadLine("Enter anything ('exit' to leave): ")
		fmt.Println(input)
		switch input {
		case "yes", "YES", "Yes", "Y", "y":
			fmt.Println("Daaaaa!")
		case "no", "NO", "No", "N", "n":
			fmt.Println("Neeeet!")
		}
		if input == "exit" {
			break
		}
	}
}
