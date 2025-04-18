# 📘 Variables and Data Types in Go

---

## ✅ Variables in Go

### 🔹 Declaration

```go
var name string
var age int
```

- `var` is used to declare a variable.
- If no value is assigned, Go gives it a **zero value**.

### 🔹 Declaration with Initialization

```go
var name string = "John"
var age int = 30
```

### 🔹 Type Inference (Short Declaration)

```go
name := "Alice"
age := 25
```

- The `:=` operator is a **short form declaration**.
- Only allowed **inside functions**.

### 🔹 Multiple Declarations

```go
var x, y int = 1, 2
```

---

### ⚙️ What Happens Internally?

When you declare a variable:
- The Go compiler reserves **stack memory**.
- If it **escapes** (used outside scope), it's placed on the **heap**.
- Compiler maps names to **memory addresses**.

Example:

```go
x := 5
```

Becomes in simplified assembly:

```asm
MOV [SP-4], 5  ; store 5 in stack location
```

---

## 📦 Data Types in Go

### 📌 1. **Basic Types**

| Type     | Description                  | Zero Value |
|----------|------------------------------|------------|
| `int`    | Integer (platform dependent) | `0`        |
| `float64`| 64-bit float                 | `0.0`      |
| `string` | UTF-8 text                   | `""`       |
| `bool`   | Boolean                      | `false`    |
| `byte`   | Alias for `uint8`            | `0`        |
| `rune`   | Alias for `int32` (Unicode)  | `0`        |

---

### 📌 2. **Composite Types**

- **Array**: Fixed length list
  ```go
  var a [3]int = [3]int{1, 2, 3}
  ```

- **Slice**: Dynamically-sized list
  ```go
  s := []int{1, 2, 3}
  ```

- **Map**: Key-value store
  ```go
  m := map[string]int{"a": 1, "b": 2}
  ```

- **Struct**: Custom types
  ```go
  type Person struct {
      Name string
      Age  int
  }
  ```

---

### 📌 3. **Pointer Types**

```go
var x int = 5
var p *int = &x
```

- `*int` is a pointer to an integer.
- `&x` gets the memory address of `x`.

---

## 🔁 Type Conversion

```go
var x int = 10
var y float64 = float64(x)
```

- Go does **not allow implicit conversion**
- Must cast explicitly between types

---

## 🔍 Behind the Scenes

- Go uses **static typing**.
- All variable types are known **at compile time**.
- Stack is used for short-lived values.
- Heap is used when values **escape** their local scope.
- Types define **how data is encoded in memory**.

---

## 🆚 Constants vs Variables

```go
const Pi = 3.14159
```

| Feature   | Variable     | Constant             |
|-----------|--------------|----------------------|
| Value     | Can change   | Cannot change        |
| Memory    | Allocated    | May not use memory   |
| Time      | Runtime      | Compile time         |

---

## ✅ Summary Table

| Concept      | Description                             |
|--------------|-----------------------------------------|
| `var`        | Declare variables                       |
| `:=`         | Short variable declaration              |
| Basic Types  | int, float, string, bool, etc.          |
| Composite    | array, slice, map, struct               |
| Pointers     | Reference memory location               |
| Conversion   | Explicit type casting                   |
| Memory       | Stack for short-living, heap for escape |

---

