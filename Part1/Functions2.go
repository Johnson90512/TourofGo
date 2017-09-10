package main

import "fmt"

//parameters x and y share type int, so it can be omitted on all but the last parameter
func add1(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add1(546, 242))
}
