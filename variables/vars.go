package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// Declaration
var (
	text    string
	address string
	age     int
)

// Declaration & Initialization
var (
	// Type inference at work
	name   = "Will Wheaton"
	course = "Star Wars Academy"
	module = "4"
	clip   = 10
)

func main() {
	fmt.Println("Initial values of text: [", text, "], address: [", address, "], age; [", age, "].")
	fmt.Println("Initial values of name: [", name, "], course: [", course, "].")
	fmt.Println("Initial values of module: [", module, "], clip: [", clip, "].")

	fmt.Println("---")
	fmt.Println("Type of 'text':", reflect.TypeOf(text))
	fmt.Println("Type of 'age':", reflect.TypeOf(age))

	iModule, err := strconv.Atoi(module)
	if err == nil {
		fmt.Println("Module added to Clips:", iModule+clip)
	} else {
		fmt.Println("Could not convert module to int:", err)
	}

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
