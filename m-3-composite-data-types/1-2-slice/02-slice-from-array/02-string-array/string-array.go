package main

import "fmt"

func main() {
	arr := [...]string{"a", "b", "c", "d"}
	sla := arr[1:3]
	fmt.Println(sla) // [b c]

	sla[0] = "BB"
	fmt.Println(arr) // [a BB c d] - AFFECTED!
	fmt.Println(sla) // [BB c]     - sure, we are changing it :)

	sla = append(sla, "DD")
	fmt.Println(arr) // [a BB c DD] - AFFECTED!
	fmt.Println(sla) // [BB c DD]   - sure, we are changing it :)

	sla = append(sla[:0], sla[1:]...) // remove 0 element
	fmt.Println(arr)                  // [a c DD DD] - AFFECTED! BUT: changed only elements that under the slice
	fmt.Println(sla)                  // [c DD]      - sure, we are changing it :)
}
