package main

import (
	"fmt"
	"math"
)

func main() {
	var a int = 25
	var b float64 = math.Sqrt(float64(a))
	var c uint = uint(b)

	fmt.Println(a, b, c)
}