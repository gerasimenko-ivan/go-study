package main

import "fmt"

func main() {
	// init 1 (full declaration with values)
	arr1 := [...]int64{1, 2, 3, 4}
	fmt.Println(arr1)

	// init 2 (full declaration with zero values)
	arr2 := [4]int64{}
	arr2[1] = 2
	arr2[2] = 4
	arr2[3] = 8
	fmt.Println(arr2)

	// init 3 (split declaration & initialization)
	var arr3 [4]int64
	arr3[1] = 3
	arr3[2] = 9
	arr3[3] = 27
	fmt.Println(arr3)
	fmt.Println()

	// iterate 1
	for i, v := range arr1 {
		fmt.Printf("i=%d, v=%v\n", i, v)
	}
	fmt.Println()
	// iterate 2
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("i=%d, v=%v\n", i, arr2[i])
	}
	// iterate 3 - more iteration versions in m-2-overview/3-2-control-flow/for
}
