package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// Get a greetinng message and print it.
	message := greetings.Hello("Yodai")
	fmt.Println(message)
}
