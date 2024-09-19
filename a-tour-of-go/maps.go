package main 

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// var m map[string]Vertex

func main() {
	m := make(map[string]Vertex)

	m["Bell"] = Vertex {
		40.6843, -74.39967,
	}

	m["Telus"] = Vertex {
		20.2015, 68.9567,
	}

	fmt.Println(m)
	fmt.Println(m["Bell"])
	fmt.Println(m["Telus"])
}