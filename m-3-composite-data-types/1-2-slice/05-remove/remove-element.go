package main

import "fmt"

func main() {
	sla := []int{1, 2, 3, 4, 5}

	sla = append(sla[:3]) // remove 2 last elements
	fmt.Println(sla)

	sla = []int{1, 2, 3, 4, 5}
	sla = append(sla[:0], sla[2:]...) // remove 2 first elements
	fmt.Println(sla)

	sla = []int{1, 2, 3, 4, 5, 6}
	sla = append(sla[:2], sla[4:]...) // remove 2 middle elements
	fmt.Println(sla)
}
