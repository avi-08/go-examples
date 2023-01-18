package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	// This line will not give error.
	// Program will run and print: Hello, %!s(MISSING)!
	// To check for such errors, run `go vet hello.go`
	fmt.Printf("Hello, %s!\n")

	// Correct usage
	fmt.Printf("Hello, %s!\n", "world")
}
