package main

import (
	"fmt"

	"example.com/greetings"

	"rsc.io/quote"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)

	// Call code in an external package
	fmt.Println(quote.Go())
}
