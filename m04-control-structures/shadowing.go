package main

import "fmt"

/*
 * Each place where a declaration occurs is called a block
 * Package block: Variables, constants, types, and functions declared outside of any functions
 * File block: names for other packages that are valid for the file that contains the import statement
 * Universe block: the block that contains all other blocks;
 * Universe block: Also contains predeclared identifiers like int, string, make, nil, true, false,...
 * All of the variables defined at the top level of a function (including the parameters to a function) are in a block
 * Any identifier defined in any outer block can be accessed from within any inner block
 * Shadowing: A declaration in the inner block with the same name as an identifier in the outer block
 */
func main() {
	// Shadowing
	// As long as the shadowing variable exists, you cannot access a shadowed variable
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5 // Shadows x from outer block
		fmt.Println(x)
	}
	fmt.Println(x) // Shadowing variable doesn't exist here

	if x > 5 {
		fmt.Println(x)
		x = 5
		fmt.Println(x)
	}
	fmt.Println(x)

	// Shadowing during multiple assignment
	y := 10
	if y > 5 {
		y, z := 5, 20 // y is shadowed here
		fmt.Println(y, z)
	}
	fmt.Println(y)

	// Shadowing package import
	// Results in compilation error; fmt.Println undefined (type string has no field or method Println)
	// fmt := "oops"
	// fmt.Println(fmt)

	// Never redefine any of the identifiers in the *universe block*
	fmt.Println(true)
	true := 10
	fmt.Println(true)
}
