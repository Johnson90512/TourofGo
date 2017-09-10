package main

import (
	"fmt"
	"math"
)

func main() {
	// will not run because "pi" is lowercase and can't be accessed outside the package fmt.Println(math.pi)
	fmt.Println(math.Pi)
}