package main

import (
	"fmt"
)

func main() {
	fmt.Println("Golang Testing Tutorial")
	result := Calculate(2)
	fmt.Println(result)
}

// Function that adds 2 to any number given and return the sum.
func Calculate(number int) int {
	val := 2 + number
	return val
}
