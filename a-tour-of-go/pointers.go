package main

import "fmt"

func main() {
	i, j := 12, 50

	p := &i
	fmt.Println(p)
	fmt.Println(*p)

	*p = 6
	fmt.Println(i)

	p = &j
	fmt.Println(*p)

	*p = *p + 25
	fmt.Println(j)
}