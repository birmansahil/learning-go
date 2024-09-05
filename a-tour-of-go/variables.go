package main

import "fmt"

// Declaring variables at package level
var a, b, c bool

func main() {
	// Declaring variable at function level
	var x int

	fmt.Println(a, b, c, x)
}