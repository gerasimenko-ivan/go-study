package main

import "fmt"

func main() {
	sla := []int{1, 2}
	fmt.Println(sla)

	sla = append(sla, 3, 4, 5)
	fmt.Println(sla)
}
