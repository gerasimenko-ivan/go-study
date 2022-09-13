package main

import "fmt"

func main() {
	x := [...]string{"a", "b", "c"}
	for i, v := range x {
		fmt.Printf("i = %d, v = '%s'\n", i, v)

		if i == 2 {
			// ATTENTION: changing V does NOT change element of array!
			v = "changed V"
			fmt.Printf("i = %d, v = '%s', real x[2] = '%v'\n", i, v, x[2])
		}
	}

	fmt.Println(x)
}
