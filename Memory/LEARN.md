# Memory in Go — Stack, Heap & Escape Analysis

> Complete learning guide following the Go-by-Python teaching method.

---

## 1. What is it?

Go manages memory automatically. Values live in one of two places:

- **Stack** — a fast, self-cleaning memory region per goroutine. Local variables go here by default. When a function returns, its stack space is instantly reclaimed.
- **Heap** — a slower, persistent memory region shared by all goroutines. Values here survive until nothing references them, then the **garbage collector (GC)** frees them.

**Escape analysis** is the compiler's automatic decision process that determines where each value goes. If a value's address "escapes" the current function (returned, stored in a global, passed to an interface), the compiler moves it to the heap. Otherwise, it stays on the stack.

## 2. Why do we need it?

**Problem:** Stack space is limited and lives only as long as a function. Heap allocations need garbage collection, which costs CPU time. In languages like C, you manually decide (malloc/free). In Python, everything is heap — no choice.

**Go's solution:** The compiler automatically decides, giving you:
- **Performance** — most values stay on stack (near-zero cost)
- **Simplicity** — you write normal code, compiler handles memory placement
- **Safety** — no manual memory management, no use-after-free bugs

**Python comparison:** Python has NO escape analysis. Every integer, string, and object is heap-allocated. Go's compiler can put simple values on the stack, which is a major performance advantage.

## 3. Python Comparison

| Aspect | Go | Python |
|---|---|---|
| Small ints | Stack (if no escape) | Heap (small ints cached -5 to 257) |
| Structs/objects | Stack (if no &) | Heap (always) |
| Strings | Depends (escape analysis) | Heap (always) |
| Arrays | Stack (if no escape) | Heap (list always heap) |
| Slices | Header: stack, Data: heap | N/A |
| Memory decision | Compiler (escape analysis) | Always heap |
| GC collector | Concurrent, low-pause | Reference counting + generational |
| User control | Indirect (avoid &, use values) | None |
| Memory inspection | `go build -gcflags="-m"` | `id()`, `gc.get_objects()` |

## 4. Syntax

There's no syntax for stack vs heap — Go decides automatically. But the code you write **triggers** the decision:

```go
x := 42              // stack (local value, no address taken)
p := &x              // p may be heap if p escapes

return x             // stack: returned by value (copy)
return &x            // heap: address returned → escape!

fmt.Println(x)       // heap! x goes to interface{} parameter

var global *int
global = &x          // heap! x stored in package-level variable

func() func() int {
    x := 0
    return func() int {   // heap! closure captures x
        x++
        return x
    }
}

make([]int, n)       // heap! slice backing array
make(map[string]int) // heap! map internal structures

[3]int{1, 2, 3}      // stack (fixed-size array, no escape)
```

Inspect compiler decisions with:
```bash
go build -gcflags="-m" main.go        # basic
go build -gcflags="-m -m -m" main.go  # very detailed
```

## 5. Simple Example

```go
package main

import "fmt"

// staysOnStack — value is returned by copy
func staysOnStack() int {
    x := 42        // stack
    return x       // copy goes to caller, original x is gone
}

// escapesToHeap — address is returned
func escapesToHeap() *int {
    x := 42        // compiler sees &x returned → MOVES TO HEAP
    return &x      // address leaves function → escape!
}

func main() {
    a := staysOnStack()       // a = 42 (copy)
    b := escapesToHeap()      // b points to heap {42}
    fmt.Println(a, *b)        // both print 42
}
```

**Line-by-line:**
1. `staysOnStack()` — x lives on stack, gets copied to the caller as return value
2. `escapesToHeap()` — x would be on stack, but `return &x` forces it to heap
3. After `staysOnStack()` returns, its x is gone (stack popped)
4. After `escapesToHeap()` returns, the heap {42} is still alive

## 6. Python Equivalent

```python
# Python: EVERYTHING is heap-allocated
# There is NO stack vs heap distinction for user code

def stays_on_stack():
    x = 42        # heap (always)
    return x      # returns reference to heap object

def escapes_to_heap():
    x = 42        # heap (always, same as above)
    return x      # also returns reference to heap object

# There is NO difference between these two functions in Python!
# Both allocate x on the heap. No escape analysis exists.
```

**Key insight:** Python code NEVER runs faster by avoiding pointers because there ARE no value semantics for user types. Go code can be 2-10x faster partly because of stack allocation.

## 7. Step-by-Step Execution

```
staysOnStack():
  ┌───────────────────────┐
  │  STACK FRAME          │
  │  x = 42               │  ← stored on stack
  │  return 42 (copy)     │  ← value copied to caller's stack
  │  ── frame popped ──   │  ← x is GONE
  └───────────────────────┘

escapesToHeap():
  ┌───────────────────────┐   ┌────────────────────┐
  │  STACK FRAME          │   │  HEAP              │
  │  x = — (pointer) ────│──►│  {42}              │
  │  return &x            │   │  (x moved here)    │
  │  ── frame popped ──   │   │                    │
  └───────────────────────┘   │  x still ALIVE!    │
                              └────────────────────┘

The stack frame is gone, but the heap value {42} persists.
Caller's *int points to the heap {42}.
```

## 8. Visual Explanation

```
STACK (per goroutine, LIFO):

  High addr
  ┌─────────────────────────────┐
  │  main() frame               │
  │  ├─ a: int = 42             │
  │  ├─ b: *int ────────        │
  │  └─ house: struct{Name,Cost}│
  ├─────────────────────────────┤
  │  buildOnStack() frame       │  ← GONE after return
  │  └─ brick: {Red, 4}        │
  ├─────────────────────────────┤
  │  buildOnHeap() frame        │  ← GONE, but brick STILL
  │  └─ &brick ────────►       │    alive on heap
  └─────────────────────────────┘
  Low addr

HEAP (shared, garbage collected):

  ┌─────────────────────────────┐
  │  {Blue, 6}  ← buildOnHeap()│
  │             returned &brick │
  ├─────────────────────────────┤
  │  [10, 20, 30]  ← slice     │
  │             backing array   │
  └─────────────────────────────┘

ESCAPE ANALYSIS TRIGGERS:

  Return &x:
  func f() *int {     STACK         HEAP
    x := 42           [f]          {42}
    return &x         x: ─────►
  }                   (popped)    alive!

  Pass to interface:
  func f() {          STACK         HEAP
    x := 42           [f]          {42}
    fmt.Println(x)    x: ─────►
  }                   (popped)    alive until GC

  Global assignment:
  var g *int
  func f() {          STACK         HEAP
    x := 42           [f]          {42}
    g = &x            x: ─────►
  }                   (popped)    global holds ref

  Closure capture:
  func f() func() {   STACK         HEAP
    x := 0            [f]          {0}
    return func() {   x: ─────►
        x++           (popped)    closure refs x
    }
  }
```

## 9. Real-World Analogy

**Workshop Bench (Stack) vs Storage Room (Heap):**

| Concept | Analogy |
|---|---|
| **Stack** | Your **workbench**. Tools for the current task. Clean, organized. You put down a tool, use it, clean up. When your shift ends, the bench is cleared instantly. Fast, no overhead. |
| **Heap** | The **storage room**. Bigger, shared by everyone. You go get something, bring it to your bench, return it later. Or someone else might need it — so it stays on a shelf. Takes time to access and organize. |
| **Escape analysis** | The **foreman** deciding: "Will anyone else need this after your shift? No → leave it at your bench. Yes → put it in storage." |
| **Return &x** | "Store this in the shared cabinet — the next shift needs it." |
| **Closure capture** | "Save this blueprint. A coworker will reference it later." |
| **fmt.Println(x)** | "Hand this to the front desk (they may distribute it to anyone)." |
| **GC** | **Janitor**. Periodically checks the storage room and throws away anything nobody is using. |

## 10. Real-World Use Cases

| Use Case | Why Memory Matters |
|---|---|
| **Web servers (high throughput)** | Every request handler should minimize heap allocations. Hot path values on stack = faster, less GC pressure. |
| **Real-time systems** | GC pauses can cause latency spikes. Stack allocation avoids GC entirely. |
| **Data processing pipelines** | Large data must be on heap (it's big). But small intermediate values should stay on stack. |
| **Library design** | Return by value (`User`) vs pointer (`*User`) affects caller's allocation pattern. |
| **Game engines** | Zero-allocation hot paths are critical for 60fps. Escape analysis decisions matter. |

## 11. Common Beginner Mistakes

**Mistake 1: Returning pointers when values work**
```go
// ❌ Forces heap allocation every call
func NewUser() *User {
    return &User{Name: "Alice"}
}

// ✅ Let CALLER decide
func NewUser() User {
    return User{Name: "Alice"}  // stack, unless caller takes & of it
}
```

**Mistake 2: Surprised that fmt.Println always escapes**
```go
func hotLoop() {
    for i := 0; i < 1000000; i++ {
        fmt.Println(i)  // Each i escapes to heap! (fmt takes interface{})
    }
}
```
**Fix:** For hot paths, use batch formatting: `fmt.Sprintf` with `\n` and print once.

**Mistake 3: Assuming & always means heap**
```go
func f() {
    x := 42
    _ = &x  // x might NOT escape if &x doesn't leave f()
    // &x is used locally but never leaves the function
    // Compiler may inline it on stack
}
```
**Reality:** `fmt.Println(&x)` → escapes (passed to interface{}). But `_ = &x` → might NOT escape (Go 1.x compiler keeps it on stack).

**Mistake 4: Premature optimization**
```go
// ❌ Ugly code to try to avoid escapes
func process(items []int) []int {
    result := make([]int, 0, len(items))
    for _, v := range items {
        result = append(result, v*2)
    }
    return result
}
// The slice was going to heap anyway. Write clean code first.
```

**Mistake 5: Ignoring escape analysis in hot paths**
```go
// In a high-frequency function
func handleRequest(r *Request) Response {
    var buf bytes.Buffer
    buf.WriteString(r.Name)
    // If buf's internal buffer grows, it goes to heap
    // Pre-allocate: buf.Grow(1024) to avoid reallocation
}
```

## 12. Best Practices

1. **Write clean, readable code first** — Go's escape analysis is excellent
2. **Prefer value receivers** for small structs (< 64 bytes)
3. **Check with `-gcflags="-m"`** only when profiling shows a problem
4. **Return values by value** unless nil is meaningful
5. **Pre-allocate slices**: `make([]T, 0, expectedLen)` reduces heap reallocations
6. **Use `-gcflags="-m -m -m"`** for full escape analysis detail
7. **Don't prematurely optimize** — most allocations don't measurably affect performance
8. **Large structs** (> 64 bytes) may be heap-allocated regardless — accept it
9. **Interface parameters always allocate** — unavoidable cost of polymorphism
10. **Hot paths** (inner loops, high-frequency functions) are where escape analysis matters most

## 13. Summary Table

| Go Code | Stack? | Heap? | Why |
|---|---|---|---|
| `x := 42; return x` | ✅ | | Returned by value (copy) |
| `x := 42; return &x` | | ✅ | Address escapes function |
| `fmt.Println(x)` | | ✅ | interface{} parameter |
| `global = &x` | | ✅ | Global holds pointer forever |
| Closure captures `x` | | ✅ | Closure may outlive function |
| `make([]int, n)` | | ✅ | Slice backing array |
| `make(map[K]V)` | | ✅ | Map internal structures |
| `[3]int{1, 2, 3}` | ✅ | | Fixed-size array |
| `s := struct{X int}{42}` | ✅ | | Struct, no address taken |
| `s := &struct{X int}{42}` | | ✅ | Explicit pointer |
| `LargeStruct{}` (>64 B) | | ✅ | Too big for stack |

## 14. Key Takeaways

1. **Stack** = fast, automatic, function-scoped. **Heap** = slower, persistent, GC-managed.
2. **Escape analysis** is automatic — Go's compiler decides where each value goes
3. **Returning an address** (`&x`) is the most common escape trigger
4. **Passing to `interface{}`** causes escape (affects all `fmt` calls)
5. **Global variables** storing pointers cause escape
6. **Closures** capture variables → heap
7. **Slices** have heap-allocated backing arrays (header is on stack)
8. **Check escapes**: `go build -gcflags="-m" main.go`
9. **Don't optimize prematurely** — the compiler is smarter than you
10. **Stack << Heap << GC pause** in performance cost

---

## Practice Exercises

### Easy: Check Escape Analysis
Write a function that takes an int, adds 10, and returns it as a value. Write another function that does the same but returns a pointer. Run `go build -gcflags="-m"` on both. What does the compiler say?

### Medium: Interface Escape
Write a function `printValue(v interface{})` that just calls `fmt.Println(v)`. Call it with an int, a string, and a struct. Check the escape analysis output. Then create a version that uses generics `[T any]` — does the value still escape?

### Challenging: Optimize a Hot Path
Given this hot-path function:
```go
func processUsers(users []User) []Summary {
    var summaries []Summary
    for _, u := range users {
        summaries = append(summaries, Summary{
            Name:  u.FirstName + " " + u.LastName,
            Age:   u.Age,
            Email: u.Email,
        })
    }
    return summaries
}
```
Check escape analysis. Then optimize it:
1. Pre-allocate `summaries` with known capacity
2. Avoid the string concatenation escape (use a fixed buffer)
3. Check if the optimization actually reduced heap allocations
(Note: string concatenation in Go always allocates — the exercise is about slice pre-allocation)
