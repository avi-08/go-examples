package main

import (
	"fmt"
	"reflect"
)

// accessible at package level
// untyped constant declaration
// type is inferred
const pi = 3.147

/*
 * consts can only hold values that the compiler can figure out at compile time.
 * arrays, slices, maps, or structs cannot be declared as constants as they are mutable.
 */
func main() {
	fmt.Println("Value of pi is:", pi)
	fmt.Println("Type of 'pi':", reflect.TypeOf(pi)) // inferred

	// typed constant declaration
	const g float64 = 9.8
	// Following is invalid for a typed constant.
	// cannot use 9.8 (untyped float constant) as int value in constant declaration (truncated)
	// const gerr int = 9.8
	fmt.Println("Acceleration due to gravity is:", g, "metres per second per second")

	// Uncommenting the below line will result in compilation error
	// g = 10
	// fmt.Println("Acceleration due to gravity(rounded off) is:", g, "metres per second per second")

	// unused constants are allowed
	// if a constant isn’t used, it is simply not included in the compiled binary.
	const e = 1.72
}
