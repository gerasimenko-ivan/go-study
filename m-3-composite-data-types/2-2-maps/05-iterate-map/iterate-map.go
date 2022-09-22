package main

import "fmt"

func main() {
	fmt.Println()

	idMap := map[int64]string{3: "three", 5: "five", 9: "nine", 15: "fifteen"}
	for key, val := range idMap {
		fmt.Printf("key = %d, val = %s\n", key, val)
	}

	fmt.Println()
}
