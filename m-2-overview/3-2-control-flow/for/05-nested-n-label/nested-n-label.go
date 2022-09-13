package main

import "fmt"

func main() {
	// 'C'-style 'for'

outerLoopLabel:
	for i := 0; i < 4; i++ {
		fmt.Printf("i = %d\n", i)
		fmt.Println("  |")
		for j := 0; j < 4; j++ {
			if i+j > 3 {
				//break // only breaks internal loop 'for', and then continues outer 'for'
				break outerLoopLabel // breaks outer 'for' from internal 'for'. Last line: |=> j = 2, i*j = 4
			}
			fmt.Printf("  |=> j = %d, i+j = %d\n", j, i+j)
		}
	}
}
