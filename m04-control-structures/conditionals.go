package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// If-elseIf-else ladder
	n := rand.Intn(10) // For 10, n=1; default random number seed in math/rand is hard-coded.
	if n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}

	// declare variables that are scoped to the condition and to both the if and else blocks
	if n := rand.Intn(10); n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}
	fmt.Println(n) // n from previous declaration
}
