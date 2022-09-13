package main

import "fmt"

func main() {
	type Grade int
	const (
		A Grade = iota // 0, iota - auto init from 0 & +1 each const
		B              // 1
		C              // 2
	)
	myGrade := A
	yourGrade := C
	fmt.Printf("myGrade = %d, yourGrade = %d\n", myGrade, yourGrade)

	const (
		Sun = iota
		Mon
		Tue
		Wed
		Thr
		Fri
		Sat
	)

	today := Mon
	fmt.Printf("today = %d", today)
}
