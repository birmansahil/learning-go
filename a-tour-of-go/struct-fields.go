package main

import "fmt"

type Vertex struct {
	a int
	b int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v.a)

	v.a = 4
	fmt.Println(v.a)
}