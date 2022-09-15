package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5}
	sla1 := arr[1:3]
	sla2 := arr[2:4]
	fmt.Println(sla1) // [2 3]
	fmt.Println(sla2) // [3 4]

	sla1[1] = 333
	fmt.Println(arr)  // [1 2 333 4 5]
	fmt.Println(sla1) // [2 333]
	fmt.Println(sla2) // [333 4]

	sla1 = append(sla1, 444)
	fmt.Println(arr)  // [1 2 333 444 5] - AFFECTED!
	fmt.Println(sla1) // [2 333 444]
	fmt.Println(sla2) // [333 444]       - AFFECTED!
}
