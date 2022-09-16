package main

import "fmt"

func main() {
	/* Three ways to create slice
	- initialize directly (slice literal - represented below)
	- create it after array (create array, init slice on that array)
	- make()
	*/

	// length - real number of elements in s slice
	// capacity - number of elements allocated in memory for slice
	//   capacity extends by x2 (present value x2) when len > cap after adding slice elements
	s := []int{1, 2, 3}
	fmt.Printf("s len = %d, s cap = %d\n", len(s), cap(s)) // s len = 3, s cap = 3
	s = append(s, 4)
	fmt.Printf("s len = %d, s cap = %d\n", len(s), cap(s)) // s len = 4, s cap = 6
	s = append(s, 5, 6, 7)
	fmt.Printf("s len = %d, s cap = %d\n", len(s), cap(s)) // s len = 7, s cap = 12
}
