package main

import (
	"fmt"
	"math"
)

func sqrt(a float64) string {
	if a < 0 {
		return sqrt(-a) + "i"
	}

	return fmt.Sprint(math.Sqrt(a))
}

func main() {
	fmt.Println(sqrt(4), sqrt(-4))
}