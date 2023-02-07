package main

import (
	"fmt"
	"reflect"
)

/*
 * Structs group together related data
 */
func main() {
	// All values will be assigned zero-value of their type
	var fred person
	fmt.Println(fred)

	// Using empty struct literal
	bob := person{}
	fmt.Println(bob)

	// Using struct literal
	// All values required and must correspond to the order of fields
	alice := person{
		"alice",
		38,
		"al",
	}
	fmt.Println(alice, reflect.TypeOf(alice))

	// Using struct literal with field name-values
	// Can specify some fields; Remaining values are assigned zero-value
	julius := person{
		name: "julius",
		age:  40,
	}
	fmt.Println(julius, julius.age) // Access values using *dot*

	// Anonymous structs
	// Useful for mashalling/unmarshalling: JSON / protocol buffers
	// Useful while writing tests
	fmt.Println("\nAnonymous structs:")
	var person struct {
		name string
		age  int
		pet  string
	}
	person.name = "bob"
	person.age = 50
	person.pet = "dog"
	fmt.Println(person, reflect.TypeOf(person), person.name)

	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}
	fmt.Println(pet, reflect.TypeOf(pet), pet.kind)

	// Comparable structs
	// Structs that are entirely composed of comparable types are comparable
	// Structs containing slice or map fields are not comparable
	// Go provides no magic method that can be overridden to redefine equality
	fmt.Println(fred == bob, fred == alice)

	// Conversions
	fmt.Println("\nStruct comparisons:")
	var p1 firstPerson
	var p2 secondPerson
	var p3 thirdPerson
	var p4 fourthPerson
	fmt.Println(p1, reflect.TypeOf(p1))
	fmt.Println(p2, reflect.TypeOf(p2))
	fmt.Println(p3, reflect.TypeOf(p3))
	fmt.Println(p4, reflect.TypeOf(p4))

	p12 := firstPerson(p2)
	p12.name = "Groot"
	// Cannot compare p1 == p2; But p12 can be compared with p1
	fmt.Println(p1, p12, p12 != p1)

	// Cannot convert p1 and p3 as fields are in different order
	// Cannot convert p1 and p4 as fields have different names
	// Cannot convert p1 and p5 as there is one additional field

	// Assignment and Comparisons with anonymous struct
	fmt.Println("\nAssignment and Comparisons with anonymous struct:")
	f := firstPerson{
		name: "Bob",
		age:  50,
	}
	var g struct {
		name string
		age  int
	}
	// compiles -- can use = and == between identical named and anonymous structs
	g = f
	fmt.Println(f == g)
}

type person struct {
	name string
	age  int
	pet  string
}

type firstPerson struct {
	name string
	age  int
}

type secondPerson struct {
	name string
	age  int
}

type thirdPerson struct {
	age  int
	name string
}

type fourthPerson struct {
	firstName string
	age       int
}

type fifthPerson struct {
	name          string
	age           int
	favoriteColor string
}
