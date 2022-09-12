package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		fmt.Printf("i = %d\n", i)
	}

	fmt.Println(" ======= ")

	// if 'i' is required outside 'for'
	i := 0
	for ; i < 4; i++ {
		fmt.Printf("i = %d\n", i)
	}
	fmt.Printf("Print i = %d one more time\n", i)
}
