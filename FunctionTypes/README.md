# 🧠 Go Function Types – Complete Guide

---

## 📘 What Are Function Types?

In Go, **function types** define the *signature* of a function — including its parameters and return values.

They allow functions to be:
- Assigned to variables
- Passed as parameters
- Returned from other functions
- Stored in slices, maps, and structs

```go
func(int, int) int
```

This is a function type that takes two `int`s and returns an `int`.

---

## 🔤 Syntax of Function Types

```go
func(parameter1Type, parameter2Type, ...) returnType
```

### Examples

```go
func()                    // No parameters, no return
func(int) string          // Takes int, returns string
func(a int, b int) int    // Two int parameters, returns int
```

---

## 🎯 Using Function Types

### 1. Assigning Function to a Variable

```go
var op func(int, int) int

op = func(a, b int) int {
    return a + b
}

fmt.Println(op(3, 4)) // Output: 7
```

---

### 2. Passing Function as Argument

```go
func operate(a, b int, fn func(int, int) int) int {
    return fn(a, b)
}

func main() {
    sum := operate(5, 3, func(x, y int) int {
        return x + y
    })
    fmt.Println(sum) // Output: 8
}
```

---

### 3. Returning a Function

```go
func greeter(name string) func() string {
    return func() string {
        return "Hello, " + name
    }
}

func main() {
    sayHi := greeter("GoLang")
    fmt.Println(sayHi()) // Output: Hello, GoLang
}
```

---

### 4. Storing in a Slice

```go
func main() {
    operations := []func(int, int) int{
        func(a, b int) int { return a + b },
        func(a, b int) int { return a - b },
    }

    fmt.Println(operations[0](10, 5)) // Output: 15
    fmt.Println(operations[1](10, 5)) // Output: 5
}
```

---

## 🔁 Higher-Order Functions

A **higher-order function** is a function that:
- Takes one or more functions as parameters
- Returns a function

### Example:

```go
func apply(f func(int) int, value int) int {
    return f(value)
}

func square(x int) int {
    return x * x
}

func main() {
    result := apply(square, 4)
    fmt.Println(result) // Output: 16
}
```

---

## 🧠 Behind the Scenes: Function Types in Go

| Feature               | Behavior                                                           |
|----------------------|---------------------------------------------------------------------|
| Named Functions       | Stored in the code section; address resolved at compile time       |
| Anonymous Functions   | Allocated at runtime; pointer stored in variable                   |
| Closures              | Functions that "capture" external variables and store them in heap |
| Function Types        | Describes the input/output of functions, used to declare variables |

---

### 🔍 Memory & Compiler Notes

- **Function literals** and **closures** may escape to the **heap** if stored beyond the stack scope.
- Use `go build -gcflags="-m"` to analyze escape behavior.

---

## 🧪 Custom Function Type Declarations

```go
type Operation func(int, int) int

func multiply(a, b int) int {
    return a * b
}

func applyOp(a, b int, op Operation) int {
    return op(a, b)
}

func main() {
    fmt.Println(applyOp(3, 4, multiply)) // Output: 12
}
```

---

## ⚠️ Notes on Function Equality

- You **cannot compare** two function values unless you are comparing them to `nil`.
  
```go
var f1 func()
var f2 func()

fmt.Println(f1 == f2) // ❌ Compilation error
```

---

## ✅ Summary

| Concept                  | Description                                                  |
|--------------------------|--------------------------------------------------------------|
| Function Type            | Signature of parameters and return values                    |
| Higher-Order Function    | Function that accepts/returns another function               |
| Anonymous Function       | Function without a name, often used inline                   |
| Closure                  | Function that captures variables from outer scope            |
| Custom Type Alias        | Use `type` to define reusable function type names            |
| Function as Value        | Store function in variables, pass to or return from functions|

---

## 📚 References

- [Go Language Spec - Function Types](https://golang.org/ref/spec#Function_types)
- [Effective Go - Functions](https://golang.org/doc/effective_go#functions)
- [Go Blog - Closures and Functions](https://go.dev/blog/closures)

---
