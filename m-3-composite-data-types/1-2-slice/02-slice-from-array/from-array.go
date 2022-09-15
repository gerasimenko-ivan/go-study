package main

import "fmt"

func main() {
	array := [...]int{1, 2, 3, 4, 5}
	slice := array[1:3] // array[from:to]
	// creates slice from array
	//  - 'from' - from element 1 (including) - so capacity - 4
	//  - 'to' - to element 3 (excluding)     - so length   - 2
	//  - length - number of elements "copied" to slice "from"-"to"
	//  - capacity - length of array from "from" to the end of the array

	// slice len = 2, s cap = 4
	fmt.Printf("slice len = %d, s cap = %d\n", len(slice), cap(slice))

	sla1 := array[:]  // make slice from FULL array
	fmt.Println(sla1) // [1 2 3 4 5]

	sla2 := array[2:] // make slice from array: from 2-index to the end of array
	fmt.Println(sla2) // [3 4 5]

	sla3 := array[:3] // make slice from array: from the beginning of array till the 3-index (NOT included)
	fmt.Println(sla3)
}
