package main

import (
	"fmt"
	"reflect"
	"strconv"
)

/*
 * Go uses a sequence of *bytes* to represent a string
 * string literals support UTF-8 encoding
 * strings are immutable in golang
 */
func main() {
	var name string
	name = "My name is Michael Scott!"
	loc := "Location -> Scranton, Pennsylvania"
	var work = []string{"Regional Manager", "Dunder Mifflin"}

	fmt.Println(name, "@", work, loc)

	// Accessing index values and slicing
	fmt.Println("\nAccessing index values and slicing")
	fmt.Println(loc[0], reflect.TypeOf(loc[0])) // Returns byte
	fmt.Println(name[11:])
	fmt.Println(work[1][:6])

	// Length of string returns number of bytes, not the code points.
	// For characters represented in more than 1 bytes, eg emojis, this difference is critical
	fmt.Println("\nUnderstanding length function:")
	status := "All systems live \u2713"
	fmt.Println(status)
	var n = len(status)
	fmt.Println(n, status[n-1], status[n-5:n-2], status[n-3:]) // [n-5:n-2] Does not print "e ✓"

	// rune and byte variables can be converted into string
	fmt.Println("\nRunes and bytes with strings:")
	var a rune = 'x' // rune is 32 bits long; equivalent to int32; Represents UTF-8
	fmt.Println(a, reflect.TypeOf(a))
	var s1 string = string(a)
	var b byte = 'y' // byte is 8 bits long; equivalent to uint8; Represents ASCII
	var s2 string = string(b)
	fmt.Println(b, reflect.TypeOf(b))
	fmt.Println(s1, s2)

	// This will print "A", not string "65"
	var x int = 65
	var y = string(x)
	fmt.Println(x, y) // int cannot be converted into string this way
	xstr := strconv.FormatInt(int64(x), 10)
	fmt.Println(xstr)

	// String to []byte and []rune
	// Observe the difference in length due to "✓"
	var bs []byte = []byte(status)
	var rs []rune = []rune(status)
	fmt.Println(bs, len(bs))
	fmt.Println(rs, len(rs))

}
