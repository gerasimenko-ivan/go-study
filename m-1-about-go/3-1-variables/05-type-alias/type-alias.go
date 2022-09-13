package main

import "fmt"

func main() {
	type temperature int8
	var temperatureToday temperature = +26
	fmt.Printf("temperatureToday = %d\n", temperatureToday)

	type ID uint32
	var myId ID
	myId = 123456789
	fmt.Printf("myId = %d\n", myId)
}
