package main

import "fmt"

func deferTest() {
	defer fmt.Println("Print second")
	fmt.Println("Print first")
}

func main() {
	defer fmt.Println("by Sahil :)")
	deferTest()
	fmt.Println("Hello")
	fmt.Println("World")
}