package main

import "fmt"

const Pi = 3.14
//setting constants and returning them
func main() {
	const World = "world"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}