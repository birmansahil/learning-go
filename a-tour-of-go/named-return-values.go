package main

import "fmt"

func division(num int) (a, b int) {
	a = num / 2
	b = num / 5
	return
}

func main() {
	fmt.Println(division(5))
}