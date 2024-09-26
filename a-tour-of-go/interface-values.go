package main

import(
	"fmt"
	"math"
)

type I interface {
	Message()
}

type T struct {
	Say string
}

func (t *T) Message() {
	fmt.Println(t.Say)
}

type F float64 

func (f F) Message() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	i.Message()
	describe(i)
	
	i = F(math.Pi)
	i.Message()
	describe(i)
}

func describe(i I) {
	fmt.Printf("%v, %T\n", i, i)
}