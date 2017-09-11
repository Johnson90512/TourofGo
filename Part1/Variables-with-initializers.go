package main

import "fmt"

var i, j int = 1, 2
//types can be omitted if an initializer is present and the variable will take the type of the initializer
func main() {
	var c, python, java = true, false, "No!"
	fmt.Println(i, j, c, python, java)
}
