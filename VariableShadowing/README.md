# 🎭 GoLang Variable Shadowing – Complete Guide <a name="top"></a>

📺 [Watch Video Explanation](https://www.youtube.com/watch?v=TZNdcEHCxgY&t=13s&ab_channel=GoWithHabib)


---

## 🔍 What is Variable Shadowing?

**Variable shadowing** in Go occurs when a **variable declared in an inner scope (block or function)** has the **same name** as a variable declared in an **outer scope**. The inner variable **"shadows" or hides** the outer one within its own scope.

It doesn't override or change the outer variable — the outer one becomes **inaccessible** while the inner one is in scope.

---

## 🧠 How It Works

Go follows **lexical scoping**, meaning:
- When a variable is referenced, Go starts by looking in the **innermost scope**.
- If it doesn’t find the variable there, it checks **one level up**, and so on until it finds a match.

---

## 📦 Memory and Compiler Behavior

- Each scope gets its own **stack frame** during execution.
- Shadowed variables are stored in **separate memory locations**.
- The compiler resolves which variable to use **at compile time**, based on scope.

---

## 📄 Example of Shadowing

```go
package main

import "fmt"

var x = 100 // package-level variable

func main() {
    x := 50 // shadows the package-level x
    fmt.Println("Inside main:", x)

    if true {
        x := 10 // shadows main's x
        fmt.Println("Inside if block:", x)
    }

    fmt.Println("Back in main:", x)
}
```

### 🔍 Output:

```bash
Inside main: 50
Inside if block: 10
Back in main: 50
```

### 📘 Explanation:
- Global `x = 100` is shadowed in `main()` by `x := 50`.
- That is further shadowed in the `if` block by `x := 10`.
- Each variable is distinct and exists in a separate scope.

---

## ⚠️ Common Mistake: Unintentional Shadowing

```go
func updateCount(count int) {
    if true {
        count := count + 1 // OOPS! Declares new `count` and shadows parameter
        fmt.Println("Updated:", count)
    }
    fmt.Println("Original:", count) // count parameter unchanged
}
```

### ✅ Fix:

```go
func updateCount(count int) {
    if true {
        count = count + 1 // Updates the outer count
        fmt.Println("Updated:", count)
    }
    fmt.Println("Original:", count)
}
```

---

## ✅ When to Use Shadowing

✅ Acceptable:
- In **short inner blocks**, when you're doing calculations or condition-specific logic.

🚫 Avoid:
- Shadowing **function parameters**, **globals**, or **loop variables** unless absolutely needed — it can lead to **bugs** or **confusion**.

---

## 🔎 How Go Detects Shadowing

Go **allows shadowing**, but it will **warn you** about unused variables:

```go
x := 5
x := 10 // compiler error: no new variables on left side of :=
```

If all variables on the left of `:=` are already declared, Go expects at least **one new variable**.

---

## 🧪 Example: Function Parameter Shadowing

```go
func greet(name string) {
    fmt.Println("Outer name:", name)
    {
        name := "John" // shadows the parameter
        fmt.Println("Inner name:", name)
    }
    fmt.Println("Back to outer name:", name)
}
```

---

## 🧠 Go Behind the Scenes (Compiler + Stack)

- Each scope block is translated to its own **activation record** on the **stack**.
- The Go compiler maps each identifier to a **symbol table entry** based on scope.
- Inner scope variables don’t overwrite outer variables in memory — they just exist in a **deeper frame**.

---

## 📋 Summary Table

| Scope      | Shadowed | Memory Location | Scope Lifetime    |
|------------|----------|------------------|-------------------|
| Global     | No       | Heap             | Entire program    |
| Function   | Yes      | Stack            | Until function ends |
| Block      | Yes      | Stack            | Until block ends  |

---

## 📚 Best Practices

- ✅ Use meaningful, different variable names to avoid confusion.
- 🔄 Don't reuse variable names unless the scope is **tight** and obvious.
- 🔬 Be cautious when copying variables inside `if`, `for`, or `switch`.

---

## 🧠 Pro Tip

Use `go vet` and `golangci-lint` to detect potential shadowing issues early:

```bash
go vet ./...
golangci-lint run
```

---

## 📌 Conclusion

- Shadowing is a valid part of Go's scoping rules, but it should be used carefully.
- Understand that **shadowing hides but doesn’t replace** outer variables.
- Go compiles shadowed variables into **separate symbols and memory locations**, avoiding conflicts.
- Awareness and good naming prevent bugs caused by unexpected variable behavior.

