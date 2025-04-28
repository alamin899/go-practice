# Go Language Memory Simulation

This document explains how memory works internally when you run a Go program, based on the following example code:

```go
package main
import "fmt"

var globalVar int = 10 // Global variable

func add(a int, b int) int {
    sum := a + b       // Local variable
    return sum
}

func main() {
    result := add(5, 15)
    result2 := add(globalVar, 10)
    fmt.Println("Result2:", result2)
    fmt.Println("Result:", result)
}
```

# Memory Layout Overview

| Section | Description |
|:--------|:------------|
| **Text** | Compiled code for `main`, `add`, and `fmt.Println`. |
| **Data** | Global variables like `globalVar = 10`. |
| **BSS** | Global variables declared but not initialized (not used here). |
| **Heap** | Dynamically allocated memory (not needed here). |
| **Stack** | Function calls and local variables like `a`, `b`, `sum`, `result`, and `result2`. |

---

# Memory Actions Step-by-Step

### Step 1: Program Load
- **Text Section**: `main`, `add`, `fmt.Println` loaded.
- **Data Section**: `globalVar = 10` stored.
- Stack and Heap are initially empty.

### Step 2: main() starts
- Stack frame for `main` function created.

### Step 3: add(5, 15) call
- Stack grows.
- New stack frame for `add` created.
- Local variables `a=5`, `b=15` created.
- Local variable `sum=20` calculated.

### Step 4: add() returns
- Stack frame for `add` is popped.
- `result=20` stored in `main` frame.

### Step 5: add(globalVar, 10) call
- Stack grows again.
- New stack frame for `add` created.
- Local variables `a=10`, `b=10` created.
- Local variable `sum=20` calculated.

### Step 6: add() returns
- Stack frame for `add` popped again.
- `result2=20` stored in `main` frame.

### Step 7: fmt.Println()
- Stack temporarily holds parameters for printing.
- Prints "Result2: 20" and "Result: 20".

### Step 8: Program Exit
- Stack cleared.
- Process ends.

---

# Memory Flow Visual

![Go Memory Management Simulation GIF](sandbox:/mnt/data/A_GIF_exhibits_eight_distinct_frames,_illustrating.png)

This GIF illustrates the full memory simulation step-by-step:
- Stack frames created and destroyed.
- Global variables remain in the Data section.
- No Heap usage because no pointers or dynamic allocations were involved.

---

# Important Points

- **Stack** grows and shrinks as functions are called and return.
- **Global variables** persist for the program's entire life.
- **Escape Analysis** moves variables to Heap only if their address escapes (not needed here).
- **Garbage Collector** automatically cleans unused memory.

# Quick Summary

> Go manages memory automatically with a structured layout:
> - **Text** for code
> - **Data** for global variables
> - **Stack** for function execution
> - **Heap** for dynamic memory (if needed)
> Go's memory model is efficient, using fast stack allocations and garbage-collected heap memory.

---

# End of Simulation

This covers the full internal memory working for the given Go code!
