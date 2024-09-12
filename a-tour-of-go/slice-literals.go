package main

import "fmt"

func main() {
	q := [] int {0, 2, 4, 6, 8}
	fmt.Println(q)

	r := [] bool {true, true, false, true, false}
	fmt.Println(r)

	s := [] struct {
		name string
		age int
		hasCar bool
	} {
		{"Chris", 52, true},
		{"Alex", 35, false},
		{"Hans", 20, false},
	}
	fmt.Println(s)

	for i:=0; i < 3; i++ {
		fmt.Println(s[i].name)
	}
}