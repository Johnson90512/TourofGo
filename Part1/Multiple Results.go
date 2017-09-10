package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}
//will return opposite of what is input.
func main() {
	a,b := swap("test", "program")
	fmt.Println(a,b)
}