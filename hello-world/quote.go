package main

// Run go get rsc.io/quote before executing this.
import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
	fmt.Println(quote.Glass())
}
