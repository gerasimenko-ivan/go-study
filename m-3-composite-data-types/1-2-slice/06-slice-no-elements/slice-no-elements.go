package main

import "fmt"

func main() {
	// WARNING! Bad style!
	// We should expect average number of element that will be added later
	var sla []string // equivalent to:
	//sla := []string{}
	fmt.Println(sla)
	fmt.Printf("slice %v, len = %d, cap = %d\n", sla, len(sla), cap(sla))

	// GOOD Style
	// We expect to add around 10 elements later (so there will be NO memory reallocation)
	good_sla := make([]string, 0, 10)
	fmt.Printf("slice %v, len = %d, cap = %d", good_sla, len(good_sla), cap(good_sla))
}
