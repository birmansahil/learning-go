package main

import "fmt"

func multiply(a, b, c int) int {
	return a * b * c
}

func main() {
	fmt.Println(multiply(2,3,4))
}