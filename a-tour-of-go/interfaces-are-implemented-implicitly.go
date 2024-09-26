package main

import "fmt"

type I interface {
	Message()
}

type T struct {
	Say string
}

func (t T) Message() {
	fmt.Println(t.Say)
}

func main() {
	var i I = T{"Hello"}
	i.Message()

	t := T{"Hello again"}
	t.Message() 
}

