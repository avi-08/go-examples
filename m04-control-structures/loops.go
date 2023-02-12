package main

import (
	"fmt"
	"reflect"
)

/*
 * *for* is the only looping keyword in Go
 * There are 4 different formats to use the *for* keyword
 * While iterating, the *for-range* loop copies the value from the compound type to the value variable
 * Loops can be labelled; Labels are always indented to the same level as the braces for the block
 */
func main() {
	// Use1: Complete for statement
	// must use := to initialize the variables; var is not legal here
	fmt.Println("\nA loop using complete *for* statement")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Use2: Condition-only for statement
	fmt.Println("\nA loop using condition-only *for* statement")
	i := 1
	for i < 10 {
		fmt.Println(i)
		i = i * 2
	}

	// Use3: Infinite for statement
	// fmt.Println("\nA loop using infinite *for* statement")
	// for {
	// 	fmt.Println("Hello") // Runs forever
	// }

	// break and continue keywords
	fmt.Println("\nUsing break and continue keywords")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 && i%3 == 0 {
			fmt.Println("FizzBuzz")
			break
		}
		if i%2 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}

	// Use4: for-range statement
	fmt.Println("\nA loop using for-range statement")
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v) // enumeration in-order
	}

	fmt.Println("\nLooping over maps")
	nameMap := map[string]string{
		"prefix":    "Mr",
		"firstName": "Michael",
		"lastName":  "Scott",
		"suffix":    "",
	}
	// The formatting functions (like fmt.Println) always out‐ put maps with their keys in ascending sorted order.
	fmt.Println(nameMap)
	for k, v := range nameMap {
		fmt.Println(k, v) // enumeration not necessarily in-order
	}

	// Skipping value from maps
	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
	for k := range uniqueNames {
		fmt.Println(k)
	}

	// for-range loop iterates over the runes, not the bytes
	fmt.Println("\nLooping over strings")
	samples := []string{"hello", "apple_π!"}
	for _, sample := range samples {
		for i, r := range sample { // returns index, byte
			fmt.Println(i, reflect.TypeOf(r), string(r))
		}
		fmt.Println()
	}

	// v contains copy of the value; Modifications will not change the values in slice
	fmt.Println("\n Example of copy values during iteration")
	for _, v := range evenVals {
		v *= 2
	}
	fmt.Println(evenVals)

	// Labelling the loops; Also notice the label indentation
	fmt.Println("\nLabelled loops:")
outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}
		}
		fmt.Println()
	}

}
