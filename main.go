package main

import (
	"fmt"
	"test.com/custompackage"
)

func main() {
	var sum int = custompackage.Add(5, 10)
	// Print the sum
	fmt.Println("Sum:", sum)
	fmt.Println("Hello")
}
