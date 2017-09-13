package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		//for loop add 9+8+7+6+5+4+3+2+1 = 45
	}
	fmt.Println(sum)
}

