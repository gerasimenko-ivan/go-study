package main

import (
	"fmt"
	"go-study/helpers/inp"
	"math/rand"
	"strconv"
)

func main() {
	// intro
	fmt.Println()
	fmt.Println("**************************************************")
	fmt.Println("You have 3 attempts to guess a number from 1 to 9.")
	fmt.Println("Let's go! Enter a number!")
	fmt.Println()

	// game
	rnd := rand.Intn(9) + 1 // TODO: find solution for 'real' random (not only 6 every time)
	rndSecret := strconv.Itoa(rnd)
	for i := 0; i < 3; i++ {
		num := inp.ReadLine(fmt.Sprintf("Attempt No.%d: ", i))
		if num == rndSecret {
			fmt.Println("***********************")
			fmt.Println("*      You WIN!!!     *")
			fmt.Printf("*   Correct num is %s  *\n", rndSecret)
			fmt.Println("***********************")
			return
		}
	}

	// fail
	fmt.Println()
	fmt.Println("***********************")
	fmt.Println("*     You loose!!!    *")
	fmt.Printf("*   Correct num is %s  *\n", rndSecret)
	fmt.Println("***********************")
}
