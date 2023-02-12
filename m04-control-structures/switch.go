package main

import "fmt"

/*
 * You can switch on any type that can be compared with ==
 * Thus, it cannot work with slices, maps, channels, functions, and structs that contain fields of these types
 * separate multiple matches with commas. Eg: case 1,2,3
 * In Go, an empty case means nothing happens
 * By default, cases in switch statements in Go don’t fall through; break statement not required
 * Go does include a fallthrough keyword, which lets one case continue on to the next one
 * Blank switch: a switch statement that doesn’t specify the value that you’re comparing against
 * A blank switch allows you to use any boolean comparison for each case
 * Favor blank switch statements over if/else chains when you have multiple related cases
 */
func main() {
	fmt.Println("\nA sample switch-case:")
	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}

	fmt.Println("\nBreaking from the switch-case(unexpected output):")
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "is even")
		case i%3 == 0:
			fmt.Println(i, "is divisible by 3 but not 2")
		case i%7 == 0:
			fmt.Println("exit the loop!")
			break // !!This breaks only the switch-case; not the for loop
		default:
			fmt.Println(i, "is boring")
		}
	}

	// Using labels with break
	fmt.Println("\nUsing labels within switch-case to break:")
loop:
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "is even")
		case i%3 == 0:
			fmt.Println(i, "is divisible by 3 but not 2")
		case i%7 == 0:
			fmt.Println("exit the loop!")
			break loop // This breaks the for loop
		default:
			fmt.Println(i, "is boring")
		}
	}

	// Blank switch statement
	fmt.Println("\nThe blank switch statement")
	words = []string{"hi", "salutations", "hello"}
	for _, word := range words {
		wordLen := len(word)
		switch {
		case wordLen < 5:
			fmt.Println(word, "is a short word!")
		case wordLen > 10:
			fmt.Println(word, "is a long word!")
		default:
			fmt.Println(word, "is exactly the right length.")
		}
	}

}
