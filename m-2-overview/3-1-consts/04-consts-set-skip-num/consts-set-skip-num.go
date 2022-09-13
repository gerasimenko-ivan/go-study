package main

import "fmt"

/* BTW: I DO NOT know practical usages of this skipping */

type Animals int

const (
	Rabbit Animals = iota // 0
	_                     // skip 1
	Snake                 // 2
	_                     // skip 3
	Lion                  // 4
)

func main() {
	fmt.Printf("Rabbit = %d, Snake = %d, Lion = %d", Rabbit, Snake, Lion)
}
