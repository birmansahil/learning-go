package main

import "fmt"

type Person struct {
	age int
	hasCar bool
}

var p = map[string]Person {
	"Sahil": Person {
		25, true,
	},
	"John": Person {
		47, false,
	},
}

func main() {
	fmt.Println(p)
	fmt.Println(p["John"])
}