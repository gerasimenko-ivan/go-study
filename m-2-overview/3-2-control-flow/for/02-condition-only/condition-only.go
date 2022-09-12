package main

import "fmt"

func main() {
	/* Condition only 'for' - imitates 'while' in other languages
	Condition only 'for' looks like a BAD STYLE.
	Reasons:
	 - you can forget '++' part & get infinite loop
	 - you have to search for '++' part in code (if 'for' is more then one line, this takes time)
	*/

	i := 0
	for i < 4 {
		fmt.Printf("i = %d\n", i)
		i++
	}
	fmt.Printf("Print i = %d one more time (outside the loop)\n", i)
}
