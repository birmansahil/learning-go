package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("length=%d capacity=%d %v\n", len(s), cap(s), s)
}

func main() {
	var s [] int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 4, 5)
	printSlice(s)
}