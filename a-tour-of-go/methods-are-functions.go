package main 

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

func abs(v Vertex) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	v := Vertex{12, 5}
	fmt.Println(abs(v))
}