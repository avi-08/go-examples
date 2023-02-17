package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

/*
 * func <functionName> (arg1 type1, arg2, arg3 type2,...) <returnType> {}
 * *main* function doesn't accept arguments and returns nothing
 * Go doesn’t have: named and optional input parameters
 * Go supports variadic parameters
 * By convention, the error is always the last (or only) value returned from a function
 * If the function completes successfully, we return nil for the error’s value
 * Return values can be *named* in Go
 */
func main() {
	fmt.Println("\nFunctions sample:")
	MyFunc(MyFuncOpts{
		LastName: "Patel",
		Age:      50})
	MyFunc(MyFuncOpts{
		FirstName: "Joe",
		LastName:  "Smith",
	})

	fmt.Println("\nVariardic input functions sample:")
	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 2, 4, 6, 8))
	a := []int{4, 3}
	fmt.Println(addTo(3, a...))
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))

	fmt.Println("\nMultiple return values sample:")
	// must assign each value returned from a function
	// v = divAndRemainder(5,2) // !!!Compilation error
	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)

	// Ignoring returned values
	res, _, err := divAndRemainder(5, 2) // Ignore remainder using *underscore*
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(res)

	findDivAndRemainder(5, 2) // Result ignored implicitly

	// Blank returns
	fmt.Println(calculateDivAndRemainder(10, 3))

	// The type of a function is built out of the keyword func and the types of the parameters and return values
	// This combination is called the signature of the function
	// Any function that has the exact same number and types of parameters and return values meets the type signature
	fmt.Println("\nFunction Types")
	fmt.Println(reflect.TypeOf(findDivAndRemainder), reflect.TypeOf(calculateDivAndRemainder))

	// Calculator using function as values
	fmt.Println("\nSimple Calculator")
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}
	CalculatorDemo(expressions)

	// Functions using *type* keyword
	fmt.Println("\nFunctions using *type* keyword")
	fmt.Println(opMap)

}

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

// Functions with optional arguments using structs
func MyFunc(opts MyFuncOpts) error {
	if opts.Age == 0 {
		fmt.Println(fmt.Errorf("invalid age: %+v", opts))
		return errors.New("invalid age")
	}
	_, err := fmt.Printf("My name is %s %s. I'm %d years old.\n", opts.FirstName, opts.LastName, opts.Age)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to process opts: %+v", opts))
	}
	return err
}

// Variardic parameters
// The variadic parameter must be the last (or only) parameter in the input parameter list
// Indicated using 3-dots ... before the type
// The variable that’s created within the function is a slice of the specified type
func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

// Go allows multiple *return* values
// types of the return values are listed in parentheses, separated by commas
func divAndRemainder(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}

// Named return values
// pre-declaring variables that you use within the function to hold the return values
// must surround named return val‐ ues with parentheses, even if there is only a single return value
// Named return values are initialized to their zero values when created
// we can return them before any explicit use or assignment
// a named returned value is local to the function block
// For naming only few of the return values, use _ as the name for any return values you want to remain nameless
func findDivAndRemainder(numerator int, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		err = errors.New("cannot divide by zero") // Notice "err =" and not "err :="
		return result, remainder, err
	}
	result, remainder = numerator/denominator, numerator%denominator
	return result, remainder, err
}

// Blank or naked returns
// Just write return in named-return functions.
// This returns the last values assigned to the named return values.
// If not assigned anything, zero-values are returned
// blank returns make it harder to understand data flow. AVOID IT
func calculateDivAndRemainder(numerator, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return
	}
	result, remainder = numerator/denominator, numerator%denominator
	return
}

// Functions are values
// Can be assigned to variables and passed along
func add(i int, j int) int { return i + j }
func sub(i int, j int) int { return i - j }
func mul(i int, j int) int { return i * j }
func div(i int, j int) int { return i / j }

func CalculatorDemo(expressions [][]string) {
	var opMap = map[string]func(int, int) int{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := opFunc(p1, p2)
		fmt.Println(result)
	}
}

// Functions can be declared using *type* keyword
type opFuncType func(int, int) int

var opMap = map[string]opFuncType{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}
