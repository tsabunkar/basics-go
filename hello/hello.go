package main

import (
	"fmt"

	"example.com/greetings" // From the ./greetings/ module
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Tejas Sabunkar")
    fmt.Println(message)
}
