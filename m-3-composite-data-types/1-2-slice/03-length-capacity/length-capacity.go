package main

import "fmt"

func main() {
	arr := [...]string{"a", "b", "c", "d"}
	sla := arr[1:3]
	fmt.Println(sla)
	// len = 2, cap = 3
	fmt.Printf("len = %d, cap = %d\n", len(sla), cap(sla))

	sla = append(sla, "A", "B")
	fmt.Println(sla)
	// len = 4, cap = 6 - capacity doubled!
	fmt.Printf("len = %d, cap = %d\n", len(sla), cap(sla))

	sla = append(sla, "C", "D", "E")
	fmt.Println(sla)
	// len = 7, cap = 12 - capacity doubled! (From last capacity value - 6
	fmt.Printf("len = %d, cap = %d\n", len(sla), cap(sla))

	sla = append(sla[:0], sla[6:]...) // remove 0 - 5 elements
	fmt.Println(sla)
	// len = 1, cap = 12 - capacity NOT decreased!
	fmt.Printf("len = %d, cap = %d\n", len(sla), cap(sla))
}
