package main

import "fmt"

func main() {
	i := 0
	for { // DO
		fmt.Printf("i = %d\n", i)
		i++
		if i > 4 { // WHILE
			break
		}
	}
	fmt.Printf("Print i = %d one more time (outside the loop)\n", i)
}
