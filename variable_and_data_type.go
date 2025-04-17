package main

import "fmt"

func main() {
	/**
	We use data type only
	int
	float32
	bool
	string
	*/

	//Declare and initialize variable

	//int type data different way
	var x int = 10
	a := 20

	//float32 type data different way
	var y float32 = 10.20
	y = 9.99

	fmt.Println(x, a, y)
}


/**
✅ Basic Types
	var a int = 10          // Integer type (default size based on system)
	var b float64 = 3.14    // 64-bit floating point number
	var c bool = true       // Boolean value (true or false)
	var d string = "Hello"  // A string (sequence of characters)
	var e byte = 'A'        // Byte is alias for uint8 (represents ASCII characters)
	var f rune = '😊'       // Rune is alias for int32 (for Unicode characters)

✅ Composite Types
	// Array: Fixed-size list of integers
	var arr [3]int = [3]int{1, 2, 3}


	// Slice: Dynamic-size list of integers
	var slice []int = []int{4, 5, 6}


	// Struct: Custom data type with named fields
	type Person struct {
		Name string
		Age  int
	}
	var p = Person{Name: "John", Age: 30}  // Creating a struct instance


	// Map: Key-value pair collection
	var m map[string]int = map[string]int{"a": 1, "b": 2}


	// Pointer: Holds memory address of a variable
	var x int = 5
	var ptr *int = &x   // ptr now points to x


	// Function: First-class citizen in Go
	func greet(name string) string {
		return "Hello " + name
	}


	// Channel: Used for communication between goroutines
	ch := make(chan int)
	go func() {
		ch <- 42        // Sending data to channel
	}()


	// Interface: Defines a set of methods
	type Speaker interface {
		Speak() string
	}


✅ Custom Types
type MyInt int        // Define a new type based on int
var num MyInt = 25    // Create a variable of type MyInt

*/