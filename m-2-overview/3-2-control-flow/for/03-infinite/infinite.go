package main

import "fmt"

func main() {
	i := 0
	for {
		if i > 4 {
			break
		}
		fmt.Printf("i = %d\n", i)
		i++
	}
	fmt.Printf("Print i = %d one more time (outside the loop)\n", i)
}
