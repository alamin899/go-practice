package main
import "fmt"

var globalVar int = 10 // Global variable

func add(a int, b int) int {
    sum := a + b       // Local variable
    return sum
}

func main() {
    result := add(5, 15)
	result2 := add(globalVar, 10) // Using global variable
	fmt.Println("Result2:", result2)   // Output: 20
	fmt.Println("Result:", result) // Output: 20
}