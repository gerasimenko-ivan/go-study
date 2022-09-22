package main

import "fmt"

func main() {
	fmt.Println()

	userId := map[int]string{13: "Vasya", 14: "Tolya"}
	fmt.Printf("len(userId) = %d\n", len(userId)) // 2

	userId[22] = "Sonya"
	userId[25] = "Kate"
	fmt.Printf("len(userId) = %d\n", len(userId)) // 4

	delete(userId, 14)
	fmt.Printf("len(userId) = %d", len(userId)) // 3

	fmt.Println()
}
