package main 

import "fmt"

type Person struct {
	age int
	hasCar bool
}

var m = map [string] Person {
	"Sahil": {25, true},
	"John": {48, false},
}

func main() {
	fmt.Println(m)
}