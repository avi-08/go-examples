package main

import (
	"fmt"
	"reflect"
)

// accessible at package level
const pi = 3.147

func main() {
	fmt.Println("Value of pi is:", pi)

	const g = 9.8
	fmt.Println("Acceleration due to gravity is:", g, "metres per second per second")
	fmt.Println("Type of 'g':", reflect.TypeOf(g))

	// Uncommenting the below line will result in compilation error
	// g = 10
	// fmt.Println("Acceleration due to gravity(rounded off) is:", g, "metres per second per second")
}
