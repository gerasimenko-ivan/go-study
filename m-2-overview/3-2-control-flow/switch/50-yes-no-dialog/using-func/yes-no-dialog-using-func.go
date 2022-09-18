package main

import (
	"fmt"
	"go-study/helpers/inp"
)

func main() {
	fmt.Println("Is it Ok to terminate this program?")
	isOk := inp.YesNoRequest()
	fmt.Printf("You entered '%v'", isOk)
}
