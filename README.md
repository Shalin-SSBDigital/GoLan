# 🐹 Go Learning Docs

I'm learning Go and making my own docs because the official ones are **so boring**.

This repo is my personal reference — every concept explained in plain language with runnable examples, detailed comments, and Python comparisons where it helps.

**Progress:** 21 topics covered across Level 1 (Core Go). See the full [[Roadmap]] in the Second Brain.

---

## 📂 Project Structure

```
GoLan/
├── README.md
├── Get Started/
│   └── main.go                     # Hello World + package/import basics
├── Variables/
│   ├── Declare Variables.go        # var vs :=, data types
│   ├── Go Multiple Variable Declaration.go  # Parallel declaration
│   └── Constants.go                # const, iota, immutability
├── Output/
│   └── output funtion.go           # fmt.Println, functions, basic func syntax
├── If-Else/
│   └── Else.go                     # if, else, else if, logical operators
├── For/
│   └── loop.go                     # All 4 for-loop forms + break/continue
├── Array /                         (note: trailing space in folder name)
│   ├── arr.go                      # 1D arrays, 2D arrays, sparse init
│   ├── Slices.go                   # make, append, copy, slice expressions
│   └── maps.go                     # CRUD, comma-ok, nil map behavior
├── Switch/
│   └── switch.go                   # Expression switch, fallthrough, type switch
├── Functions/
│   ├── function.go                 # Basic func syntax, named returns
│   ├── Multiple Return Values.go   # Tuple-like returns, error pattern
│   ├── Variadic Functions.go       # ...type, slice unpacking
│   └── Closures.go                 # Closures, loop-variable gotcha
├── Methods/
│   └── Methods.go                  # Value/pointer receivers, Stringer
├── Interface/                      (older duplicate — see Interfaces/)
│   └── Interfaces.go               # Basic interface examples
├── Interfaces/
│   └── interfaces.go               # Comprehensive: composition, any, type switch
├── Structs/
│   └── Structs.go                  # ⚠️ Stub — placeholder only
├── Struct-Embedding/
│   └── struct-embedding.go         # Composition over inheritance
├── Pointers/
│   └── pointers.go                 # & and * operators
├── strings-and-runes/
│   └── strings-and-runes.go        # UTF-8, runes, byte vs char
└── Generics/
    └── generics.go                 # Type params, constraints, generic stack
```

---

## 📘 Topics Covered

### 1. Getting Started — [`Get Started/main.go`](Get%20Started/main.go)

**Concepts:** package main, import, func main(), fmt.Println

```go
package main
import "fmt"
func main() {
    fmt.Println("Hello World!")
}
```

| Concept | Explanation |
|---------|-------------|
| `package main` | Tells Go to build an executable (not a library) |
| `import "fmt"` | Imports the format package for printing |
| `func main()` | Entry point — execution starts here |
| `fmt.Println()` | Prints text + newline to console |

---

### 2. Variables — [`Variables/Declare Variables.go`](Variables/Declare%20Variables.go)

**Concepts:** var keyword, := short declaration, type inference, basic data types

```go
var shalin int = 25    // explicit type
name := "Shalin"       // inferred as string
```

**Data types covered:**

| Type | Example | Description |
|------|---------|-------------|
| `int` | `42` | Whole numbers |
| `float64` | `3.14` | Decimal numbers |
| `string` | `"hello"` | Text (double quotes only) |
| `bool` | `true` | true or false |

**var vs := comparison:**

| | `var` | `:=` |
|---|---|---|
| Scope | package or function | function only |
| Type | explicit or inferred | always inferred |
| Zero values | supported | not supported |

---

### 3. Multiple Variables — [`Variables/Go Multiple Variable Declaration.go`](Variables/Go%20Multiple%20Variable%20Declaration.go)

**Concepts:** parallel declaration, same-type grouping

```go
var a, b, c, d int = 1, 3, 5, 7
```

Values are assigned positionally. All variables must share the same type.

---

### 4. Constants — [`Variables/Constants.go`](Variables/Constants.go)

**Concepts:** const keyword, immutability, compile-time evaluation, iota

```go
const Pi float64 = 3.14
// Pi = 2  // COMPILE ERROR — constants cannot change
```

**Key rules:**
- Value must be known at compile time
- Reassignment causes a compile error
- Only primitive types: numeric, string, bool

**Bonus patterns:**
- Parallel const blocks
- `iota` auto-incrementing generator

---

### 5. Output — [`Output/output funtion.go`](Output/output%20funtion.go)

**Concepts:** fmt.Println, multi-argument printing, user-defined functions

```go
a, b := 10, 20
fmt.Println("The sum is: ", a+b)
```

**Function syntax in Go:**
```go
func add(a int, b int) int {
    return a + b
}
```

Type comes **after** the variable name (unlike C/Java). If consecutive params share a type, you can write `a, b int`.

---

### 6. If-Else — [`If-Else/Else.go`](If-Else/Else.go)

**Concepts:** if, else, else if, `||` operator, short-statement if

```go
if 7%2 == 0 {
    fmt.Println("even")
} else {
    fmt.Println("odd")
}

if num := 9; num < 10 {
    fmt.Println(num, "has 1 digit")
}
```

**All forms:**

| Form | Use case |
|------|----------|
| `if c { }` | Single case |
| `if c { } else { }` | Two branches |
| `if c { } else if { } else { }` | Multiple branches |
| `if s; c { }` | Scoped variable + condition |

**Note:** Go has NO ternary operator (`? :`). Use if-else everywhere.

---

### 7. For Loops — [`For/loop.go`](For/loop.go)

**Concepts:** All 4 loop forms, break, continue, range

**Loop variants:**

| Form | Example | Use case |
|------|---------|----------|
| Three-part | `for i := 0; i < 5; i++` | Counted iteration |
| Condition-only | `for m < 5` | While-style loop |
| Infinite | `for { }` | Server loops / until break |
| Range | `for i, v := range slice` | Iterate collections |

```go
// Standard
for i := 0; i < 5; i++ { }

// While-style
m := 0
for m < 5 { m++ }

// Infinite + break
for { if done { break } }

// Range over slice
for idx, val := range []string{"a", "b"} { }
```

**Flow control:**
- `break` — exits the loop entirely
- `continue` — skips to next iteration

---

### 8. Arrays — [`Array/arr.go`](Array%20/arr.go)

**Concepts:** 1D arrays, 2D arrays, literal init, sparse init, len(), zero values

```go
var a [5]int              // [0 0 0 0 0]
b := [5]int{1, 2, 3, 4, 5}
c := [...]int{100, 3: 400, 500}   // [100 0 0 400 500]
```

**Key properties:**

| Property | Behavior |
|----------|----------|
| Fixed size | Set at compile time — never changes |
| Value type | Assignment COPIES all elements |
| Zero values | All elements auto-initialized to 0 |
| `len()` | Returns compile-time length |
| Comparison | `arr1 == arr2` allowed (same type only) |

**2D arrays:**
```go
var twoD [2][3]int
twoD = [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
}
```

---

### 9. Slices ⭐ — [`Array/Slices.go`](Array%20/Slices.go)

**Concepts:** make, append, copy, slice expressions, variadic expansion

```go
s := make([]string, 3)           // ["", "", ""]
s[0] = "a"
s = append(s, "d", "e")          // ["a", "", "", "d", "e"]
c := make([]string, len(s))
copy(c, s)                       // deep copy
l := s[2:5]                      // slice expression: [2:5]
```

**Slice vs Array:**

| | Array | Slice |
|---|---|---|
| Size | Fixed at compile time | Dynamic (grows with append) |
| Type includes size? | Yes: `[3]int` | No: `[]int` |
| Passed by | Value (copies) | Reference (shares backing array) |
| Zero value | Fixed-length array | `nil` (length 0, usable) |

**Key patterns:**
- `make([]T, len, cap)` — create with initial length + capacity
- `append(s, x...)` — variadic expansion to append another slice
- `s[low:high]` — creates a view into the backing array (no copy)

---

### 10. Maps ⭐ — [`Array/maps.go`](Array%20/maps.go)

**Concepts:** make, CRUD operations, comma-ok idiom, nil map behavior

```go
m := make(map[string]int)
m["k1"] = 7
m["k2"] = 13
delete(m, "k2")
val, ok := m["k1"]     // 7, true
_, exists := m["nope"] // 0, false
```

**Compared to Python:**

| Operation | Go | Python |
|-----------|----|--------|
| Create | `m := make(map[K]V)` | `d = {}` |
| Get with check | `v, ok := m["k"]` | `d.get("k")` or try/except |
| Missing key | Returns zero value | Raises `KeyError` |
| Delete | `delete(m, "k")` | `del d["k"]` |
| Equal check | `maps.Equal(a, b)` (Go 1.21+) | `a == b` |
| Iteration order | Random (deliberate) | Insertion order (3.7+) |

**Gotcha:** Missing keys return the zero value (0, "", false). Always use comma-ok to distinguish "missing" from "stored as zero."

---

### 11. Switch — [`Switch/switch.go`](Switch/switch.go)

**Concepts:** expression switch, fallthrough, type switch basics

```go
switch i {
case 1:
    fmt.Println("one")
case 2, 3:
    fmt.Println("two or three")
default:
    fmt.Println("other")
}
```

**Go switch quirks:**
- **No `break` needed** — each case automatically breaks (no fallthrough by default)
- **`fallthrough`** — explicitly continue to next case (rarely used)
- **Multiple values** — `case 2, 3:` matches either
- **Type switch** — `switch v.(type)` to branch on the dynamic type of an interface

---

### 12. Functions — [`Functions/function.go`](Functions/function.go)

**Concepts:** func syntax, named returns, type-after-name convention

```go
func add(a int, b int) int {
    return a + b
}

func greet(name string) string {
    return "Hello " + name
}

func getCoords() (x, y int) {    // named returns
    x = 10                        // bare return
    y = 20
    return
}
```

**Compared to Python:**

| Feature | Go | Python |
|---------|----|--------|
| Return type | After params: `func() int` | Before params: `def f() -> int` |
| Named returns | First-class, bare return | Not supported |
| Multiple returns | True multi-return | Tuples (single value) |
| Default args | Not supported | Supported |

---

### 13. Multiple Return Values — [`Functions/Multiple Return Values.go`](Functions/Multiple%20Return%20Values.go)

**Concepts:** true multi-return, `_` discard, (result, error) pattern

```go
func vals() (int, int) {
    return 3, 7
}

a, b := vals()             // both values
_, c := vals()             // discard first with _
```

**The (result, error) pattern** is ubiquitous in Go:
```go
f, err := os.Open("file.txt")
if err != nil {
    // handle error
}
```

This is Go's idiomatic error handling — no try/catch, no exceptions. Every function that can fail returns `(result, error)`.

---

### 14. Variadic Functions — [`Functions/Variadic Functions.go`](Functions/Variadic%20Functions.go)

**Concepts:** `...type` syntax, slice unpacking

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

sum(1, 2)       // 3
sum(1, 2, 3, 4) // 10

nums := []int{1, 2, 3}
sum(nums...)    // unpack slice with ...
```

**Compared to Python:** Go's `...int` = Python's `*args`. The variadic parameter must be the last parameter.

---

### 15. Closures — [`Functions/Closures.go`](Functions/Closures.go)

**Concepts:** closure creation, mutation capture, loop-variable gotcha

```go
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

nextInt := intSeq()    // i is captured
fmt.Println(nextInt()) // 1
fmt.Println(nextInt()) // 2
```

**Loop-variable gotcha (same as Python):**
```go
for _, v := range values {
    v := v                // <-- FIX: create a new variable per iteration
    funcs = append(funcs, func() { fmt.Println(v) })
}
```

Without `v := v`, all closures capture the SAME loop variable.

---

### 16. Methods — [`Methods/Methods.go`](Methods/Methods.go)

**Concepts:** value vs pointer receivers, methods on any type, Stringer, Python comparison

```go
type Rectangle struct {
    Width, Height float64
}

// Value receiver — operates on a copy
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Pointer receiver — can mutate
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}
```

**Receiver types:**

| Receiver | Copies? | Can mutate? | When to use |
|----------|---------|-------------|-------------|
| Value `(r T)` | Yes | No | Small types, read-only |
| Pointer `(r *T)` | No | Yes | Large structs, mutation needed |

**Methods on ANY type:**
```go
type MyFloat float64
func (f MyFloat) Abs() float64 {
    if f < 0 { return float64(-f) }
    return float64(f)
}
```

---

### 17. Interfaces ⭐ — [`Interfaces/interfaces.go`](Interfaces/interfaces.go)

**Concepts:** implicit satisfaction, composition, empty interface (any), type assertion, type switch

```go
type Speaker interface {
    Speak() string
}

type Dog struct{ Name string }

// Dog implicitly satisfies Speaker — no "implements" keyword
func (d Dog) Speak() string {
    return d.Name + " says woof!"
}
```

**Why Go interfaces are special:**
- **Implicit satisfaction** — no `implements` keyword. If a type has the methods, it satisfies the interface
- **Duck typing** at compile time — "if it walks like a duck..."
- **Composable** — interfaces can embed other interfaces
- **`any`** (formerly `interface{}`) — empty interface, holds any type

**Type assertion:**
```go
var i any = "hello"
s := i.(string)          // panics if wrong type
s, ok := i.(string)      // safe: ok=false if wrong type
```

**Standard interfaces:**
- `fmt.Stringer` — `String() string` (like Python's `__str__`)
- `error` — `Error() string` (like Python's `__str__` on exceptions)

---

### 18. Struct Embedding — [`Struct-Embedding/struct-embedding.go`](Struct-Embedding/struct-embedding.go)

**Concepts:** composition over inheritance, field/method promotion, shadowing

```go
type Base struct {
    Num int
}

type Container struct {
    Base                    // embedded (no field name)
    Str string
}

c := Container{Base{10}, "hello"}
fmt.Println(c.Num)          // promoted from Base — accessed directly
```

**Rules of embedding (NOT inheritance):**

| Concept | Go Embedding | Python Inheritance |
|---------|-------------|-------------------|
| Relationship | HAS-A (composition) | IS-A (inheritance) |
| Method resolution | Promotion (flat) | MRO (chain) |
| Diamond problem | Compile error | Resolved via MRO |
| `super()` | Not supported | Supported |
| Shadowing | Outer overrides inner | `super()` can access parent |

Go embedding is **composition with syntactic sugar** — promoted methods become part of the outer type's API.

---

### 19. Pointers — [`Pointers/pointers.go`](Pointers/pointers.go)

**Concepts:** `&` (address-of), `*` (dereference), mutating through pointers

```go
i := 42
p := &i                 // p is a pointer to i
fmt.Println(*p)         // 42 (dereference)
*p = 21                 // mutate i through pointer
fmt.Println(i)          // 21
```

**Pointers enable mutation:**
```go
func zeroVal(val int) {
    val = 0             // only modifies local copy
}

func zeroPtr(ptr *int) {
    *ptr = 0            // modifies original
}
```

**In Go, you mostly use pointers with:**
- Large structs (avoid copying)
- Methods that need to mutate the receiver
- nilable fields (nil pointer = zero value)

---

### 20. Strings & Runes — [`strings-and-runes/strings-and-runes.go`](strings-and-runes/strings-and-runes.go)

**Concepts:** UTF-8 byte sequences, rune type, range loops, conversions

```go
s := "Hello, 世界"
fmt.Println(len(s))            // 13 bytes (not 9 characters)
fmt.Println(utf8.RuneCountInString(s)) // 9 characters

for i, r := range s {          // range decodes UTF-8 automatically
    fmt.Printf("%d → %c (%U)", i, r, r)
}
```

**Byte vs Rune:**

| | byte | rune |
|---|---|---|
| Alias for | `uint8` | `int32` |
| Represents | Raw byte | Unicode code point |
| ASCII? | One byte = one char | One rune = one char |
| Non-ASCII? | Multi-byte sequence | Still one value |

**Key gotchas:**
- `len(s)` counts **bytes**, not characters — use `utf8.RuneCountInString(s)` for char count
- `s[i]` gives a raw byte, not a character — use `[]rune(s)[i]` or range loop
- String slicing `s[:5]` creates a view (O(1)), but can slice into the middle of a multi-byte rune

---

### 21. Generics — [`Generics/generics.go`](Generics/generics.go)

**Concepts:** type parameters `[T any]`, constraints, generic structs, generic functions

```go
func Map[T any](items []T, fn func(T) T) []T {
    result := make([]T, len(items))
    for i, item := range items {
        result[i] = fn(item)
    }
    return result
}

type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}
```

**Available since Go 1.18.** Uses square brackets `[T any]` instead of angle brackets `[T]` (which would conflict with the I/O operators `<` and `>`).

**Built-in constraints:**
- `any` — any type (replaces `interface{}`)
- `comparable` — types that support `==` and `!=`
- `constraints.Ordered` — types with `<`, `>`, `<=`, `>=` (golang.org/x/exp)

---

## 🚀 How to Run

```bash
# From the GoLan directory:

# Executable files (package main)
go run "Get Started/main.go"
go run For/loop.go
go run "If-Else/Else.go"
go run "Array /arr.go"
go run "Array /Slices.go"
go run "Array /maps.go"
go run "Output/output funtion.go"
go run "Variables/Go Multiple Variable Declaration.go"
go run Switch/switch.go
go run "Functions/function.go"
go run "Functions/Multiple Return Values.go"
go run "Functions/Variadic Functions.go"
go run "Functions/Closures.go"
go run Methods/Methods.go
go run Interfaces/interfaces.go
go run "Struct-Embedding/struct-embedding.go"
go run Pointers/pointers.go
go run strings-and-runes/strings-and-runes.go
go run Generics/generics.go

# Reference-only files (package not main — read for learning)
# Variables/Declare Variables.go   — package variables
# Variables/Constants.go          — package variables
# Interface/Interfaces.go         — older duplicate, see Interfaces/
# Structs/Structs.go              — stub (placeholder)
```

> **Note:** Some files in `Variables/` and `Interface/` use non-main packages for naming or organizational reasons. They're meant for reading/reference, not standalone execution. Most files in this repo are `package main` and fully runnable.

---

## 📈 Progress Overview

Check out the full [[Roadmap]] in the Second Brain vault for a complete topic tracker across 105 items. Current status:

| Level | Coverage |
|---|---|
| Level 1 — Core Go (37 topics) | 23 ✅ (62%) |
| Level 2 — Production Go (27 topics) | ❌ Not started |
| Level 3 — Industry Backend (41 topics) | ❌ Not started |
| **Total** | **22%** |

**Up next:** Implement full Structs examples, then error handling → defer/panic → packages/modules.

---

## 📝 Why This Exists

The official Go docs are technically correct but **painfully dry**. This repo is my attempt to document what I learn in a way that's:

- **Readable** — plain language, no jargon for the sake of it
- **Runnable** — every example is a real .go file you can execute
- **Annotated** — detailed inline comments explaining every line
- **Comparable** — Python side-by-side where the conceptual difference matters

If you're also learning Go, feel free to use this as a reference. Contributions and suggestions are welcome!

---

## 🔗 Related

- [[Roadmap]] — Full Go industry roadmap in the Second Brain
- [[Notes]] — Gotchas, gotchas, and more gotchas
- [[Tasks]] — What's next
