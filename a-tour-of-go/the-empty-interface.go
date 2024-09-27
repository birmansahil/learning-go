package main 

import "fmt"

func main() {
	var i interface {}
	describe(i)
	PrintAnything(i)

	i = 42
	describe(i)
	PrintAnything(i)

	i = "hello"
	describe(i)
	PrintAnything(i)
}

func describe(i interface{}) {
	fmt.Printf("%v %T\n", i, i)
}

func PrintAnything(i interface{}) {
	fmt.Println(i)
}