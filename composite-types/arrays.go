package main

import (
	"fmt"
	"reflect"
)

/*
 * During declaration/initialization, the unspecified values assume zero-value of the data-type.
 * Type of array also contains size of array.
 * Cannot use variable to specify array size. Array size must be known at compile time
 * Type conversion cannot be applied on different size arrays
 * Functions cannot accept arrays of *any* size
 * Indices cannot be negative or greater than array size
 */
func main() {
	// Declaration
	// In this case, all values will be 0
	var x [3]int
	x[1] = -193
	fmt.Println(x, "of size", len(x), "of type", reflect.TypeOf(x))
	// Use of in-built len function to get size of array
	// Type of array also contains size of array

	// Declaration and initialization
	// Initialization with size and array literal
	var y = [3]int{1, 2, 3}
	fmt.Println(y, "of size", len(y))

	// Initialization without size and array literal
	var a = [...]string{"a", "b", "c"}
	fmt.Println(a, "of size", len(a))

	// Initialize specific index values. Unspecified values will assume a null string
	var b = [...]string{2: "a", "b", 10: "c", "d"}
	fmt.Println(b, "of size", len(b))

	// multi-dimensional array
	fmt.Println("\nMulti-dimensional arrays")
	var c [2][3]int
	fmt.Println(c, "of size", len(c), "x", len(c[0]))

	// multi-dimensional array. Unspecified values initialized to zero-value
	var d = [3][4]int{{1, 2}, {3, 4, 5}}
	fmt.Println(d)

	// Comparisons check for respective values of 2 arrays
	fmt.Println(x, y, "equals:", x == y)
	x[0] = 1
	x[1] = 2
	x[2] = 3
	fmt.Println(x, y, "equals:", x == y)

	// Comparisons can only be done on arrays of equal size and type.
	// Following lines will give compilation error
	// fmt.Println(x, a, "equals:", x == a)
	// fmt.Println(b, a, "equals:", b == a)

}
