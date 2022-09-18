package inp

import (
	"fmt"
)

func YesNoRequest() bool { // TODO: find out how to overload & add default value to inputs (greeting string) {
	requestsLimit := 10
	requestNum := 0
	for {
		//fmt.Println("")
		input := ReadLine("Enter 'yes' or 'no' (y/n): ")
		switch input {
		case "yes", "YES", "Yes", "y":
			return true
		case "no", "NO", "No", "n":
			return false
		default:
			fmt.Printf("Your input '%v' was not recognized as 'yes' or 'no'.\n", input)
			fmt.Println("Please, check your input and try again.")
		}
		requestNum++
		fmt.Printf("requestNum = '%v'\n", requestNum)
		if requestNum >= requestsLimit {
			fmt.Printf("Looks like you can not handle input (for %v retries)\n", requestsLimit)
			fmt.Println("So program exits now")
			// TODO: find out how to throw errors
		}
	}
}
