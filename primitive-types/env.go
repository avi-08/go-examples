package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Listing all environment variables\n")
	fmt.Println(os.Environ())

	name := os.Getenv("USER")
	fmt.Println("\nAhoy", name, "! Welcome aboard")

	fmt.Println("KUBECONFIG: ", os.Getenv("KUBECONFIG"))
}
