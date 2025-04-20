# 📦 GoLang Custom Package: Complete Guide

---

## 🔰 What is a Package in Go?

In Go, a **package** is a way to organize and reuse code. A Go program is made up of **packages**, and every Go file starts with a `package` declaration. You can create **custom packages** to modularize your code and make it easier to manage and reuse across different parts of your project.

---

## 🛠️ Steps to Create and Use a Custom Package

---

### 1. 🧱 Create the Project Structure

```bash
go-custom-pkg/
│
├── main.go             # Main application file
└── myutils/            # Custom package directory
    └── mathutils.go    # Custom package file
```

---

### 2. 📝 Create the Custom Package

**File:** `myutils/mathutils.go`

```go
package myutils

import "fmt"

// Exported function (starts with a capital letter)
func Add(a int, b int) int {
    return a + b
}

// Exported function
func Multiply(a int, b int) int {
    return a * b
}

// Unexported function (only available inside the package)
func subtract(a int, b int) int {
    return a - b
}

// Exported function calling unexported function
func PrintDifference(a, b int) {
    result := subtract(a, b)
    fmt.Println("Difference is:", result)
}
```

> 🔒 Note: Only identifiers (functions, variables, types) that **start with an uppercase letter** are **exported** (public). Lowercase ones are **unexported** (private).

---

### 3. 📥 Use the Custom Package in `main.go`

**File:** `main.go`

```go
package main

import (
    "fmt"
    "go-custom-pkg/myutils"  // Importing the custom package
)

func main() {
    sum := myutils.Add(5, 3)
    product := myutils.Multiply(5, 3)

    fmt.Println("Sum:", sum)
    fmt.Println("Product:", product)

    myutils.PrintDifference(10, 4)
}
```

---

### 4. 🏃‍♂️ Run the Program

Navigate to your project folder and run:

```bash
go run main.go
```

Expected output:

```bash
Sum: 8
Product: 15
Difference is: 6
```

---

## 🔍 How It Works Behind the Scenes

---

### 🧠 Compilation

- Go **compiles each package separately** into `.a` files (archives).
- The compiler creates an internal symbol table of all **exported identifiers**.
- During compilation of `main.go`, the compiler links it with the `myutils` package.

---

### 📦 Import Resolution

- Go uses the **import path** (`go-custom-pkg/myutils`) to locate the package.
- Imports are resolved from:
  - The current module
  - `GOPATH`
  - External modules (`go mod`)

---

### 📂 Module-Aware Mode (`go mod`)

If you're using **Go modules**, make sure to initialize the module:

```bash
go mod init go-custom-pkg
```

> This creates a `go.mod` file and allows Go to treat the folder as a module.

---

### 🧹 Visibility and Access

- **Exported identifiers** (capitalized) are accessible from other packages.
- **Unexported identifiers** (lowercase) are private to the package.
- You can wrap unexported logic inside exported functions to expose limited behavior (like `PrintDifference` above).

---

## ✅ Best Practices for Custom Packages

- Keep functions logically grouped (e.g., all math operations in one package).
- Keep files short and focused.
- Use meaningful names for packages and exported identifiers.
- Avoid circular imports (Go doesn't allow them).
- Use Go’s `doc` tool to generate package documentation:
  
```bash
go doc myutils
```

---

## 🔄 Reusability Example

You can import your custom package in **multiple Go files** across the same module:

```go
import "go-custom-pkg/myutils"
```

---

## 🗃️ Summary

| Concept                | Description                                      |
|------------------------|--------------------------------------------------|
| `package`              | Keyword to define a file's package               |
| Custom Package         | A folder with `.go` files under a named package  |
| Exported Function      | Function name starts with a capital letter       |
| Unexported Function    | Starts with lowercase, private to the package    |
| `import`               | Brings other packages into the current file      |
| `go mod init`          | Initializes a module-aware project               |
| `go run main.go`       | Compiles and runs the project                    |

---

## 🧪 Advanced: Nesting & Sub-packages

You can nest packages like:

```
utils/
├── math/
│   └── calc.go
└── string/
    └── text.go
```

Then import like:

```go
import "go-custom-pkg/utils/math"
```

---

Let me know if you’d like this exported as an actual `.md` file or want to explore how packages work with third-party libraries!
