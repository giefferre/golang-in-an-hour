package main

import (
	"fmt"
)

func main() {
	// START1 OMIT
	// Declaration AND assignment
	aString := "Ciao Mondo!"
	// STOP1 OMIT

	// START2 OMIT
	// Declaration
	var anotherString string

	// Assignment
	anotherString = "Hallo Welt!"
	// STOP2 OMIT

	// START3 OMIT
	fmt.Println(aString)
	fmt.Println(anotherString)
	// STOP3 OMIT
}
