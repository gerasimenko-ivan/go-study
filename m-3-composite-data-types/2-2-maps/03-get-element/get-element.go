package main

import "fmt"

func main() {
	idMap := map[string]int64{
		"zero": 0,
		"one":  1,
		"two":  2,
	}
	fmt.Println(idMap)
	fmt.Println(idMap["zero"])
	fmt.Println(idMap["minus one"]) // still 0

	// diff '0' default from '0' from element
	val, ok := idMap["zero"]
	fmt.Println(val, ok)
	val2, ok2 := idMap["minus one"]
	fmt.Println(val2, ok2)

	// typical check
	if val, ok := idMap["zero"]; ok {
		fmt.Printf("'%v' is from map", val)
	} else {
		fmt.Printf("'%v' is default", val)
	}
}
