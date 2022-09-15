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
	fmt.Printf("slice len = %d, s cap = %d\n", len(slice), cap(slice)) // slice len = 2, s cap = 4

	slice[0] = 333
	fmt.Println(slice) // output: [333 3]
	fmt.Println(array) // array is still underlying slice, so output: [1 333 3 4 5]

}
