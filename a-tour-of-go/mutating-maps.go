package main 

import "fmt"

func main() {
	m := make(map [string] int)

	m["age"] = 24
	fmt.Println("The age is:", m["age"])

	m["age"] = 35
	fmt.Println("The age is:", m["age"])

	v, ok := m["age"]
	fmt.Println("The age is:", v, "Is present?", ok)

	delete(m, "age")
	fmt.Println("The age is:", m["age"])

	v, ok = m["age"]
	fmt.Println("The age is:", v, "Is present?", ok)
}