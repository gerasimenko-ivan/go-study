package main

import "fmt"

func main() {
	type Grade int
	const (
		A Grade = iota
		B
		C
		D
		F
	)
	myGrade := A
	yourGrade := D
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
