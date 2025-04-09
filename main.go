package main

import (
	"ExampleAntlr/listener"
	"fmt"
)

// main is the entry point of the application.
// It evaluates a mathematical expression using the Calc function from the listener package
// and prints the result to the standard output.
//
// The expression "3 * 5 + 2 * 10" is evaluated, and the expected result is 35.
func main() {
	fmt.Println(listener.Calc("3 * 5 + 2 * 10"))
}
