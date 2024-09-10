package main 

import (
	"fmt"
	"math"
)

func pow(a, b, c float64) float64 {
	if z := math.Pow(a, b); z < c {
		return z
	}

	return c
}

func main() {
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}