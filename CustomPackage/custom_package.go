package custompackage

import "fmt"

func Add(a int, b int) int {
	var sum int = a + b
	fmt.Println("Sum:", sum)
	return sum
}