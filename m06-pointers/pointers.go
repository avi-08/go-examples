package main

import (
	"fmt"
	"reflect"
)

/*
 * pointers store memory address of other variables
 * pointer itself is a variable
 * every pointer, no matter what type it is pointing to, is always the same size:
 * zero value for a pointer is nil
 * nil is an untyped identifier that represents the lack of a value for certain types
 * Before dereferencing a pointer, always make sure that it is not *nil*, program panics otherwise
 * slices, maps, and functions are implemented using pointers
 */
func main() {
	fmt.Println("\npointer example1:")
	// The & is the address operator.
	// It precedes a value type and returns the address of the memory location where the value is stored
	greet := "Hello"
	greetPtr := &greet
	fmt.Println("initial greet:", greet)
	fmt.Println("stored at:", greetPtr, reflect.TypeOf(greetPtr))

	// The * is the indirection operator.
	// It precedes a variable of pointer type and returns the pointed-to value.
	// This is called dereferencing
	fmt.Println("value at the address *greetPtr: ", *greetPtr)
	*greetPtr = "Hello World!"
	fmt.Println("updated greet using pointer:", greet)

	// Another example
	fmt.Println("\npointer example2:")
	x := 10
	pointerToX := &x
	fmt.Println(pointerToX)  // prints a memory address
	fmt.Println(*pointerToX) // prints 10
	z := 5 + *pointerToX
	fmt.Println(z) // prints 15

	// program will panic if you attempt to dereference a nil pointer
	fmt.Println("\nnil pointer example:")
	var xptr *int
	fmt.Println(xptr, xptr == nil) // prints true
	// Uncomment below line to see panic
	// fmt.Println(*xptr)       // panics

	// built-in function *new* creates a pointer variable.
	// It returns a pointer to a zero value instance of the provided type
	fmt.Println("\npointer using *new* function:")
	var xp = new(int)
	fmt.Println(xp == nil) // prints false
	fmt.Println(*xp)       // prints 0

	// You can’t use an & before a primitive literal (numbers, booleans, and strings) or a constant
	// This is because they don’t have memory addresses; they exist only at compile time
	// Use a helper function to turn a constant value into a pointer
	fmt.Println("\ngetting pointers of constants/literals:")
	p := person{
		FirstName: "Pat",
		// MiddleName: "Perry", // This will give compilation error
		MiddleName: stringp("Perry"), // This works because of call-by-value
		LastName:   "Peterson",
	}
	fmt.Println(p)

	// Go is a call by value language, the values passed to functions are copies
	// However, if a pointer is passed to a function, the function gets a copy of the pointer.
	// This still points to the original data, which means that the original data can be modified by the called function.

	// A good convention
	// Rather than populating a struct by passing a pointer to it into a function
	// have the function instantiate and return the struct
	fmt.Println(BadMakeFoo(&Foo{}))
	fmt.Println(GoodMakeFoo())
}

// Helper function to get a pointer from string contant
func stringp(s string) *string {
	return &s
}

type person struct {
	FirstName  string
	MiddleName *string // pointer
	LastName   string
}

type Foo struct {
	Field1 string
	Field2 int
}

// Don't use like this
func BadMakeFoo(f *Foo) error {
	f.Field1 = "val"
	f.Field2 = 20
	return nil
}

// This is better
func GoodMakeFoo() (Foo, error) {
	f := Foo{
		Field1: "val",
		Field2: 20,
	}
	return f, nil
}
