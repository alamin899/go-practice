# 🚀 How to Run a Go Project

Running a Go project is quite simple once you have the Go environment set up. Follow the steps below to start running a Go project.

---

## 📥 1. Install Go

Make sure that Go is installed on your system. To install Go, visit the official Go website and follow the installation instructions:

- [Download Go](https://golang.org/dl/)

Once installed, verify the installation by running:

```bash
go version
```

This should display the Go version installed.

---

## 📝 2. Set Up Your Go Project

- Create a project directory on your local machine.

```bash
mkdir my-go-project
cd my-go-project
```

- Inside the project directory, create a Go file (for example, `main.go`).

```go
// main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

---

## ⚙️ 3. Run the Project

To run the Go project, you can use the `go run` command.

```bash
go run main.go
```

This command compiles and runs the Go file directly.

---

## 🛠️ 4. Build the Project (Optional)

If you want to **compile the Go code** into an executable file, you can use the `go build` command.

```bash
go build main.go
```

This will generate an executable file named `main` (or `main.exe` on Windows) in the current directory. You can then run it as follows:

```bash
./main  # On Linux/macOS
main.exe # On Windows
```

---

## 🌍 5. Running a Go Project with Modules

If your project uses **Go modules**, you should initialize a Go module for dependency management:

1. Inside the project directory, initialize the module:

   ```bash
   go mod init my-go-project
   ```

2. After initializing the module, you can run the Go project as usual with `go run`.

---

## 💡 6. Run Tests (Optional)

To run tests within your Go project (if any):

```bash
go test
```

This will run any test functions in your project.

---

## 🔧 7. Common Commands

| Command               | Description                           |
|-----------------------|---------------------------------------|
| `go run <file>.go`     | Run a Go file directly.              |
| `go build`             | Compile the Go code into an executable.|
| `go test`              | Run tests for the project.           |
| `go mod init`          | Initialize a Go module.              |
| `go get <package>`     | Get a package dependency.            |
| `go fmt`               | Format the code according to Go's standards. |

---

## ✅ Summary

1. Install Go and set up your environment.
2. Write a Go program (e.g., `main.go`).
3. Run the program using `go run <file>.go`.
4. Optionally, build an executable using `go build`.
5. Use Go modules for managing dependencies with `go mod init`.

---

