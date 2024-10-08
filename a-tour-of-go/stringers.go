package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
 
func main() {
	a := Person{"Alex Morre", 36}
	b := Person{"Karima Taos", 67}
	fmt.Println(a, b)
}