package main

import (
	"fmt"
	"math"
)

func Newt(x float64) float64 {

	a := 1.0
	for i :=1; i < 15; i++ {
		a = a - ((a*a) - x) / (2 * a)
	}
	return a
}

func main() {
	fmt.Println(Newt(256))
	fmt.Println(math.Sqrt(256))
}