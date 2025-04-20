package main
import "fmt"
func main() {
	// Declare a variable with the same name in the same scope
	var x int = 10
	fmt.Println("Value of x:", x) // Output: Value of x: 10

	// Shadowing the variable x
	var x int = 20
	fmt.Println("Shadowed value of x:", x) // Output: Shadowed value of x: 20

	// Using a different scope (inside a function)
	{
		var x int = 30
		fmt.Println("Inner scope value of x:", x) // Output: Inner scope value of x: 30
	}

	fmt.Println("Outer scope value of x:", x) // Output: Outer scope value of x: 20
}