package main

import "fmt"

func reverse(a int, b string) (string, int) {
	return b, a
}

func main() {
	fmt.Println(reverse(4, "Number"))
}