package main

import "fmt"

type Vertex struct {
	a int
	b int
}

func main() {
	v := Vertex{1, 2}

	p := &v
	fmt.Println(*p)
	fmt.Println((*p).a)

	p.a = 5
	fmt.Println(v)
}