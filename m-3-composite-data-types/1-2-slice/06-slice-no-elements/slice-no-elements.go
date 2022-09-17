package main

import "fmt"

func main() {
	var sla []string // equivalent to:
	//sla := []string{}
	fmt.Println(sla)
	fmt.Printf("len = %d, cap = %d", len(sla), cap(sla))
}
