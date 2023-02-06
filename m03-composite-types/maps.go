package main

import (
	"fmt"
	"reflect"
)

/*
 * map[keyType]valueType is implemented as a hash map in Go
 * Zero-value of map is *nil*
 * len function tells you the number of key-value pairs in a map
 * Nil maps: Read returns zero-value of valueType; write causes panic
 * Empty maps: Read returns zero-value of valueType; writes can be perfomed
 * Using *make*: map will have a length of 0; can grow past the initially specified size
 * Maps automatically grow as you add key-value pairs to them
 * maps are not comparable; Can only be compared with *nil*
 * key for a map can be any comparable type; Cannot use slice or map as key
 */
func main() {
	// Declaration and initializations
	// Zero-value of map is *nil*
	fmt.Println("\nNil map:")
	var nilMap map[string]int // nil map
	fmt.Println(nilMap, len(nilMap), reflect.TypeOf(nilMap), nilMap == nil)
	fmt.Println(nilMap["x"]) // Print 0; zero-value of the valueType *int*
	// nilMap["x"] = 12         // Panic!! assignment to entry in nil map

	// Using empty map literal
	fmt.Println("\nUsing empty map literal:")
	totalWins := map[string]int{} // Empty map
	fmt.Println(totalWins, len(totalWins), reflect.TypeOf(totalWins), totalWins == nil)
	fmt.Println(totalWins["x"]) // Print 0; zero-value of the valueType *int*
	totalWins["x"]++            // This works because a zero-value is returned even if key is absent
	fmt.Println(totalWins, totalWins["x"])

	// Non-empty map literal
	// Notice the trailing comma after the last key-value pair
	fmt.Println("\nUsing non-empty map literal:")
	teams := map[string][]string{
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}
	fmt.Println(teams, len(teams), reflect.TypeOf(teams), teams == nil)
	fmt.Println(teams["x"]) // Print []; zero-value of the valueType *[]string*
	fmt.Println(teams["Orcas"])

	// Using in-built make function
	fmt.Println("\nUsing make:")
	ages := make(map[int][]string, 10)                              // No values, but also not nil map
	fmt.Println(ages, len(ages), reflect.TypeOf(ages), ages == nil) // has a length of 0; can grow past the initially specified size

	// The comma-ok idiom
	// If ok is true, the key is present in the map.
	// If ok is false, the key is not present.
	fmt.Println("\nThe comma-ok idiom:")
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	fmt.Println(v, ok)
	v, ok = m["world"] // Returns the assigned value 0
	fmt.Println(v, ok)
	v, ok = m["goodbye"] // key not present; returns zero-value for int which is 0
	fmt.Println(v, ok)

	// Deletions in map; using in-built *delete* function
	fmt.Println("\nDeleting from map:")
	fmt.Println("Map: ", m)
	delete(m, "world")
	fmt.Println("After deleting 'world': ", m)
	delete(m, "goodbye") // No error
	delete(nilMap, "")   // No error

	// Using map as sets: *map[keyType]bool*
	fmt.Println("\nUsing map as sets:")
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}

	// Another implementation using struct as valueType
	// This has a memory advantage as structs don't occupy memory; bool occupies 1 byte
	// However, readability is affected as comma-ok idiom is required to check if element exists in the set
	intSet2 := map[int]struct{}{}
	for _, v := range vals {
		intSet2[v] = struct{}{}
	}
	if _, ok := intSet2[5]; ok {
		fmt.Println("5 is in the set")
	}
}
