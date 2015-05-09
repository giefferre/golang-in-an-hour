package main

import (
	"fmt"
)

// START1 OMIT
func Sum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	a := 7
	b := 5

	output := fmt.Sprintf("Given a %d, b %d \nThe sum is: %d", a, b, Sum(a, b))
	fmt.Println(output)
}

//STOP1 OMIT
