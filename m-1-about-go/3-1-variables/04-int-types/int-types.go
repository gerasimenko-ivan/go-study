package main

import "fmt"

func main() {
	/* usually if you declare int variable, you choose int 64 */

	var integer8 int8 = -128         // -128; 127
	var unsignedInteger8 uint8 = 255 // 0; 255
	// (u)int 16, 32, 64
	// byte = uint8
	// rune
	var b byte = 255        // 0; 255
	var r rune = 2147483647 // -2147483648; 2147483647

	fmt.Printf("integer8 = %d\n", integer8)
	fmt.Printf("unsignedInteger8 = %d\n", unsignedInteger8)
	fmt.Printf("b = %d\n", b)
	fmt.Printf("r = %d\n", r)
}
