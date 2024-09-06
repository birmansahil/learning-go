package main

import "fmt"

const (
	a = 12.45
	b = 2
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", a, a)
	fmt.Printf("Type: %T Value: %v\n", b, b)
}