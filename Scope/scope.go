package main

import "fmt"

// Global variable (accessible throughout the package)
var globalVar = "I am a global variable"

func main() {
    // Function scope (accessible only within this function)
    localVar := "I am a local variable"
    fmt.Println(globalVar) // Accessing global variable
    fmt.Println(localVar)  // Accessing local variable

    // Block scope (accessible only within the block)
    if true {
        blockVar := "I am a block variable"
        fmt.Println(blockVar)
    }

    // Loop scope (accessible only within the loop)
    for i := 0; i < 1; i++ {
        loopVar := "I am a loop variable"
        fmt.Println(loopVar)
    }

    // Anonymous function scope (accessible only within the anonymous function)
    anonFunc := func() {
		anonVar := "I am an anonymous function variable"
        fmt.Println(anonVar)
    }
    anonFunc()

    // Closure scope (variable captured by a closure)
    closureVar := "I am a closure variable"
    closureFunc := func() {
        fmt.Println(closureVar) // Accessing closure variable
    }
    closureFunc()
}