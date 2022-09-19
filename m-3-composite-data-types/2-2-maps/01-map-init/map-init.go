package main

import "fmt"

func main() {
	// init using make
	idMap := make(map[string]int64)
	idMap["Tim"] = 123
	fmt.Println(idMap)

	// init using map literal
	idMap2 := map[string]int64{}
	idMap2["Tom"] = 234
	fmt.Println(idMap2)
	// or with values
	idMap3 := map[string]int64{
		"Tor": 345,
		"Tuk": 456,
		"Yuk": 567,
	}
	fmt.Println(idMap3)
}
