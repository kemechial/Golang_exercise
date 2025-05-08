package calculator

import (
	"fmt"
	"math"
)

func Add(a, b int) int {
	fmt.Println("Sum is:", a+b)
	return a + b
}

func Subtract(a, b int) int {
	return int(math.Abs(float64(a - b)))
}

