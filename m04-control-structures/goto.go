package main

import (
	"fmt"
	"math/rand"
)

/*
 * In Go, a goto statement specifies a labeled line of code and execution jumps to it
 * Go forbids jumps that skip over variable declarations
 * Go forbids jumps that go into an inner or parallel block
 */
func main() {
	a := rand.Intn(10)
	for a < 100 {
		if a%5 == 0 {
			goto done
		}
		a = a*2 + 1
	}
	fmt.Println("do something when the loop completes normally")
done:
	fmt.Println("do complicated stuff no matter why we left the loop")
	fmt.Println(a)
}
