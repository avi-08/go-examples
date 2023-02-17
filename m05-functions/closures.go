package main

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

/*
 * Closures are functions defined within a function
 * All values accessible to parent are accessible to the closure as well
 * In Go, the cleanup code is attached to the function with the defer keyword
 * Go allows to defer multiple closures in a function
 * They run in last-in-first-out order; the last defer registered runs first
 */
func main() {
	// Passing closure to a function
	fmt.Println("\nClosures can be passed to a function:")
	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println(people)

	sort.Slice(people, func(i int, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)

	sort.Slice(people, func(i int, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)

	// Returning funtion from function
	fmt.Println("\nFunctions can also be returned from a function:")
	twoBase := makeMult(2)
	threeBase := makeMult(3)
	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}

	// Using defer for resource cleanup
	fmt.Println("\nUsing defer for resource cleanup:")
	deferExample()

	fmt.Println("\nAnother defer example:")
	deferExample2()

	fmt.Println("\nThe resource cleanup return pattern:")
	getFileExample()

}

// Returning function from a function
func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func deferExample() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	// defer delays the invocation until the surrounding function exits
	// The code within defer closures runs after the return statement
	// any variables passed into a deferred closure aren’t evaluated until the closure runs
	defer f.Close()
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}

// Return values from deferred functions cannot be read
func deferExample2() {
	defer func() int {
		fmt.Println(" running defer!")
		return 2 // there's no way to read this value
	}()
	fmt.Print("defer runs after this ->")
}

// deferred functions can read return values of the surrounding function
// This can be done using named return values
func DoSomeInserts(ctx context.Context, db *sql.DB, value1, value2 string) (err error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	// defer is called after DoSomeInserts returns
	// since err is the named return value, it can be accessed from within
	// defer block
	defer func() {
		if err == nil {
			err = tx.Commit()
		}
		if err != nil {
			tx.Rollback()
		}
	}()
	_, err = tx.ExecContext(ctx, "INSERT INTO FOO (val) values $1", value1)
	if err != nil {
		return err
	}
	// use tx to do more database inserts here
	return nil
}

// A common pattern in Go is for a function that allocates a resource to also return
// a closure that cleans up the resource
func getFile(name string) (*os.File, func(), error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}
	return file, func() {
		fmt.Println("Closing file", file.Name())
		file.Close()
	}, nil
}

func getFileExample() {
	_, closer, err := getFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Didn't really want to read the file!")
	defer closer()
}
