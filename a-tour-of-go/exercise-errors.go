package main

import "fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
	return fmt.Sprintf("Cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0
	
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2*z)
	}
	
	return z, nil
}

func PrintResult(result float64, err error) {
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

func main() {
	result, err := Sqrt(2)
	PrintResult(result, err)

	result, err = Sqrt(-2)
	PrintResult(result, err)
}