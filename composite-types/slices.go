package main

import (
	"fmt"
	"reflect"
)

/*
 * Slice are sequence of objects
 * Unlike arrays, length is not a part of the type for a slice
 * Functions can accept slices of variable size for same variable
 * Comparisons cannot be perfomed on slices
 * Length of a slice can be less than capacity
 */
func main() {
	fmt.Println("\nDeclaring and intializing slices:")
	// Declaring slice; Initialized with zero-value as *nil*
	var x []int

	// Initializing slice using slice literal
	var y = []int{10, 20, 30}
	fmt.Println(y, "of size", len(y), "of type", reflect.TypeOf(y))

	var z = []int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(z, "of size", len(z))

	// Multi-dimensional slices
	var a [][]int
	fmt.Println(a, "of size", len(a))

	// Comparisons cannot be perfomed on slices
	// Only comparison can be done with *nil*
	fmt.Println("\nComparing slices:")
	fmt.Println("Check x for nil:", x, x == nil)
	fmt.Println("Check y for nil:", y, y == nil)
	// fmt.Println("Check x == y:", x, y, y == x) // Compilation error

	// Modifying slices
	fmt.Println("\nAppending to slices:")
	y = append(y, -46)
	fmt.Println("Updated y:", y)

	// Append multiple values
	x = append(x, 1, 2, 42, 100)
	fmt.Println("Updated x:", x)

	// Append using another slice
	x = append(x, z...)
	fmt.Println("Appended z to x:", x)

	// Capacity of a slice
	// Length of a slice can be less than capacity; For arrays its always same
	// Capacity is doubled till 1024, and incremented by 25% thereafter
	fmt.Println("\nCapacity & length of a slice:")
	x = []int{}
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))

	// Using built-in *make* function to delare and initialize slices
	// Allows us to specify the type, length, and, optionally, the capacity.
	// Length <= capacity.
	fmt.Println("\nDeclaring slices using make:")

	x1 := make([]int, 5)
	fmt.Println("Init:", x1, len(x1), cap(x1))
	x1 = append(x1, 10)
	fmt.Println("Appended:", x1, len(x1), cap(x1)) // Cap increased from 5 to 10

	xs := make([]int, 0, 10)
	fmt.Println("Init:", xs, len(xs), cap(xs), xs == nil)
	xs = append(xs, 5, 6, 7, 8)
	fmt.Println("Modified:", xs)

	// Using variables to create slice
	xlen := 5
	xcap := 10
	x2 := make([]string, xlen, xcap)
	fmt.Println(x2, len(x2), cap(x2))

	// A good convention for declaring slices is to use make
	// to create a slice of 0 length and specified capacity
	// and then append to it as required.
	var ys = make([]float64, 0, 10)
	ys = append(ys, 1.09)

	// Slicing slices
	// The sub-slices are variables pointing to the same memory; Just with different capacity
	// Capacity of sub-slice = capacity of parent - starting offset of sub-slice
	fmt.Println("\nSlicing slices:")
	fmt.Println(xs[:2])  // len = 2, capacity = 10
	fmt.Println(xs[4:])  // len = 0, capacity = 6
	fmt.Println(xs[1:3]) // len = 2, capacity = 9
	fmt.Println(xs[:])   // len = 4, capacity = 10

	// Arrays to slices
	fmt.Println("\nArrays to slices:")
	xa := [4]int{5, 6, 7, 8}
	ya := xa[:2]
	za := xa[2:]
	xa[0] = 10
	fmt.Println("xa:", xa)
	fmt.Println("y:", ya, reflect.TypeOf(ya), len(ya), cap(ya))
	fmt.Println("z:", za, reflect.TypeOf(za), len(za), cap(za))

	// In-built *copy* function
	// copy(destination_slice, source_slice)
	fmt.Println("\nUsing the copy function:")
	xc := []int{1, 2, 3, 4}
	yc := make([]int, 4)
	num := copy(yc, xc)
	yc[0] = -1
	fmt.Println(xc, yc, num)

	// Copying overlapping sub-slices of same slice
	num = copy(xc[:3], xc[1:])
	fmt.Println(xc, num)

	// Copying arrays to/from slices
	d := [4]int{5, 6, 7, 8}
	ycc := make([]int, 2)
	copy(ycc, d[:])
	fmt.Println(ycc, d)
	copy(d[:], xc)
	fmt.Println(d, reflect.TypeOf(d))

}
