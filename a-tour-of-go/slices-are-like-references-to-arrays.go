package main 

import "fmt"

func main() {
	names := [6] string {
		"John",
		"Chris",
		"Alex",
		"Jessica",
		"Anna",
		"Emma",
	}

	fmt.Println(names)

	males := names[0 : 3]
	females := names[3:6]

	fmt.Println(males, females)

	males[0] = "Sahil"
	fmt.Println(males)
	fmt.Println(names)
}