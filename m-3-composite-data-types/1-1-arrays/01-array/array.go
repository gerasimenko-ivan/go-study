package main

import "fmt"

func main() {
	// fixed-length / same data-type / indexes from 0

	// full declaration (redundant)
	var point [3]float64 = [3]float64{1, 2, 3}
	// could be simplified to:
	var point_ = [3]float64{1, 3, 5}
	fmt.Println(point)
	fmt.Println(point_)

	// split declaration & initialization
	var x [5]int
	x[0] = 9
	fmt.Println(x)

	// no explicit length
	y := [...]int{3, 33, 333}
	fmt.Println(y)
}
