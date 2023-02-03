package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// Declaration
// Accessible at package level
// Note: Only declare variables in the package block that are effectively immutable
var (
	text    string
	address string
	age     int
)

// Declaration & Initialization
// Accessible at package level
// Note: Only declare variables in the package block that are effectively immutable
var (
	// Type inference at work
	name   = "Will Wheaton"      // string
	course = "Star Wars Academy" // string
	module = "4"                 // string
	clip   = 10                  // int
)

// Package level variables can remain unused.
var unused = "This is unused variable."

func main() {
	fmt.Println("Initial values of text: [", text, "], address: [", address, "], age; [", age, "].")
	fmt.Println("Initial values of name: [", name, "], course: [", course, "].")
	fmt.Println("Initial values of module: [", module, "], clip: [", clip, "].")

	fmt.Println("---")
	fmt.Println("Type of 'text':", reflect.TypeOf(text))
	fmt.Println("Type of 'age':", reflect.TypeOf(age))

	// Yet another way of declaring using :=
	// := can create new vars as well as update existing vars
	// For updating, atleast one new var on LHS must be defined
	// Note: Good practice to use this way within function blocks only.
	iModule, err := strconv.Atoi(module)
	if err == nil {
		fmt.Println("Module added to Clips:", iModule+clip)
	} else {
		fmt.Println("Could not convert module to int:", err)
	}

	// every declared local variable must be read
	// unused_local declared but not used
	// var unused_local = "This will raise compilation error."

	// Pointers
	fmt.Println("---")
	var ptr *string = &course
	fmt.Println("Type of 'ptr':", reflect.TypeOf(ptr))
	fmt.Println("Memory location of 'course':", ptr)
	fmt.Println("Value at the address 'ptr':", *ptr)

	fmt.Println("---")
	text = "Live long and prosper"
	fmt.Println("Initialized text to:", text)
	updateTextByVal(text) // Pass by value
	fmt.Println("After applying updates on text:", text)
	updateTextByRef(&text) // Pass by reference
	fmt.Println("After applying updates on text:", text)
}

// Pass by value
func updateTextByVal(text string) string {
	text = "Now is the time"
	fmt.Println("Updated text to:", text)
	return text
}

// Pass by reference
func updateTextByRef(text *string) string {
	*text = "Godspeed ahead!"
	fmt.Println("Updated text to:", *text)
	return *text
}
