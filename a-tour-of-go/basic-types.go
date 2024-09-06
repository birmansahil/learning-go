package main

import (
	"fmt"
	"math/cmplx"
)

var (
	hasDiscount bool = true
	discount int = 10
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", hasDiscount, hasDiscount)
	fmt.Printf("Type: %T Value: %v\n", discount, discount)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}