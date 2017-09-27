package main

import (
	"fmt"
	"math"
)

func Newt(x float64) float64 {
	a, b := x,0.0
	iterations := 0
	for {
		iterations++
		a, b = a - ((a * a) - x) / (2 * a), a
		fmt.Println(a - b)
		if a - b < 0 {
			b = -(a - b)
		}
		if (a - b) < 1e-8 {
			break
		}
	}
	fmt.Println("Iterations: ", iterations)
	return a
}

func main() {
	fmt.Println(Newt(13))
	fmt.Println("Expected: ", math.Sqrt(19283389217398127))
}