package main

import "fmt"

func main() {
	type Email string
	type Id uint64
	idMap := make(map[Email]Id)

	// add elements
	idMap["vasia@gmail.com"] = 12345678
	idMap["yoy@ya.ru"] = 55555555
	idMap["go@home.ru"] = 13131313

	fmt.Println(idMap)

	// remove elements
	delete(idMap, "yoy@ya.ru")

	fmt.Println(idMap)
}
