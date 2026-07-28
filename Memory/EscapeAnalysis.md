---
tags: [go, memory, escape-analysis, type/learning]
created: 2026-07-28
status: active
project: go-lan
updated: 2026-07-28
---

# Escape Analysis in Go

## 1. What is it?

Escape analysis is Go compiler's automatic decision process that determines whether a value should be stored on the stack (fast, function-scoped) or the heap (slower, garbage-collected). It analyzes where pointers go — if a value's address "escapes" the function, it goes to the heap.

## 2. Why do we need it?

Without escape analysis, you would have to manually decide stack vs heap for every variable. Go automates this so you write simpler code. The compiler optimizes for performance: stack allocation is near-zero cost, heap allocation triggers garbage collection overhead.

## 3. Python Comparison

| Aspect | Go | Python |
|---|---|---|
| Stack allocation | ints, structs, arrays (when no escape) | NONE for user types |
| Heap allocation | When address escapes, slices, maps | EVERYTHING is heap |
| Decision | Compiler (escape analysis) | Always heap |
| User control | Indirect (avoid &, use values) | None |
| Tool | go build -gcflags="-m" | id() (object identity) |

## 4. Syntax

No syntax for escape analysis — it's a compiler optimization. But triggers involve:

```go
return &x          // x escapes!
fmt.Println(&x)    // x escapes! (address passed to interface{})
global = &x        // x escapes!
go func() { x }()  // x escapes (closure capture)
make([]int, n)     // backing array on heap
```

Inspect decisions:
```bash
go build -gcflags="-m" main.go
```

## 5. Simple Example

```go
func staysOnStack() int {
    x := 42        // stack — returned by value
    return x
}

func escapesToHeap() *int {
    x := 42        // heap! — address returned
    return &x
}
```

## 6. Python Equivalent

```python
# Python: ALWAYS heap-allocated, no escape analysis
def stays_on_stack():
    x = 42         # heap (everything is)
    return x

def escapes_to_heap():
    x = 42         # heap (everything is)
    return x       # no difference — both heap
```

## 7. Step-by-Step Execution

```
staysOnStack():
  [Stack Frame]
  +-- x = 42 (on stack)
  +-- return 42 (copy to caller)
  +-- frame popped -> x GONE

escapesToHeap():
  [Stack Frame]        [Heap]
  +-- x = 42 ----------> {42} (x moved here)
  +-- return &x          ↑
  +-- frame popped      │
                          x still ALIVE on heap!
```

## 8. Visual Explanation

```
HEAP ESCAPE TRIGGERS:

1. Return &x:
   func f() *int {       STACK               HEAP
       x := 42           f() frame
       return &x         x: ------> {42}
   }

2. Pass to interface{}:
   func f() {
       x := 42           x: ------> {42}
       fmt.Println(x)    (fmt takes interface{})
   }

3. Global assignment:
   var g *int
   func f() {
       x := 42           x: ------> {42}
       g = &x            (global holds addr forever)
   }

4. Closure capture:
   func f() func() int {
       x := 0            x: ------> {0}
       return func() int {
           x++           (closure refs x)
       }
   }

NO ESCAPE — stays on stack:
   func f() int {
       x := 42           x stays on STACK
       return x          returned by value (copy)
   }
```

## 9. Real-World Analogy

**Restaurant Kitchen:**

Stack = Chef's immediate workbench. Ingredients for the current dish. Cleaned instantly when dish leaves. Fast, no cleanup cost.

Heap = Walk-in refrigerator. Ingredients that multiple dishes need, or that need to survive between orders. Takes time to access and restock.

Escape Analysis = Chef deciding: "Do I need this after this dish? No? Keep on bench. Yes? Put in fridge."

Returning a pointer = "Save this sauce — the next dish needs it too."

## 10. Real-World Use Cases

- Performance profiling: hot functions with unexpected GC pressure -> check escape analysis
- API design: value receivers vs pointer receivers affect escape
- High-throughput servers: every allocation matters under load
- Library design: returning by value lets caller decide allocation

## 11. Common Mistakes

**Using pointers when values work** — unnecessary heap allocations:
```go
// BAD — forces heap
func NewUser() *User {
    return &User{Name: "Alice"}
}
// BETTER — let caller decide
func NewUser() User {
    return User{Name: "Alice"}
}
```

**Assuming & always means heap** — not always, but fmt.Println always escapes:
```go
func f() {
    x := 42
    fmt.Println(&x)  // x escapes because &x goes to interface{}
}
```

## 12. Best Practices

1. Write clean, readable code first — compiler optimizes well
2. Check with -gcflags="-m" only when profiling shows issues
3. Prefer value receivers for small structs
4. Use -gcflags="-m -m -m" for full detail
5. Don't obsess — most allocations don't matter
6. Focus on hot paths (inner loops, high-frequency functions)
7. Interface parameters always allocate — unavoidable
8. Return values by value unless nil is meaningful

## 13. Summary Table

| Trigger | Stack | Heap | Reason |
|---|---|---|---|
| x := 42; return x | Yes | | Returned by value |
| x := 42; return &x | | Yes | Address escapes |
| fmt.Println(x) | | Yes | interface{} param |
| global = &x | | Yes | Global holds pointer |
| closure := func() { x } | | Yes | Closure captures x |
| make([]int, n) | | Yes | Slice backing array |
| [3]int{1,2,3} | Yes | | Fixed-size array |
| LargeStruct{} | | Yes | > ~64 bytes |
| fmt.Printf("%v", x) | | Yes | interface{} box |

## 14. Key Takeaways

1. Escape analysis is automatic — Go's compiler decides stack vs heap
2. Returning address (&x) causes escape
3. Passing to interface{} causes escape
4. Global variables storing pointers cause escape
5. Closures capture variables -> heap
6. Slices always have heap-allocated backing arrays
7. fmt.Printf causes escape (interface{} params)
8. Check with: go build -gcflags="-m" main.go
9. Don't optimize prematurely — write clean code
10. Stack << Heap << GC in performance cost