# 🐹 Go Learning Docs

I'm learning Go and making my own docs because the official ones are **so boring**.

This repo is my personal reference — every concept explained in plain language with runnable examples, detailed comments, and Python comparisons where it helps.

**Progress:** Level 1 complete ✅ (36/37). See the full [[Roadmap]] in the Second Brain.

> ⚠️ **Notice:** The `Interface/` folder is a duplicate of `Interfaces/` and will be cleaned up. This is the 1 remaining Level 1 item.

---

## 📂 Project Structure

```
GoLan/
├── README.md
├── CLAUDE.md
├── Get Started/
│   ├── main.go                     # Hello World + package/import basics
│   └── README.md
├── Variables/
│   ├── Declare Variables.go        # var vs :=, data types
│   ├── Go Multiple Variable Declaration.go  # Parallel declaration
│   ├── Constants.go                # const, iota, immutability
│   └── README.md
├── Type Conversion/
│   └── conversion.go               # int/float, strconv, string↔[]byte
├── Operators/
│   └── operators.go                # Arithmetic, comparison, bitwise, etc.
├── Output/
│   ├── output funtion.go           # fmt.Println, functions, basic func syntax
│   └── README.md
├── If-Else/
│   ├── Else.go                     # if, else, else if, logical operators
│   └── README.md
├── For/
│   ├── loop.go                     # All 4 for-loop forms + break/continue
│   └── README.md
├── Switch/
│   ├── switch.go                   # Expression switch, fallthrough, type switch
│   └── README.md
├── Array /                         (note: trailing space in folder name)
│   ├── arr.go                      # 1D arrays, 2D arrays, sparse init
│   ├── Slices.go                   # make, append, copy, slice expressions
│   └── maps.go                     # CRUD, comma-ok, nil map behavior
├── Functions/
│   ├── function.go                 # Basic func syntax, named returns
│   ├── Multiple Return Values.go   # Tuple-like returns, error pattern
│   ├── Variadic Functions.go       # ...type, slice unpacking
│   ├── Closures.go                 # Closures, loop-variable gotcha
│   └── README.md
├── Recursion/
│   └── recursion.go                # Factorial, fibonacci, tree traversal
├── Methods/
│   ├── Methods.go                  # Value/pointer receivers, Stringer
│   └── README.md
├── Interface/                      ⚠️ Duplicate — see Interfaces/
│   ├── Interfaces.go               # Basic interface examples
│   └── README.md
├── Interfaces/
│   ├── interfaces.go               # Comprehensive: composition, any, type switch
│   └── README.md
├── Structs/
│   ├── Structs.go                  # Creation, tags, embedding, constructor pattern
│   └── README.md
├── Struct-Embedding/
│   ├── struct-embedding.go         # Composition over inheritance
│   └── README.md
├── Pointers/
│   ├── pointers.go                 # & and * operators
│   └── README.md
├── Enums/
│   └── enums.go                    # iota, String(), bitmasks
├── strings-and-runes/
│   ├── strings-and-runes.go        # UTF-8, runes, byte vs char
│   └── README.md
├── Strings-Builder/
│   └── strings_builder.go          # Builder methods, Grow, vs +=
├── Error Handling/
│   └── errors.go                   # error interface, panic, recover, defer
├── Packages-Modules/
│   ├── go.mod                      # Module definition
│   ├── main.go                     # Package usage examples
│   ├── calculator/                 # Sub-package
│   ├── config/                     # Sub-package
│   ├── models/                     # Sub-package
│   └── utils/                      # Sub-package
├── Generics/
│   ├── generics.go                 # Type params, constraints, generic stack
│   └── README.md
├── GoRoutines/
│   ├── goroutines.go               # go keyword, WaitGroup, closure gotcha
│   ├── LEARN.md                    # Deep dive
│   └── README.md
├── Channels/
│   ├── channel.go                  # Unbuffered/buffered, close, range, Lego analogy
│   ├── LEARN.md                    # Deep dive
│   └── README.md
├── Concurrency/
│   ├── concurrency.go              # Fan-out/fan-in, select, mutex
│   ├── LEARN.md                    # Deep dive
│   └── README.md
├── Memory/
│   ├── memory.go                   # Stack vs heap, escape analysis examples
│   ├── escape-analysis.go          # Deep dive: -gcflags="-m", 10 escape rules
│   ├── LEARN.md
│   └── README.md
├── Go-GPM-Engine/
│   ├── LEARN.md                    # G-P-M scheduler deep dive
│   └── README.md
├── Go-Garbage-Collector/
│   ├── LEARN.md                    # Tri-color mark-sweep, write barrier, GOGC
│   └── README.md
├── Go-Pipelines/
│   ├── LEARN.md                    # Pipeline pattern, fan-out/fan-in
│   └── README.md
└── Go-Interface-Internals/
    ├── LEARN.md                    # iface, eface, itab, nil-is-not-nil
    └── README.md
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

### 21. Type Conversion — [`Type Conversion/conversion.go`](Type%20Conversion/conversion.go)

**Concepts:** int/float conversion, strconv (Itoa/Atoi), string↔[]byte, rune conversion, type assertion vs conversion

```go
var i int = 42
var f float64 = float64(i)          // explicit conversion (no implicit)

s := strconv.Itoa(42)               // "42"
n, err := strconv.Atoi("42")        // 42, nil (returns error on fail)

b := []byte("hello")                 // string → bytes
str := string([]byte{104, 101})     // bytes → string
```

| Operation | Go | Python |
|-----------|----|--------|
| int → float | `float64(i)` | `float(i)` |
| int → string | `strconv.Itoa(i)` | `str(i)` |
| string → int | `strconv.Atoi(s)` | `int(s)` |
| Type check | `v.(T)` or type switch | `isinstance(v, T)` |

**Key rule:** Go requires **explicit** conversion — no automatic int→float like Python/Python.

---

### 22. Operators — [`Operators/operators.go`](Operators/operators.go)

**Concepts:** arithmetic, comparison, logical, bitwise, assignment, increment/decrement

```go
a, b := 10, 3
fmt.Println(a + b)    // 13 — arithmetic
fmt.Println(a > b)    // true — comparison
fmt.Println(a & b)    // 2   — bitwise AND
a++                   // 11  — increment (statement, NOT expression)
```

**Compared to Python:**

| Category | Go | Python |
|----------|----|--------|
| Increment | `i++` (statement only) | `i += 1` |
| Exponentiation | `math.Pow(2, 3)` | `2 ** 3` |
| Integer division | `3 / 2` = `1` (int) | `3 // 2` = `1` |
| Ternary | ❌ No ternary | `x if c else y` |

---

### 23. Recursion — [`Recursion/recursion.go`](Recursion/recursion.go)

**Concepts:** recursive functions, stack limits, iteration vs recursion

```go
func fact(n int) int {
    if n == 0 { return 1 }
    return n * fact(n - 1)
}
```

| Aspect | Go | Python |
|--------|----|--------|
| Tail-call optimization | ❌ No | ❌ No |
| Default stack limit | ~1 GB (goroutine) | ~1000 frames |
| Recursion depth | Very deep (goroutine stack grows) | Shallow (fixed stack) |

**Pattern:** Go favors iteration over recursion (no perf penalty for loops).

---

### 24. Structs — [`Structs/Structs.go`](Structs/Structs.go)

**Concepts:** struct creation (4 ways), value/pointer semantics, nested/anonymous structs, struct tags, comparison, constructor pattern

```go
type User struct {
    Name string
    Age  int
}

// Creation patterns
u1 := User{"Alice", 30}                    // positional
u2 := User{Name: "Bob", Age: 25}           // named fields
u3 := new(User)                             // pointer (all fields zero)
u4 := &User{Name: "Charlie"}                // pointer with fields
```

| Feature | Go Struct | Python Class |
|---------|-----------|--------------|
| Constructor | Constructor function `NewUser()` | `__init__` |
| Zero values | Auto-initialized | `__init__` required |
| Inheritance | Embedding (composition) | Class inheritance |
| Methods | Defined separately | Inside class body |
| Tags | Struct tags (JSON, validation) | Decorators |

---

### 25. Enums (iota) — [`Enums/enums.go`](Enums/enums.go)

**Concepts:** custom types, `iota` auto-increment, `String()` method, bitmasks, skip values

```go
type Day int
const (
    Monday Day = iota  // 0
    Tuesday            // 1
    Wednesday          // 2
)

func (d Day) String() string {
    return [...]string{"Mon", "Tue", "Wed"}[d]
}
```

| Pattern | Example | Use case |
|---------|---------|----------|
| Basic iota | `iota` from 0 | Sequential constants |
| Skip values | `_`, `_` | Reserve numbers |
| Bitmasks | `Flag1 = 1 << iota` | `1, 2, 4, 8...` |

---

### 26. String Builder — [`Strings-Builder/strings_builder.go`](Strings-Builder/strings_builder.go)

**Concepts:** `strings.Builder`, `WriteString`, `Grow`, vs `+=` benchmark

```go
var sb strings.Builder
sb.Grow(100)                            // pre-allocate (performance!)
sb.WriteString("Hello")
sb.WriteString(" ")
sb.WriteString("World")
result := sb.String()                   // "Hello World"
```

| Method | Performance | Use case |
|--------|-------------|----------|
| `+=` | ❌ Slow (new string each time) | Occasional concat |
| `strings.Join` | ✅ Fast | Join slice with delimiter |
| `strings.Builder` | ✅✅ Fastest | Many appends (loops) |
| `fmt.Sprintf` | ⚠️ Medium | Formatting needed |

---

### 27. Error Handling — [`Error Handling/errors.go`](Error%20Handling/errors.go)

**Concepts:** `error` interface, `errors.New`, `fmt.Errorf`, custom error types, wrapping, `errors.Is`/`As`, panic, recover, defer

```go
// Returning errors (idiomatic Go)
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Defer — cleanup that always runs
func readFile(name string) ([]byte, error) {
    f, err := os.Open(name)
    if err != nil { return nil, err }
    defer f.Close()                    // runs on return or panic
    return io.ReadAll(f)
}
```

| Concept | Go | Python |
|---------|----|--------|
| Error creation | `errors.New("msg")` | `Exception("msg")` |
| Control flow | `if err != nil` | `try/except` |
| Stack trace | Manual wrapping | Automatic |
| "Finally" | `defer` | `finally` |
| Crash | `panic()` | `raise` |
| Recovery | `recover()` in defer | `except` |

**Key pattern:** Go's "if err != nil" is NOT boilerplate — it's explicit error handling. Every path is visible.

---

### 28. Packages & Modules — [`Packages-Modules/`](Packages-Modules/)

**Concepts:** exported vs unexported, multi-file packages, `init()`, `go.mod`, import paths, `go get`, blank import

```go
package calculator

// Exported (capital letter)
func Add(a, b int) int { return a + b }

// Unexported (lowercase) — package-private
func helper() int { return 42 }
```

| Concept | Go | Python |
|---------|----|--------|
| Module | `go.mod` file | `pyproject.toml` |
| Import | `import "module/pkg"` | `import module.pkg` |
| Export rule | Capital letter = public | Explicit `__all__` |
| Init | `init()` per package | Module-level code |
| Blank import | `import _ "pkg"` (for side effects) | Not idiomatic |

---

### 29. Channels — [`Channels/channel.go`](Channels/channel.go)

**Concepts:** unbuffered channels, buffered channels, send/receive, close, range, goroutine synchronization

**Analogy:** Two friends (Alex & Sam) passing Lego bricks through a pipe in the wall.

```go
pipe := make(chan string)          // unbuffered — one brick at a time
shelfPipe := make(chan string, 3)  // buffered — shelf holds 3 bricks

pipe <- "🧱 Red brick"              // send (Alex pushes into pipe)
brick := <-pipe                    // receive (Sam pulls from pipe)

close(shelfPipe)                   // "No more bricks!"
for brick := range shelfPipe { }   // Keep pulling until closed
```

| Go Concept | Lego Analogy |
|---|---|
| `make(chan T)` | Empty pipe, 1 brick fits |
| `make(chan T, N)` | Pipe with shelf for N bricks |
| `ch <- value` | Push brick into pipe |
| `value := <-ch` | Pull brick from pipe |
| `close(ch)` | "That's all I have!" |
| `for v := range ch` | Keep pulling until empty |

**Compared to Python:**

| Operation | Go | Python |
|---|---|---|
| Create | `ch := make(chan T)` | `q = queue.Queue()` |
| Send | `ch <- val` | `q.put(val)` |
| Receive | `val := <-ch` | `q.get()` |
| Close | `close(ch)` | Sentinel value |

**Scenario 1:** Unbuffered — sync handoff, both must be ready
**Scenario 2:** Buffered — async up to buffer size
**Scenario 3:** Two-way channels — swap bricks between builders
**Scenario 4:** Close + range — "no more bricks"

---

### 30. Goroutines — [`GoRoutines/goroutines.go`](GoRoutines/goroutines.go)

**Concepts:** go keyword, sync.WaitGroup, closure capture, anonymous goroutines, concurrent timing

```go
go myFunction()        // Start goroutine — returns immediately

var wg sync.WaitGroup
wg.Add(1)              // Register worker
go func() {
    defer wg.Done()    // Signal completion
    // ... work ...
}()
wg.Wait()              // Wait for all workers
```

**Loop closure gotcha (Go < 1.22):**
```go
// WRONG — all goroutines see LAST value of i
for i := 0; i < 3; i++ {
    go func() { fmt.Println(i) }()  // prints 3, 3, 3
}

// RIGHT — pass as argument (copy)
for i := 0; i < 3; i++ {
    go func(id int) { fmt.Println(id) }(i)  // prints 0, 1, 2
}
```

| Concept | Explanation |
|---|---|
| `go fn()` | Start fn() as goroutine (non-blocking) |
| `sync.WaitGroup` | Counter: Add → Done → Wait |
| Stack size | ~2 KB (vs ~8 MB for OS threads) |
| Closure capture | Loop variables shared by ref — pass as arg! |

---

### 31. Concurrency — [`Concurrency/concurrency.go`](Concurrency/concurrency.go)

**Concepts:** concurrency vs parallelism, select statement, mutex, fan-out/fan-in pattern

```go
// Select — wait on multiple channels (race)
select {
case msg := <-ch1:
    fmt.Println("Cache responded first:", msg)
case msg := <-ch2:
    fmt.Println("DB responded first:", msg)
case <-time.After(100 * time.Millisecond):
    fmt.Println("Timeout!")
}

// Mutex — protect shared data
var mu sync.Mutex
mu.Lock()
sharedCounter++  // Only one goroutine at a time
mu.Unlock()

// Fan-out / Fan-in — distribute work to workers, collect results
jobs := make(chan int, 5)
results := make(chan int, 5)
// Start 3 workers (fan-out)
for w := 1; w <= 3; w++ {
    go worker(w, jobs, results)
}
// Collect results (fan-in)
for result := range results {
    fmt.Println(result)
}
```

**Concurrency vs Parallelism:**

| | Concurrency | Parallelism |
|---|---|---|
| What | Structure (dealing with many things) | Execution (doing many things) |
| How | Goroutines switch between tasks | Multiple CPU cores run simultaneously |
| Go's role | Built-in (`go` keyword + `chan`) | Runtime decides, hardware dependent |

| Python | Go |
|---|---|
| `threading.Lock()` | `sync.Mutex{}` |
| `asyncio.wait(FIRST_COMPLETED)` | `select { case ... }` |
| Complex worker pools | Simple `go fn()` + `chan` |

---

### 32. Memory (Stack vs Heap) — [`Memory/memory.go`](Memory/memory.go)

**Concepts:** stack allocation, heap allocation, escape analysis, garbage collection

```go
// STACK — returned by value (copy)
func staysOnStack() int {
    x := 42        // stays on stack
    return x       // copy goes to caller
}

// HEAP — address returned (escape!)
func escapesToHeap() *int {
    x := 42        // compiler sees &x returned → HEAP
    return &x
}
```

**Visual memory map:**
```
STACK (fast, automatic):          HEAP (slower, GC-managed):
┌──────────────────────┐         ┌──────────────────────┐
│ main() frame         │         │ {Blue, 6}            │
│  ├─ heapBrick = ptr ─┼──────→  │ [10, 20, 30]         │
│  └─ house =          │         └──────────────────────┘
│     {Lego House,100} │
└──────────────────────┘
```

**Escape analysis — `go build -gcflags="-m"`**

See dedicated file: [`Memory/escape-analysis.go`](Memory/escape-analysis.go) with 10 rules showing exactly when values escape to heap:

| Trigger | Stack? | Heap? |
|---|---|---|
| `x := 42; return x` | ✅ | |
| `x := 42; return &x` | | ✅ |
| `fmt.Println(x)` | | ✅ (interface{}) |
| `global = &x` | | ✅ |
| Closure capture | | ✅ |
| `make([]int, n)` | | ✅ |

**Compared to Python:** Python has NO stack allocation for user types — everything is heap-allocated. Go's escape analysis is a major performance advantage.


### 33. Go Scheduler (G-P-M Engine) — [`Go-GPM-Engine/LEARN.md`](Go-GPM-Engine/LEARN.md)

**Concepts:** G (goroutine), P (processor), M (machine/thread), work stealing, continuation stealing, GOMAXPROCS, M:N scheduling

The Go scheduler multiplexes millions of goroutines onto a few OS threads. The three actors: **G** (task), **P** (logical CPU/desk), **M** (OS thread). Key insights:
- Each P has its own **local run queue** of Gs
- **Work stealing** — idle Ps steal half the work from busy Ps
- **Continuation stealing** — Go jumps into new goroutines immediately (cache warm)
- Blocking syscalls: M releases P, new M takes over (no CPU idle)
- Monitor with: `GODEBUG=schedtrace=1000`

### 34. Garbage Collector — [`Go-Garbage-Collector/LEARN.md`](Go-Garbage-Collector/LEARN.md)

**Concepts:** Non-generational tri-color mark-sweep, write barrier, GC assist, GOGC tuning

Go's GC is **concurrent** (most phases run alongside your program):
- **White** = unvisited (dead), **Gray** = visited, in-progress, **Black** = fully scanned (alive)
- **Write barrier** — prevents hiding a white object behind a black one during GC
- **GC Assist** — if you allocate faster than GC can mark, you help clean
- **GOGC=100** — start GC when heap doubles (tunable)
- Monitor with: `GODEBUG=gctrace=1`

### 35. Concurrency Pipelines — [`Go-Pipelines/LEARN.md`](Go-Pipelines/LEARN.md)

**Concepts:** Pipeline pattern, fan-out, fan-in, worker pools, context cancellation, rate limiting, bounded parallelism

Build scalable data processing pipelines:
- **Pipeline** = stages connected by channels (each stage = goroutine)
- **Fan-out** = distribute work across N worker goroutines
- **Fan-in** = merge N result channels into one (use WaitGroup)
| Factory assembly line analogy
- Use `context.Context` for cancellation, bounded parallelism for stability

### 36. Interface Internals (iface & eface) — [`Go-Interface-Internals/LEARN.md`](Go-Interface-Internals/LEARN.md)

**Concepts:** iface, eface, itab, type assertions, "nil is not nil" bug

Go interfaces are **two-word data structures** (16 bytes):
- **eface** (`any`): `{_type, data}` — used for empty interface
- **iface** (method interface): `{itab, data}` — itab caches method dispatch (O(1))
- **"Nil is not nil"** — returning nil pointer as interface creates non-nil iface (itab non-nil, data nil)
| Type assertion is O(1) pointer comparison
---

### 37. Generics — [`Generics/generics.go`](Generics/generics.go)

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
go run "Type Conversion/conversion.go"
go run Operators/operators.go
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
go run Recursion/recursion.go
go run Methods/Methods.go
go run Interfaces/interfaces.go
go run "Struct-Embedding/struct-embedding.go"
go run Structs/Structs.go
go run Pointers/pointers.go
go run Enums/enums.go
go run strings-and-runes/strings-and-runes.go
go run Strings-Builder/strings_builder.go
go run "Error Handling/errors.go"
go run Generics/generics.go
go run Channels/channel.go
go run GoRoutines/goroutines.go
go run Concurrency/concurrency.go
go run Memory/memory.go

# Escape analysis compiler output
go build -gcflags="-m" Memory/escape-analysis.go

# Reference-only files (package not main — read for learning)
# Variables/Declare Variables.go   — package variables
# Variables/Constants.go          — package variables
# Interface/Interfaces.go         — older duplicate, see Interfaces/
# Structs/Structs.go              — now fully implemented
# Packages-Modules/main.go        — module main (read for package organization)
```

> **Note:** Some files in `Variables/` and `Interface/` use non-main packages for naming or organizational reasons. They're meant for reading/reference, not standalone execution. Most files in this repo are `package main` and fully runnable.

---

## 📈 Progress Overview

Check out the full [[Roadmap]] in the Second Brain vault for a complete topic tracker across 105 items. Current status:

| Level | Coverage |
|---|---|
| Level 1 — Core Go (37 topics) | 36 ✅ (97%) |
| Level 2 — Production Go (27 topics) | 8 ✅ (30% — Concurrency + Runtime) |
| Level 3 — Industry Backend (41 topics) | ❌ Not started |
| **Total** | **42% (44/105)** |

**Up next:** Standard Library → File I/O → JSON → HTTP → Testing → Level 3 backend topics.

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
