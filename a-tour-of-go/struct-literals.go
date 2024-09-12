package main

import "fmt"

type Person struct {
	fname, lname string 
} 

var (
	p1 = Person{"Sahil", "Birman"}
	p2 = Person{fname: "John", lname: "Doe"}
	p3 = Person{}
	p4 = &Person{"Shalender", "Singh"}
)

func main() {
	fmt.Println(p1, p2, p3, p4)
}