package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("length = %d capacity = %d %v\n", len(s), cap(s), s)
}

func main() {
	s := [] int {2, 3, 5, 7, 11, 13}
	printSlice(s)

	// zero length
	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)
}