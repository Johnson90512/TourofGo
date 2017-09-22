package main

import (
	"fmt"
	//"math"
)

func Newt(x float64) float64 {

	a := 1.0
	for i :=1; i < 10; i++ {
		a = a - ((a*a) - x) / (2 * a)
		fmt.Println(a)
	}
	return a
}

func main() {
	fmt.Println(Newt(100))
}