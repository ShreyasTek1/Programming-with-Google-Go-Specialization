## ✅ **What Are Composite Data Types?**

**Composite (or aggregate) data types** **group multiple values together**—typically of the same or different types—into a single logical unit. They're essential for handling **complex data structures**, like a person's profile, a game state, or a database record.

---

## 🧩 Composite Data Types in Go

|Type|Description|Real-life Analogy|
|---|---|---|
|**Array**|Fixed-size collection of same type elements|A row of lockers|
|**Slice**|Dynamically-sized version of an array|A flexible list of items|
|**Map**|Key-value pairs (like a dictionary)|A phonebook|
|**Struct**|Custom type that groups fields of any type|A form with name, age, address, etc.|

---

## 🔹 1. **Arrays**

### ✅ Definition:

A fixed-length sequence of elements of the same type.

```go
var arr [3]int = [3]int{1, 2, 3}
fmt.Println(arr[0]) // Output: 1
```

- **Length is fixed** once declared.
    
- Can access via **indexing**, starting at `0`.
    

---

## 🔹 2. **Slices**

### ✅ Definition:

A **flexible**, **resizable**, and more commonly used alternative to arrays.

```go
nums := []int{10, 20, 30}
nums = append(nums, 40)
fmt.Println(nums) // [10 20 30 40]
```

- Built on top of arrays.
    
- Can grow or shrink dynamically.
    
- Use `append`, `copy`, slicing (`nums[1:3]`), etc.
    

---

## 🔹 3. **Maps**

### ✅ Definition:

A built-in associative data structure (like hash tables).

```go
ages := map[string]int{
    "Alice": 25,
    "Bob": 30,
}
fmt.Println(ages["Bob"]) // Output: 30
```

- Key-value pairs.
    
- Keys must be unique.
    
- Values can be updated, added, or deleted.
    

---

## 🔹 4. **Structs**

### ✅ Definition:

Custom composite data types that group together fields (of any type).

```go
type Person struct {
    Name string
    Age  int
}

p := Person{Name: "Shreyas", Age: 20}
fmt.Println(p.Name) // Output: Shreyas
```

- Extremely useful for modeling **real-world entities**.
    
- Supports methods (functions attached to struct types).
    

---

## 🎯 Why Are Composite Types Important?

- Real-world data is **not flat** — you need to model **hierarchies, groupings, relationships**.
    
- These types let you organize your code and data **logically and efficiently**.
    
- You’ll use them **everywhere**: APIs, files, databases, user interfaces, etc.
    

---

## ✅ After This Module

You’ll be able to:

- Store and manage collections (`arrays`, `slices`)
    
- Store key-value pairs (`maps`)
    
- Create custom blueprints for real-world models (`structs`)
    

---
## 🔷 **ARRAYS in Go**

### ✅ What is an Array?

An array is a **fixed-size** collection of elements **of the same type**. It’s **indexed**, and each element is initialized to the type’s **zero value**.

---

### 📌 **Key Properties of Arrays:**

|Feature|Description|
|---|---|
|**Fixed Size**|Must be defined at compile-time (e.g. `[5]int`)|
|**Same Type Elements**|All elements must be of the same type|
|**Indexing**|Accessed using indices: `arr[0]`, `arr[1]`, etc.|
|**Zero Value Init**|Automatically initialized (e.g. int = 0, string = "")|
|**Index Starts at 0**|Just like C, Python, JavaScript, etc.|

---

### 🧪 **Example: Array Declaration and Access**

```go
var x [5]int       // An array of 5 integers: initialized to [0, 0, 0, 0, 0]
x[0] = 2           // Set first element to 2
fmt.Println(x[1])  // Output: 0 (because it's auto-initialized)
```

---

## 🔸 **Array Literals**

You can **initialize an array inline** using curly braces `{}`:

```go
var x = [5]int{1, 2, 3, 4, 5}  // Literal initialization
```

---

### 🆗 **Using `...` (Ellipsis) to Infer Array Length**

Go lets you omit the length using `...` if you're using a literal:

```go
x := [...]int{10, 20, 30, 40}  // Go infers the array is length 4
```

> ✅ This is very convenient and prevents mismatch between array size and literal length.

---

## 🔁 **Iterating Over Arrays**

### ✅ With a `for` loop and `range`:

This is the **idiomatic way** to loop through arrays in Go.

```go
x := [...]int{1, 2, 3}

for i, v := range x {
    fmt.Println("Index:", i, "Value:", v)
}
```

- `i` is the index.
    
- `v` is the value at that index.
    

🔸 Each iteration assigns:

- `i = 0`, `v = 1`
    
- `i = 1`, `v = 2`
    
- `i = 2`, `v = 3`
    

---

## 🧵 Summary

|Concept|Example|Notes|
|---|---|---|
|Declare array|`var x [3]int`|Fixed size, auto-initialized to zero|
|Access element|`x[0] = 5`|Indexing starts at 0|
|Array literal|`x := [3]int{1, 2, 3}`|Init with values|
|Ellipsis|`x := [...]int{1, 2, 3}`|Let Go infer size|
|Iterate with loop|`for i, v := range x {}`|Get both index and value|

---
## 🧱 Composite Data Types in Go

### 🔹 What Are Composite Data Types?

- Composite types aggregate other types.
    
- Examples: **arrays**, **slices**, **maps**, and **structs**.
    
- Useful for modeling complex data in real-world applications.
    

---

## 📦 Arrays

### 🔸 Definition:

- **Fixed-length** sequence of elements of the same type.
    
- Length is known at **compile time**.
    

### 🔸 Syntax:

```go
var x [5]int // array of 5 ints, initialized to zero
x[0] = 2     // set first element
fmt.Println(x[1]) // prints 0
```

### 🔸 Array Literal:

```go
var x = [5]int{1, 2, 3, 4, 5}
```

### 🔸 Using `...`:

- Go can infer array length:
    

```go
x := [...]int{1, 2, 3, 4}
```

### 🔸 Iteration:

```go
arr := [3]int{1, 2, 3}
for i, v := range arr {
    fmt.Println("Index:", i, "Value:", v)
}
```

---

## 🍰 Slices

### 🔸 Definition:

- **Flexible**, **dynamic-sized** view over an **underlying array**.
    
- A "window" into part (or all) of an array.
    

### 🔸 Three Key Properties:

1. **Pointer** to the first element
    
2. **Length** – number of accessible elements
    
3. **Capacity** – number of elements from the pointer to the end of the array
    

### 🔸 Creation from an Array:

```go
arr := [7]string{"a", "b", "c", "d", "e", "f", "g"}
s1 := arr[1:3] // includes "b", "c"
s2 := arr[2:5] // includes "c", "d", "e"
```

- Slices can **overlap**.
    
- Modifying one can affect the other.
    

### 🔸 `len()` and `cap()`:

```go
fmt.Println(len(s1)) // 2
fmt.Println(cap(s1)) // capacity from s1[0] to end of array
```

### 🔸 Writing to Slices:

- Modifies the **underlying array**:
    

```go
s1[1] = "X"
fmt.Println(arr) // original array changes too
```

---

## ✏️ Slice Literals

### 🔸 Define a slice without creating an explicit array:

```go
sli := []string{"x", "y", "z"}
```

- This creates an **underlying array** automatically.
    
- `len(sli)` and `cap(sli)` are the same.
    

---
### 🧠 **Concept Recap: Slices in Go**

#### ✅ **1. What is a Slice?**

- A slice is a **flexible, growable** view into an array.
    
- It has:
    
    - `length` → how many elements are currently in it.
        
    - `capacity` → how many elements it can hold before it needs to grow.
        

---

### 🛠️ **`make` Function for Slices**

#### ✌️ Two Forms of `make`:

##### 1. `make([]T, length)`

- Creates a slice of type `T` with length `length`.
    
- Capacity = Length (underlying array is exactly `length` long).
    

```go
s := make([]int, 5)  // len = 5, cap = 5
```

##### 2. `make([]T, length, capacity)`

- Creates a slice of type `T` with:
    
    - Length = `length`
        
    - Capacity = `capacity`
        

```go
s := make([]int, 5, 10)  // len = 5, cap = 10
```

---

### ➕ **Appending to Slices**

#### 🧩 Use the `append()` Function:

- Syntax:
    
    ```go
    s = append(s, value)
    ```
    
- It adds `value` at the end of slice `s`.
    
- It **returns a new slice**, which might point to:
    
    - The same underlying array (if capacity allows),
        
    - OR a new array (if old capacity is full).
        

#### 🧪 Example:

```go
s := make([]int, 0, 3)   // len = 0, cap = 3
s = append(s, 100)       // s = [100], len = 1, cap = 3
```

If you go beyond capacity:

```go
s = append(s, 200, 300, 400) // Now cap is exceeded, Go allocates a new array
```

> ❗ Important: Always assign back the result of `append`, since it could return a new slice reference.

---

### 🧬 Under the Hood

- Slices are just a **descriptor**:
    
    - A pointer to an array,
        
    - A length,
        
    - A capacity.
        
- Arrays are not resizable, but slices simulate dynamic arrays by copying to larger arrays when needed.
    

---

### ✅ Summary

|Function|Purpose|
|---|---|
|`make([]T, l)`|Create slice with length = capacity = l|
|`make([]T, l, c)`|Create slice with len = l and cap = c|
|`append(s, x)`|Add `x` to slice `s`; returns new slice|

---
## 🧠 What is a Hash Table?

- A **hash table** is a data structure that maps **unique keys** to **values**.
    
- Think of it as a **dictionary**: you give it a key, and it gives you the associated value.
    
- It enables **very fast data access** — typically in **constant time** `O(1)`.
    

---

## 🔑 Key Concepts

### 1. **Key-Value Pairs**

- Each value is paired with a **unique key**.
    
    - Example: `"Joe" → x`, `"Jane" → y`, `"Pat" → z`
        
- Keys must be unique. Values don’t have to be.
    

### 2. **Hash Function**

- A function that takes a **key** and calculates a **slot (index)** in an array.
    
- You don't write or call the hash function — it's used **behind the scenes**.
    
- Example:
    
    ```go
    m := map[string]string{
        "Joe": "x",
        "Jane": "y",
        "Pat": "z",
    }
    ```
    

---

## 💡 Analogy

Imagine:

- A filing cabinet with numbered drawers.
    
- The key (e.g., "Joe") goes into a **hash function**, which tells you to put or look inside drawer `3`.
    
- You don’t care how it gets to `3`—you just use the key and it works.
    

---

## ✅ Advantages

|Feature|Benefit|
|---|---|
|🔍 Fast Lookup|Constant-time `O(1)` access on average.|
|🧠 Easy Key Names|You can use meaningful keys like `"username"` or `"email"` instead of integers.|
|🧺 Key Flexibility|Works with strings, numbers, or custom types (if hashable).|

---

## ⚠️ Disadvantage: **Collisions**

- A **collision** occurs when two different keys hash to the **same index**.
    
    - Example: Both "Joe" and "Jane" go to slot `2`.
        
- Go handles collisions automatically (with techniques like **chaining**).
    
- ⚡ Collisions can reduce speed but are usually rare with a good hash function.
    

---

## 🛠️ Hash Tables in Go → `map`

In Go, the built-in **`map`** type is your hash table.

### Declaring and Using a `map`:

```go
// Declare and initialize
emails := map[string]string{
    "Joe":  "joe@example.com",
    "Jane": "jane@example.com",
    "Pat":  "pat@example.com",
}

// Access a value
fmt.Println(emails["Jane"])  // Output: jane@example.com

// Add a new entry
emails["Alice"] = "alice@example.com"

// Delete an entry
delete(emails, "Pat")

// Check if key exists
email, exists := emails["Bob"]
if exists {
    fmt.Println("Bob's email:", email)
} else {
    fmt.Println("Bob not found")
}
```

---

## 🧠 Summary Table

|Term|Meaning|
|---|---|
|**Hash Table**|A data structure to map keys to values.|
|**Hash Function**|Maps a key to an index.|
|**Collision**|Two keys map to the same index.|
|**Go’s `map`**|Hash table implementation.|

---
## 🗺️ **Go Maps = Hash Tables**

Maps in Go store **key-value pairs**, with fast access via unique keys — just like hash tables.

---

## 🔨 **Creating Maps**

### 1. **Declare and make (empty map)**

```go
var idMap map[string]int      // Declares the map variable (but it's nil — not usable yet)
idMap = make(map[string]int)  // Initializes the map (ready to use)
```

### 2. **Declare and initialize with literals**

```go
idMap := map[string]int{
    "Joe": 123,
    "Jane": 456,
    "Pat": 789,
}
```

---

## 🎯 **Accessing Values**

### Get a value by key:

```go
fmt.Println(idMap["Joe"])  // Output: 123
```

### Add or update a value:

```go
idMap["Jane"] = 999  // Updates value for key "Jane"
idMap["Alice"] = 111 // Adds new key-value pair
```

### Delete a key-value pair:

```go
delete(idMap, "Joe")  // Removes "Joe" from the map
```

---

## ❓ **Check if a key exists**

Use a **two-value assignment** to check presence:

```go
val, ok := idMap["Jane"]
if ok {
    fmt.Println("Jane is in the map with value", val)
} else {
    fmt.Println("Jane not found")
}
```

- `ok` is `true` if the key exists, `false` otherwise.
    

---

## 📏 **Get map length**

```go
fmt.Println(len(idMap))  // Prints number of key-value pairs
```

---

## 🔁 **Iterate through a map**

Use a `for` loop with `range`:

```go
for key, val := range idMap {
    fmt.Println("Key:", key, "Value:", val)
}
```

You can also ignore a value using `_`:

```go
for key := range idMap {
    fmt.Println("Key only:", key)
}
```

---

## 🧠 Key Notes

|Feature|Behavior|
|---|---|
|Map keys must be **unique**|Re-assigning a key replaces the value|
|Keys can be of **comparable types**|e.g., `int`, `string`, `bool`, structs (not slices or maps)|
|Maps are **reference types**|Passed by reference (not copied)|
|Reading a non-existent key returns **zero value**|`0` for `int`, `""` for `string`, etc.|

---

## ✅ Example: Complete Go Map Program

```go
package main

import "fmt"

func main() {
    // Create and initialize a map
    idMap := map[string]int{
        "Joe":  123,
        "Jane": 456,
    }

    // Add/update values
    idMap["Alice"] = 789
    idMap["Jane"] = 999  // update

    // Access a value
    fmt.Println("Joe's ID:", idMap["Joe"])

    // Check if key exists
    val, ok := idMap["Bob"]
    if ok {
        fmt.Println("Bob's ID:", val)
    } else {
        fmt.Println("Bob not found")
    }

    // Delete a key
    delete(idMap, "Joe")

    // Length of map
    fmt.Println("Total entries:", len(idMap))

    // Iterate through map
    for name, id := range idMap {
        fmt.Printf("Name: %s, ID: %d\n", name, id)
    }
}
```

---
## 🔶 What is a Struct?

A **struct** (short for _structure_) is a **composite or aggregate data type** in Go (and other languages like C) that **groups together variables (called fields)** under one name. These fields can be of different types.

It helps you **organize related data together** in a single logical unit.

---

## 🔸 Why Use a Struct?

Imagine trying to manage a person's:

- Name
    
- Address
    
- Phone number
    

You **could** create:

```go
var name1, address1, phone1 string
var name2, address2, phone2 string
```

But this is messy. You must **manually keep track** of which name goes with which address and phone.

Instead, with a `struct`, you can group them together:

```go
type Person struct {
    Name    string
    Address string
    Phone   string
}
```

Now, all the data for one person is in **one object**.

---

## 🔸 Declaring a Struct

```go
type Person struct {
    Name    string
    Address string
    Phone   string
}
```

This defines a new **type** called `Person`, with three fields: `Name`, `Address`, and `Phone`.

---

## 🔸 Creating Struct Variables

You can now create individual people:

```go
var p1 Person
var p2 Person
```

Each one (`p1`, `p2`) is an independent person with their own Name, Address, and Phone.

---

## 🔸 Accessing Fields (Dot Notation)

Use `.` to **read** or **write** fields:

```go
p1.Name = "Alice"
p1.Address = "123 Main St"
p1.Phone = "9876543210"

fmt.Println(p1.Name) // Output: Alice
```

This is called **dot notation**.

---

## 🔸 Initializing Structs

### 1. **Zero Initialization**

```go
p1 := new(Person)
```

- This gives you a pointer to a new `Person` object
    
- All fields are set to **zero values** (`""` for strings)
    

### 2. **Struct Literal (Best Way)**

```go
p2 := Person{
    Name:    "Bob",
    Address: "456 Park Ave",
    Phone:   "1234567890",
}
```

This creates and fills in all the fields at once.

---

## 🔸 Summary

|Feature|Syntax / Example|
|---|---|
|Define Struct|`type Person struct { Name string ... }`|
|Create Struct Variable|`var p Person` or `p := Person{...}`|
|Access Fields|`p.Name = "Alice"`|
|Zero Initialization|`p := new(Person)`|
|Full Initialization|`p := Person{ Name: "Bob", ... }`|

---

## 🔸 Why Structs Are Important

Structs make your code:

- **More organized**
    
- **Easier to maintain**
    
- **Safer**, since related data is grouped together
    

---
