# 🧠 Deep Dive: Go Functions

---

## 🔧 1. Basic Function Syntax

```go
func functionName(param1 type1, param2 type2) returnType {
    // function body
}
```

Example:

```go
func add(a int, b int) int {
    return a + b
}
```

---

## 🔁 2. Multiple Return Values

```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

---

## 🧠 3. Named Return Values

```go
func fullName() (first string, last string) {
    first = "John"
    last = "Doe"
    return // auto returns first and last
}
```

---

## 📦 4. Variadic Functions

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

---

## 🧬 5. Functions as First-Class Citizens

```go
func greet(name string) string {
    return "Hello, " + name
}

var f func(string) string = greet
fmt.Println(f("Go"))
```

---

## 🔁 6. Closures

```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

c := counter()
fmt.Println(c()) // 1
fmt.Println(c()) // 2
```

---

## 🎯 7. Anonymous Functions

```go
func() {
    fmt.Println("I’m anonymous!")
}()
```

```go
hello := func(name string) {
    fmt.Println("Hello", name)
}
hello("Go Dev")
```

---

## 🧵 8. Function Pointers

```go
func apply(fn func(int) int, val int) int {
    return fn(val)
}

func square(x int) int {
    return x * x
}

fmt.Println(apply(square, 5)) // 25
```

---

## 📌 9. Defer in Functions

```go
func demo() {
    defer fmt.Println("This prints last")
    fmt.Println("This prints first")
}
```

Deferred calls use **LIFO** order.

---

## 📍 10. Recover & Panic

```go
func safeDivide(a, b int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered:", r)
        }
    }()
    if b == 0 {
        panic("division by zero")
    }
    fmt.Println(a / b)
}
```

---

## ⚙️ 11. Function Call Internals

- Arguments evaluated **left to right**
- New **stack frame** is created
- Locals stored in stack
- Stack is popped on return

---

## 📊 12. Escape Analysis

```go
func createPointer() *int {
    x := 10
    return &x // Escapes to heap
}
```

---

## ✅ 13. Function Best Practices

- Keep functions small & focused
- Use named returns **sparingly**
- Avoid unnecessary `defer`
- Use interfaces for testability

---

## 🧪 14. Unit Testing Functions

```go
func Add(a, b int) int {
    return a + b
}
```

```go
// add_test.go
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Expected 5, got %d", result)
    }
}
```

---

## 📦 15. Higher-Order Functions

```go
func mathOperation(op string) func(int, int) int {
    switch op {
    case "add":
        return func(a, b int) int { return a + b }
    case "mul":
        return func(a, b int) int { return a * b }
    default:
        return nil
    }
}
```

---

Want even deeper content like **inline optimizations**, **generics**, **function tracing**, or **runtime performance analysis**? Just ask!
