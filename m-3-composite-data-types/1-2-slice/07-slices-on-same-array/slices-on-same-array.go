package main

import "fmt"

func main() {
	arr := [...]string{"a", "b", "c", "d", "e", "f"}
	sla1 := arr[1:4]
	sla2 := arr[2:5]
	fmt.Println(sla1)
	fmt.Println(sla2)

	fmt.Println("Changing sla2 - 2nd element")
	sla2[1] = "DD"
	fmt.Println(arr)  // AFFECTED!
	fmt.Println(sla1) // AFFECTED!
	fmt.Println(sla2)

	fmt.Println("Appending sla1")
	sla1 = append(sla1, "EE")
	fmt.Println(arr) // AFFECTED!
	fmt.Println(sla1)
	fmt.Println(sla2) // AFFECTED!
}
