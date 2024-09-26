package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof! I am " + d.Name
}

func main() {
	p := Person{"Sahil"}
	d := Dog{"Tommy"}

	fmt.Println(p.Speak())
	fmt.Println(d.Speak())
}