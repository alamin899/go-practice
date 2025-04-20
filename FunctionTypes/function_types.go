package main

import "fmt"

// --- Custom Type Declaration ---
type Operation func(int, int) int

// Named function for custom type example
func multiply(a, b int) int {
	return a * b
}

// --- Example 1: Assigning function to a variable ---
func assignToVariableExample() {
	var op func(int, int) int
	op = func(a, b int) int {
		return a + b
	}
	fmt.Println("Assign to variable:", op(3, 4)) // Output: 7
}

// --- Example 2: Passing function as an argument ---
func operate(a, b int, fn func(int, int) int) int {
	return fn(a, b)
}

func passFunctionAsArgumentExample() {
	sum := operate(5, 3, func(x, y int) int {
		return x + y
	})
	fmt.Println("Function as argument:", sum) // Output: 8
}

// --- Example 3: Returning a function ---
func greeter(name string) func() string {
	return func() string {
		return "Hello, " + name
	}
}

func returnFunctionExample() {
	sayHi := greeter("GoLang")
	fmt.Println("Return function:", sayHi()) // Output: Hello, GoLang
}

// --- Example 4: Storing functions in a slice ---
func functionSliceExample() {
	operations := []func(int, int) int{
		func(a, b int) int { return a + b },
		func(a, b int) int { return a - b },
	}
	fmt.Println("Slice[0]:", operations[0](10, 5)) // Output: 15
	fmt.Println("Slice[1]:", operations[1](10, 5)) // Output: 5
}

// --- Example 5: Higher-order function ---
func apply(f func(int) int, value int) int {
	return f(value)
}

func square(x int) int {
	return x * x
}

func higherOrderFunctionExample() {
	result := apply(square, 4)
	fmt.Println("Higher-order function:", result) // Output: 16
}

// --- Example 6: Using custom function type ---
func applyOp(a, b int, op Operation) int {
	return op(a, b)
}

func customFunctionTypeExample() {
	fmt.Println("Custom function type:", applyOp(3, 4, multiply)) // Output: 12
}

func main() {
	assignToVariableExample()
	passFunctionAsArgumentExample()
	returnFunctionExample()
	functionSliceExample()
	higherOrderFunctionExample()
	customFunctionTypeExample()
}
