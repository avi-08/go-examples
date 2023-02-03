package main

import (
	"fmt"
	"reflect"
)

/*
 * There are no implicit type conversions in go.
 * Mismatched types must be explicitly converted before performing operations.
 */
func main() {
	var x int = 10
	var y float64 = 100.91
	// invalid operation: x + y (mismatched types int and float64)
	// var zerr float64 = x + y
	var z float64 = float64(x) + y // explicit conversion needed
	var a int = x - int(y)         // explicit conversion needed
	var b = x * x
	var c bool = true
	var d = 1
	fmt.Println(reflect.TypeOf(x), x)
	fmt.Println(reflect.TypeOf(y), y)
	fmt.Println(reflect.TypeOf(z), z)
	fmt.Println(reflect.TypeOf(a), a)
	fmt.Println(reflect.TypeOf(b), b)
	fmt.Println(reflect.TypeOf(c), c)
	fmt.Println(reflect.TypeOf(d), d)
	// d is not a boolean
	// No implicit conversion from a non-zero integer to boolean true
	// if d {
	// 	fmt.Println("This will never ever print")
	// }
	if c {
		fmt.Println("I might be printed, if c allows it.")
	}

}
