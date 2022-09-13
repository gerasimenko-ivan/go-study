package main

import "fmt"

func main() {

outerLoopLabel:
	for i := 0; i < 4; i++ {
		fmt.Printf("i = %d\n", i)
		fmt.Println("  |")
		for j := 0; j < 4; j++ {
			if i+j > 3 {
				fmt.Println("  BREAK to LABEL")
				//break // only breaks internal loop 'for', and then continues outer 'for'
				break outerLoopLabel // breaks outer 'for' from internal 'for'. Last line: |=> j = 2, i*j = 4
			}
			fmt.Printf("  |=> j = %d, i+j = %d\n", j, i+j)
		}
	}

	fmt.Println("")
	fmt.Println("  +++++++++++ ")
	fmt.Println("  BREAK without LABEL")
	for i := 0; i < 4; i++ {
		fmt.Printf("i = %d\n", i)
		fmt.Println("  |")
		for j := 0; j < 4; j++ {
			if i+j > 3 {
				fmt.Println("  BREAK (only internal loop)")
				break // only breaks internal loop 'for', and then continues outer 'for'
			}
			fmt.Printf("  |=> j = %d, i+j = %d\n", j, i+j)
		}
	}
}
