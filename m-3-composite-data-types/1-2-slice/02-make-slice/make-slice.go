package main

import "fmt"

func main() {
	sla1 := make([]int, 10)                                  // len = cap = 10
	fmt.Println(sla1)                                        // [0 0 0 0 0 0 0 0 0 0]
	fmt.Printf("len = %d, cap = %d\n", len(sla1), cap(sla1)) // len = 10, cap = 10

	sla2 := make([]int, 5, 10)                             // len = 5, cap = 10
	fmt.Println(sla2)                                      // [0 0 0 0 0]
	fmt.Printf("len = %d, cap = %d", len(sla2), cap(sla2)) // len = 5, cap = 10
}
