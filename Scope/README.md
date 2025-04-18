# Why Local Scope Does Not Work Outside a Function in Go

## 🔒 What is Local Scope?

In Go, a variable declared **inside a function or block** is **local** to that function/block.

```go
func sayHello() {
    message := "Hello, world"
    fmt.Println(message) // ✅ Works here
}

fmt.Println(message) // ❌ Error: undefined
```

---

## 🔎 Why Can’t It Be Used Outside?

Because the **compiler enforces scope rules** during **compile time**.

**Behind the scenes:**
1. The compiler builds a **symbol table** (a map of variable names and where they are declared).
2. When you try to use `message` outside `sayHello()`, the compiler:
   - Checks if `message` is declared in the current or an outer scope.
   - If not found → **compile-time error**: `undefined: message`
3. The variable `message` is stored in memory (like the stack), **only while `sayHello()` is running**. When the function ends, that memory is **freed**.

---

## 🔎 Go Memory Model

Go uses a **stack and heap memory model**:

| Variable Type     | Where Stored   | Lifetime              |
|-------------------|----------------|------------------------|
| Local variables   | Stack          | Short (function call) |
| Global variables  | Data segment   | Long (entire program) |
| Heap allocations  | Heap           | Controlled by GC      |

---

## 📦 Where Tables (Maps), Slices, and Structs Are Stored

Some types in Go are **composite** and need **dynamic memory**:

| Data Type     | Stored Where      | Notes |
|---------------|-------------------|-------|
| **Map**       | Heap              | Created using `make`, automatically heap-allocated. |
| **Slice**     | Pointer on stack, backing array on heap | If you slice inside a function and return it, the backing array lives on the heap. |
| **Struct**    | Stack or Heap     | Small structs can be on stack; large structs passed via pointer are on heap. |
| **Array**     | Stack             | If used inside function and not returned. |
| **Pointer**   | Stack or Heap     | The pointer itself is on stack; it may point to a heap value. |

> 📌 Rule of Thumb:  
If a variable **escapes the function scope** (used outside), Go's **escape analysis** will allocate it on the **heap**.

---

## 🛡️ Why This is Good

1. **Encapsulation**: Keeps logic clean and local.
2. **Memory safety**: Prevents accessing garbage memory.
3. **Concurrency safety**: Avoids races between goroutines.

---

## ✅ If You Want to Use It Outside...

**Return the value:**

```go
func getMessage() string {
    message := "Hello from inside"
    return message
}

func main() {
    msg := getMessage()
    fmt.Println(msg) // ✅ Now works
}
```

**Or declare it outside:**

```go
var message string

func sayHello() {
    message = "Hello from function"
}

func main() {
    sayHello()
    fmt.Println(message) // ✅ Now works
}
```
