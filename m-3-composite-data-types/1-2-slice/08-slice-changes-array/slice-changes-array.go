package main

import "fmt"

func main() {
	arr := []int64{1, 2, 3}
	sla := arr[:]
	fmt.Println(sla)

	fmt.Println("Changing slice, changes array")
	sla[0] = 111
	fmt.Println(arr)
	fmt.Println(sla)

	fmt.Println("After increasing of slice capacity, changing slice does NOT change  array")
	sla = append(sla, 4)
	sla[0] = 11111
	fmt.Println(arr)
	fmt.Println(sla)
}
