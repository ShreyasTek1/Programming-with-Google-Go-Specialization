### 🔑 What is a Pointer?

A **pointer** is a variable that **stores the memory address** of another variable.

In Go, you use:

- `&` (ampersand) to get the address of a variable
    
- `*`(star) to **dereference** (i.e., access the value at) a pointer
    

---

### 🧠 Operators:

|Operator|Meaning|Example|
|---|---|---|
|`&`|"Address of" operator|`ptr := &x`|
|`*`|"Dereferencing" (value at address)|`val := *ptr`|

---

### 🧪 Example 1: Basic Pointers

```go
package main

import "fmt"

func main() {
    var x int = 1
    var y int
    var ip *int  // ip is a pointer to int

    ip = &x      // ip holds the address of x
    y = *ip      // y is assigned the value at the address ip points to

    fmt.Println("x:", x)     // Output: x: 1
    fmt.Println("ip:", ip)   // Output: ip: 0xc000012078 (example address)
    fmt.Println("y:", y)     // Output: y: 1
}
```

📌 **Explanation**:

- `x` has value `1`
    
- `ip = &x` means ip now holds the memory address of x
    
- `*ip` dereferences the pointer, giving you the value at that address → `1`
    
- `y = *ip` means `y` becomes `1`
    

---

### 🆕 Example 2: Using `new` Function

```go
package main

import "fmt"

func main() {
    ptr := new(int)     // ptr is a pointer to a newly allocated int (default value 0)
    fmt.Println(*ptr)   // Output: 0

    *ptr = 3            // set the value at that address to 3
    fmt.Println(*ptr)   // Output: 3
}
```

📌 **Explanation**:

- `new(int)` allocates space in memory for an integer, initialized to `0`, and returns its address
    
- `ptr` is a pointer to this new memory location
    
- `*ptr = 3` stores `3` in that memory
    

---

### 🧩 Summary:

|Concept|Explanation|
|---|---|
|`*int`|Type: Pointer to an integer|
|`&x`|Gives you the address of `x`|
|`*ptr`|Dereferences the pointer, gives the value at that address|
|`new(int)`|Allocates memory for an `int`, returns pointer to it|

---
## 🧠 What Is Variable Scope?

> **Scope** defines _where in the code a variable can be accessed_ or is _visible_.  
> In Go, this is determined using **blocks** `{}` and follows **lexical (static) scoping**.

---

## 🔄 Lexical Scoping in Go

Go is **lexically scoped**, meaning:

- The **location of a variable declaration in the source code** determines where it can be accessed.
    
- A variable is visible **in the block it’s declared in and any nested blocks**.
    

---

## 📦 Blocks in Go

A **block** is a section of code enclosed in `{}` (curly braces). There are also **implicit blocks** in Go:

|Block Type|Description|
|---|---|
|**Universe block**|Contains all Go source (predeclared identifiers)|
|**Package block**|All source files in a package|
|**File block**|Source code within a single file|
|**Function block**|Code inside a function|
|**If/For/Switch**|Introduce their own blocks|

Blocks are **nested**, and Go resolves variables from **inner → outer**.

---

## ✅ Example 1: Shared Variable in Outer Block

```go
package main

import "fmt"

var x = 1 // defined in the file block

func f() {
    fmt.Println("f:", x) // accesses x from file block
}

func g() {
    fmt.Println("g:", x) // accesses x from file block
}

func main() {
    f()
    g()
}
```

### 🔍 Scope Analysis:

- `x` is defined in **file block** `b1`
    
- Functions `f` and `g` each have their own blocks (`b2` and `b3`), both inside `b1`
    
- Both `f` and `g` can access `x` by looking up from their local blocks to the file block → ✅ No error
    

---

## ❌ Example 2: Variable Only in Inner Block

```go
package main

import "fmt"

func f() {
    x := 1
    fmt.Println("f:", x)
}

func g() {
    // fmt.Println("g:", x) // ❌ Error: x is undefined
}

func main() {
    f()
    g()
}
```

### 🔍 Scope Analysis:

- `x` is defined in `f`’s block only (`b2`)
    
- `g`’s block (`b3`) does **not** see `x` → ❌ Error
    
- There is no outer block where `x` is defined that includes both functions
    

---

## 🔄 Scope Rule (Summary):

> A variable is **accessible** in a block `bj` **if it’s declared in `bj` or any outer block `bi` where `bi ≥ bj`**.

🧠 Think of scope as **nested Russian dolls** — each block can see what's in the dolls above it, but not below or beside.

---

## 🧱 Visualizing the Hierarchy

```
File Block (b1)
├── Function f (b2)
│   └── Local variables (e.g., x)
└── Function g (b3)
    └── Cannot access x from b2
```

---

## ✨ Summary Points:

- Go uses **block-based lexical scoping**
    
- **Inner blocks** can access variables from **outer blocks**, not vice versa
    
- Blocks are defined by `{}` and language structures like `if`, `for`, `switch`
    
- **Variables should be declared in the correct block** to ensure accessibility where needed
    

---
### 🔑 **Key Concepts Explained Simply**

#### 1. **Memory Allocation and Deallocation**

- When you create a variable, it needs **space in memory**.
    
- This memory can’t be used by anything else until it's **deallocated** (freed).
    
- If memory is never freed, it's a **memory leak** – you eventually **run out of usable memory**.
    

---

### 🧠 **Where is Memory Stored?**

There are two main regions in memory:

#### ✅ **Stack**

- Used for **local variables** in a function.
    
- Memory is **automatically allocated** when a function is called.
    
- Memory is **automatically deallocated** when the function returns.
    
- Very fast and managed for you.
    
- **Example**: `int x = 5;` inside a function → x lives on the stack.
    

#### ✅ **Heap**

- Used for data that needs to **persist beyond a function call**.
    
- You must **manually allocate and deallocate** heap memory (e.g., using `malloc` and `free` in C).
    
- More **flexible** but **error-prone**.
    

---

### 🧑‍💻 **In C: Manual Heap Management**

```c
int* x = malloc(sizeof(int) * 10); // Allocates space for 10 integers
...
free(x); // Deallocates that memory
```

- If you forget `free(x);` → **Memory Leak**
    
- If you `free(x);` too early → **Dangling Pointer / Crash**
    

---

### 🧪 **In Go: Automatic Memory Management with Garbage Collection**

- Go **still has stack and heap**, but:
    
    - It **automatically chooses** where a variable lives (compiler decides).
        
    - Go uses **Garbage Collection (GC)** to deallocate heap memory.
        
- No need for `malloc` or `free` explicitly.
    
- **Less error-prone**, but **slightly slower** due to GC overhead.
    

#### Example in Go:

```go
func f() {
    x := 42  // may live on stack or heap depending on compiler's escape analysis
}
```

- If `x` is only used within `f()`, it goes on the **stack**.
    
- If `x` escapes `f()` (e.g., returned from the function), it goes on the **heap**.
    
- Go deallocates heap memory **automatically** when it's no longer reachable.
    

---

### 🧠 **Important Terms to Remember**

|Term|Meaning|
|---|---|
|**Stack**|Fast, automatic memory for function-local variables|
|**Heap**|Persistent memory requiring manual (C) or automatic (Go) cleanup|
|**Memory Leak**|Memory allocated but never freed|
|**Garbage Collection**|Automatic memory cleanup (used in Go, Java, Python, etc.)|
|**Escape Analysis**|Go compiler optimization to decide stack vs heap placement|

---

### 📌 TL;DR Summary

- Memory must be **allocated** when variables are created and **deallocated** when no longer needed.
    
- **Stack** memory is fast and managed automatically when functions are entered/exited.
    
- **Heap** memory is flexible but needs manual `malloc`/`free` in C, or is GC-managed in Go.
    
- **Go uses Garbage Collection** so you don't free memory manually, but understanding stack vs heap is still important.
    

---
### 🔍 **What's the problem with manual deallocation?**

In languages like **C/C++**, you **must deallocate memory manually** (e.g., using `free()` or `delete`). But this is **tricky and error-prone**, because:

- If you **deallocate too early**, you get a **dangling pointer** (using memory that's already freed).
    
- If you **forget to deallocate**, you get a **memory leak** (wasted memory not being used but never reclaimed).
    

This is especially problematic when:

- You return a **pointer to a local variable** in a function.
    
- You don’t clearly know who still needs a variable and who doesn’t.
    

---

### ✅ **What Go does differently**

Go avoids this headache using **garbage collection (GC)**, even though it's a **compiled language**.

#### 🔹 Example scenario:

```go
func foo() *int {
    x := 42
    return &x
}

func main() {
    p := foo()
    fmt.Println(*p)
}
```

- `x` is local to `foo()`. Normally, it would be destroyed after `foo()` ends.
    
- But since you **return a pointer to `x`**, Go’s **compiler promotes `x` to the heap** (instead of stack).
    
- The **garbage collector (GC)** now keeps track of `x`, because `main()` is still using it.
    

#### 🧠 So, Go is **smart enough** to:

1. Detect that `x` needs to live longer than the function.
    
2. Allocate `x` on the **heap**.
    
3. Let the **GC clean it up** **only** when no one uses it anymore.
    

---

### 🧹 **Garbage Collection (GC) in Go**

#### What GC does:

- **Tracks references (pointers)** to variables.
    
- When **no references remain**, it **automatically deallocates** the memory.
    
- You don’t have to decide: "stack or heap?", or "when to free memory?".
    

#### Go vs others:

|Feature|C/C++|Java/Python|**Go**|
|---|---|---|---|
|Manual memory management|✅ Yes|❌ No|❌ No|
|Garbage collection|❌ No|✅ Yes|✅ Yes|
|Compiled|✅ Yes|❌ No|✅ Yes|
|Needs interpreter|❌ No|✅ Yes|❌ No (compiled)|

So Go gives you **compiled performance** with **safe memory management**.

---

### ⚖️ **Trade-offs**

- **Pros**:
    
    - No need to manually free memory.
        
    - Avoids bugs like dangling pointers, memory leaks.
        
    - Easier and safer to use pointers.
        
- **Cons**:
    
    - **Garbage collection has a cost** (slightly reduced performance due to runtime memory checks).
        
    - You don’t control memory layout as precisely as in C/C++.
        

But Go's **GC is optimized and fast**, and for most use-cases, the safety and convenience **outweigh** the small performance trade-off.

---

### ✅ Summary

- Go’s **garbage collector** allows pointers to local variables to safely escape their scope (like returning a pointer to `x` in `foo()`).
    
- Go handles **stack vs heap allocation** **automatically**, based on variable lifetime.
    
- This makes Go a **safe**, **efficient**, and **modern** language for system-level programming without the manual memory pains of C/C++.
    

---
## 🟨 1. Comments in Go

### ✅ **Single-line Comments**

- Begin with `//`
    
- Everything to the right of `//` is ignored by the compiler.
    

```go
var x int // this is a comment
```

### ✅ **Multi-line (Block) Comments**

- Begin with `/*` and end with `*/`
    
- Can span multiple lines.
    

```go
/* This is a
multi-line comment */
```

> 💡 Use comments to clarify your code — they help readers understand your logic.

---

## 🟨 2. Print Statements in Go

### ✅ **Using the `fmt` Package**

To print anything, you must first import the `fmt` package:

```go
import "fmt"
```

### ✅ **Basic Print Example**

```go
fmt.Printf("Hi\n") // Prints: Hi
```

### ✅ **String Concatenation**

```go
x := "Joe"
fmt.Printf("Hi " + x + "\n") // Prints: Hi Joe
```

### ✅ **Format Strings**

Go uses format specifiers inside `Printf` to print variables cleanly.

#### Example:

```go
x := "Joe"
fmt.Printf("Hi %s\n", x) // %s is replaced with x → "Hi Joe"
```

#### Common Format Specifiers:

|Specifier|Meaning|
|---|---|
|`%s`|String|
|`%d`|Integer (decimal)|
|`%f`|Float|
|`%v`|Default format|

---

## 🟨 3. Integers in Go

### ✅ **Generic Declaration**

```go
var x int // Compiler decides bit size (32 or 64 depending on system)
```

### ✅ **Fixed-width Integers**

|Type|Bits|Range (signed)|Range (unsigned)|
|---|---|---|---|
|`int8`|8|-128 to 127|`uint8`: 0 to 255|
|`int16`|16|-32,768 to 32,767|`uint16`: 0 to 65,535|
|`int32`|32|~±2 billion|`uint32`: 0 to ~4B|
|`int64`|64|Huge range|`uint64`: 0 to very large|

- **Unsigned integers (`uint`)** don’t use a sign bit — all bits represent magnitude.
    

> Use fixed-size integers when you want **exact control** over memory or range.

---

## 🟨 4. Integer Operators in Go

### ✅ **Arithmetic Operators**

```go
+  // addition
-  // subtraction
*  // multiplication
/  // division
%  // modulus
```

### ✅ **Comparison Operators**

```go
==  // equal to
!=  // not equal to
>   // greater than
<   // less than
>=  // greater than or equal
<=  // less than or equal
```

### ✅ **Boolean Operators**

```go
&&  // logical AND
||  // logical OR
!   // logical NOT
```

### ✅ **Bitwise Operators**

```go
&   // bitwise AND
|   // bitwise OR
^   // bitwise XOR
&^  // bit clear (AND NOT)
<<  // left shift
>>  // right shift
```

---

## 🟨 Summary Cheat Sheet

|Topic|Syntax Example|Notes|
|---|---|---|
|Comment|`// comment` or `/* block */`|Ignored by compiler|
|Print|`fmt.Printf("Hi %s", name)`|Needs `import "fmt"`|
|Integer|`var x int`, `var y uint16`|`int` is system-dependent|
|Operators|`+`, `-`, `==`, `&&`, `<<`, etc.|Common across most programming languages|

---
## 📝 1. Comments in Go

- **Single-line Comment**:
    
    ```go
    // This is a comment
    var x int // Variable declaration with inline comment
    ```
    
- **Multi-line (Block) Comment**:
    
    ```go
    /*
       This is a multi-line comment.
       It can span multiple lines.
    */
    ```
    
- ✅ Comments are ignored by the compiler and are for humans to understand the code better.
    

---

## 📤 2. Print Statements (Using the `fmt` Package)

To use print functions:

```go
import "fmt"
```

### Printing Text:

```go
fmt.Println("Hello, World!")
```

### Concatenating Strings:

```go
x := "Joe"
fmt.Println("Hi " + x) // Prints: Hi Joe
```

### Using Format Strings (`Printf` with format specifiers):

```go
fmt.Printf("Hi %s\n", x) // %s for string
```

Common format specifiers:

- `%s` → string
    
- `%d` → decimal integer
    
- `%f` → floating point
    
- `%v` → default format (any value)
    

---

## 🔁 3. Type Conversion

### ✅ Valid Conversion (example: `int16` to `int32`):

```go
var y int16 = 2
var x int32
x = int32(y) // Explicit type conversion
```

- You **must** convert types explicitly if they differ (even if both are integers).
    

### ❌ Invalid Conversion (without casting):

```go
x = y // Error! Different types: int32 ≠ int16
```

> Not all type conversions are allowed. Some will cause compiler errors if not logically valid.

---

## 🔢 4. Integer Types

### Signed Integers:

- `int8` → -128 to 127
    
- `int16` → -32,768 to 32,767
    
- `int32` → -2B to 2B
    
- `int64` → Very large
    

### Unsigned Integers:

- `uint8` (same as `byte`) → 0 to 255
    
- `uint16`, `uint32`, `uint64` → progressively larger ranges
    

Use `int` (generic integer) unless you have a specific size requirement.

---

## 🧮 5. Operators in Go

### Arithmetic:

- `+`, `-`, `*`, `/`, `%` (modulo)
    

### Comparison:

- `==`, `!=`, `<`, `>`, `<=`, `>=`
    

### Logical:

- `&&` (AND), `||` (OR), `!` (NOT)
    

### Bitwise:

- `&`, `|`, `^`, `&^`, `<<`, `>>`
    

---

## 🌊 6. Floating Point Numbers

- `float32` → ~6 decimal digits precision
    
- `float64` → ~15 decimal digits precision (use this by default)
    

Example:

```go
var f float64 = 123.456
f = 1.23e2 // Scientific notation = 1.23 × 10² = 123
```

---

## 🔁 7. Complex Numbers

Go supports complex numbers using `complex` function:

```go
z := complex(2, 3) // 2 + 3i
```

---

## 🧵 8. Strings

- A string is an immutable sequence of **UTF-8** encoded bytes.
    
- Declared like this:
    
    ```go
    s := "Hello, world"
    ```
    

### String Facts:

- Immutable: You cannot change individual characters.
    
- UTF-8 encoded by default.
    
- Characters are called **runes** (Go term for Unicode code points).
    
- A **rune** is an alias for `int32`.
    

### String as Bytes:

```go
s := "A"
fmt.Println(s[0])       // Output: 65 (ASCII of 'A')
fmt.Printf("%c\n", s[0]) // Output: A
```

---

## 🔡 9. ASCII vs Unicode vs UTF-8

|Format|Size|Characters|Notes|
|---|---|---|---|
|ASCII|7-8 bits|128|English letters and control characters|
|Unicode|32 bits|2+ billion|Supports all global characters|
|UTF-8|8–32 bits|Variable|Default in Go, ASCII-compatible, efficient|

---

## ✅ Summary

|Topic|Key Takeaway|
|---|---|
|Comments|Use `//` or `/* */` for readability|
|Print Statements|Use `fmt` package with `Println`, `Printf`, etc.|
|Type Conversion|Must be explicit (e.g., `int32(y)`)|
|Integers|Use `int` unless you need a specific bit size|
|Floats|Prefer `float64` for more precision|
|Strings|Immutable UTF-8 sequences; made of runes|
|Unicode/UTF-8|UTF-8 is default encoding; supports global characters|

---
## 🧠 **Go Programming: Strings, Unicode, and Useful Packages**

---

### 🔤 **Strings in Go**

- Strings are **immutable sequences of bytes**.
    
- Represented in **UTF-8 encoding**, which supports Unicode characters.
    
- Each character in a string is a **rune** (alias for `int32`) that corresponds to a **Unicode code point**.
    
- You **cannot modify** a string directly, but you can create a **new modified version**.
    

---

### 🌐 **Unicode and Runes**

- A **rune** in Go = a Unicode code point (e.g., `'A'` is `0x41`).
    
- UTF-8 is a **variable-length encoding**:
    
    - ASCII characters: 1 byte
        
    - Other Unicode characters: 2–4 bytes
        

---

### 📦 **`unicode` Package**

Used for evaluating **individual characters** (runes):

#### 🔍 Predicate Functions (return `bool`)

- `unicode.IsDigit(r)` → is `r` a digit?
    
- `unicode.IsSpace(r)` → is `r` a whitespace character?
    
- `unicode.IsLetter(r)` → is `r` a letter?
    
- `unicode.IsPunct(r)` → is `r` punctuation?
    

#### 🔁 Conversion Functions

- `unicode.ToUpper(r)` → converts rune to uppercase
    
- `unicode.ToLower(r)` → converts rune to lowercase
    

---

### 📦 **`strings` Package**

Used for **whole string** operations.

#### 🔎 Search & Compare

- `strings.Compare(a, b)`  
    → returns `0` if equal, `-1` if `a < b`, `1` if `a > b`
    
- `strings.Contains(s, substr)`  
    → `true` if `substr` is in `s`
    
- `strings.HasPrefix(s, prefix)`  
    → `true` if `s` starts with `prefix`
    
- `strings.Index(s, substr)`  
    → returns **index** of first occurrence of `substr` in `s` (or `-1`)
    

#### 🔧 Manipulation (returns new strings)

- `strings.Replace(s, old, new, n)`  
    → replaces `n` instances of `old` with `new` in `s`
    
- `strings.ToLower(s)` / `strings.ToUpper(s)`  
    → converts string case
    
- `strings.TrimSpace(s)`  
    → removes leading/trailing whitespace
    

---

### 📦 **`strconv` Package**

Handles **conversion between strings and numbers**.

#### 🔁 Common Conversions

- `strconv.Atoi(s)`  
    → converts numeric **string** to `int` (e.g., `"123"` → `123`)
    
- `strconv.Itoa(i)`  
    → converts `int` to **string** (e.g., `123` → `"123"`)
    
- `strconv.ParseFloat(s, bitSize)`  
    → converts **string** to `float64` or `float32`
    
- `strconv.FormatFloat(f, fmt, prec, bitSize)`  
    → converts **float** to string (customizable format)
    

---

### 📌 Example Code Snippets

```go
import (
    "fmt"
    "unicode"
    "strings"
    "strconv"
)

func main() {
    r := 'a'
    fmt.Println(unicode.IsLetter(r))   // true
    fmt.Println(unicode.ToUpper(r))    // 'A'

    s := "Hello, World!"
    fmt.Println(strings.Contains(s, "World"))  // true
    fmt.Println(strings.ToUpper(s))            // "HELLO, WORLD!"

    numStr := "123"
    num, _ := strconv.Atoi(numStr)
    fmt.Println(num + 1)              // 124

    floatStr := "3.14"
    f, _ := strconv.ParseFloat(floatStr, 64)
    fmt.Println(f * 2)                // 6.28
}
```

---
### 🔸 **Constants in Go**

- A **constant** is a value known **at compile time**.
    
- Once declared, it **never changes** during the execution of the program.
    
- The **type is inferred** from the value on the right-hand side.
    
    ```go
    const x = 1.3  // x is a float64
    const y, z = 4, "hi"  // y is int, z is string
    ```
    

---

### 🔸 **What is `iota`?**

- `iota` is a special identifier in Go used to **generate a sequence of constants**.
    
- Useful when you want to define multiple constants with **unique values**, but the **actual value doesn’t matter**, only that they are **distinct**.
    
- Common use cases:
    
    - Days of the week
        
    - Months of the year
        
    - Grades or status codes
        
    - Enums in other languages like C
        

---

### 🔸 **Why Use `iota`?**

- You need several constants with **different values**.
    
- You don’t care **what the values actually are**.
    
- You just care that **each constant is unique**.
    

---

### 🔸 **How `iota` Works**

- Inside a `const` block, `iota` starts at `0` and increases by `1` for each new line.
    
- If `iota` is used on the first line, the following constants in the block will automatically increment.
    

---

### 🔸 **Example: Grades Using `iota`**

```go
type Grade int

const (
    A Grade = iota  // 0
    B               // 1
    C               // 2
    D               // 3
    F               // 4
)
```

- Even though we only write `iota` once, each subsequent constant automatically gets a new value.
    
- You **shouldn’t rely** on the actual numbers (e.g., 0, 1, 2...) — just that they're unique and different.
    

---

### ✅ Key Takeaways

- Use `const` for fixed, unchanging values.
    
- Use `iota` when defining a group of related constants where:
    
    - Values must be unique.
        
    - You don’t care about the actual numeric values.
        
    - You want cleaner, more maintainable code (like enums).
        

---
## 🔹 **What is Control Flow?**

Control flow is the **order** in which your code statements are **executed**.

- The default is **top-down**, one line after another.
    
- But control flow **changes** when you introduce:
    
    - Conditional statements (like `if`)
        
    - Loops (like `for`)
        
    - Multi-way branches (like `switch`)
        

---

## 🔸 `if` Statements in Go

Used to conditionally execute blocks of code.

### ✅ Basic Syntax:

```go
if condition {
    // Code to run if condition is true
}
```

- `condition` must be a boolean expression.
    
- Block is **only executed** if the condition is `true`.
    

### ✅ With `else` (optional):

```go
if x > 5 {
    fmt.Println("x is greater than 5")
} else {
    fmt.Println("x is 5 or less")
}
```

---

## 🔸 `for` Loops in Go

Go has **only one looping keyword: `for`**, but it has **3 forms**:

### ✅ 1. Traditional `for` (like C-style loop):

```go
for init; condition; update {
    // loop body
}
```

Example:

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

- **init** runs once at the start.
    
- **condition** is checked before every iteration.
    
- **update** runs after each iteration.
    
- If `condition` is always true → infinite loop.
    

---

### ✅ 2. Condition-only `for` (like a `while` loop):

```go
i := 0
for i < 10 {
    fmt.Println(i)
    i++
}
```

- Initialization happens **before** the loop.
    
- Update is done **inside** the loop body.
    

---

### ✅ 3. Infinite loop:

```go
for {
    // runs forever unless you break out
}
```

- Use `break` or `return` to exit.
    
- Rarely used except in things like servers or embedded systems.
    

---

## 🔸 `switch` Statements in Go

Used to select **one of many code paths**, based on a value.

### ✅ Basic Syntax:

```go
switch expression {
case value1:
    // do something
case value2:
    // do something else
default:
    // optional fallback
}
```

### Example:

```go
switch x {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
default:
    fmt.Println("Other")
}
```

### 🟢 Key Differences from C:

- **No need for `break`** — Go auto-breaks after each case.
    
- You can **omit `switch` expression** and use conditions instead:
    

```go
switch {
case x > 10:
    fmt.Println("x > 10")
case x < 5:
    fmt.Println("x < 5")
default:
    fmt.Println("x is between 5 and 10")
}
```

---

## ✅ Summary Table

|Control Flow|Description|Example|
|---|---|---|
|`if`|Conditional execution|`if x > 5 { ... }`|
|`for`|Repeated execution|`for i := 0; i < 10; i++ { ... }`|
|`switch`|Multi-way branching|`switch x { case 1: ... }`|

---
## 🔹 Tagless `switch` in Go

A **tagless switch** is a variant of the `switch` statement where no expression (or "tag") follows the `switch` keyword.

### ✅ How it works:

- Each `case` is a **boolean expression**.
    
- Go evaluates each `case` **in order**.
    
- It runs the **first `case` that is true**.
    
- If none are true, it runs the `default` (if present).
    

### 🧠 Think of it as a cleaner alternative to:

```go
if ... else if ... else ...
```

### ✅ Example:

```go
x := 2

switch {
case x > 1:
    fmt.Println("x is greater than 1")
case x < -1:
    fmt.Println("x is less than -1")
default:
    fmt.Println("x is between -1 and 1")
}
```

- Since `x > 1` is true, the first case runs and prints:  
    **`x is greater than 1`**
    

---

## 🔸 `break` and `continue` in Loops

### ✅ `break` — Exit the Loop Early

Used to **terminate** the loop immediately.

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break
    }
    fmt.Println(i)
}
```

**Output:**

```
0  
1  
2  
3  
4
```

- The loop stops when `i == 5`.
    

---

### ✅ `continue` — Skip Current Iteration

Used to **skip** the rest of the loop **for that iteration** only.

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        continue
    }
    fmt.Println(i)
}
```

**Output:**

```
0  
1  
2  
3  
4  
6  
7  
8  
9
```

- Skips printing when `i == 5`, but continues the loop.
    

---

## 🔹 Reading User Input with `fmt.Scan`

### ✅ Purpose:

Used to **get input from the user** (like reading from the keyboard).

### ✅ Syntax:

```go
var myVar int
fmt.Print("Enter a number: ")
fmt.Scan(&myVar)
fmt.Println("You entered:", myVar)
```

### 🔍 How it works:

- `Scan` reads **space-separated tokens** (words, numbers, etc.)
    
- You pass **pointers** to variables (`&myVar`)
    
- The function **blocks (waits)** for the user to press `Enter`
    
- It returns **two values**:
    
    1. Number of scanned items (e.g. `1` if you typed one value)
        
    2. An `error` (useful to detect invalid input)
        

### ✅ Example:

```go
var appleNum int
fmt.Print("Number of apples: ")
n, err := fmt.Scan(&appleNum)

if err != nil {
    fmt.Println("Input error:", err)
} else {
    fmt.Printf("You entered: %d (items scanned: %d)\n", appleNum, n)
}
```

---

## ✅ Summary Table

|Feature|Description|Example Code Snippet|
|---|---|---|
|**Tagless switch**|Cleaner alternative to if-else chains|`switch { case x > 0: ... }`|
|**`break`**|Exits the current loop immediately|`if i == 5 { break }`|
|**`continue`**|Skips to next iteration of the loop|`if i == 5 { continue }`|
|**`fmt.Scan`**|Reads user input into variables|`fmt.Scan(&x)`|

---

