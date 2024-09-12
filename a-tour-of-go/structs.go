package main

import "fmt"

type Vertex struct {
	a int
	b int
	c string
}

func main() {
	fmt.Println(Vertex{2, 4, "Hello"})
}