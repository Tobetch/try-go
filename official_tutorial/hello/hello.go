package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")

	// Call code in an external package
	fmt.Println(quote.Go())
}
