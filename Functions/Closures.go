// ============================================================
//  CLOSURES IN GO — with Python Comparisons
// ============================================================
//  A closure is a function that references variables from outside
//  its own body — it "closes over" those variables.
//
// ┌──────────────────────┬────────────────────────────────────────────┐
// │        Go            │                 Python                     │
// ├──────────────────────┼────────────────────────────────────────────┤
// │ func counter() func()│ def counter():                             │
// │     int {            │     count = 0                              │
// │     count := 0       │     def inner():                           │
// │     return func()    │         nonlocal count  # <-- REQUIRED!    │
// │         int {        │         count += 1                         │
// │         count++      │         return count                       │
// │         return count │     return inner                           │
// │     }                │                                            │
// │ }                    │                                            │
// │                      │                                            │
// │ NO "nonlocal" needed │ "nonlocal" keyword REQUIRED to modify      │
// │ — Go captures by     │ outer variables. Without it, Python        │
// │ reference (mutation  │ creates a NEW local variable instead.      │
// │ works automatically) │                                            │
// └──────────────────────┴────────────────────────────────────────────┘
// ============================================================

package main

import "fmt"

// ============================================================
//  COUNTER CLOSURE
//  Returns a function that "remembers" and increments a count.
// ============================================================
// counter() initializes 'count' to 0 and returns an anonymous
// function that captures 'count' by reference.
//
// Each call to the returned function:
//   1. Increments count (count++)
//   2. Returns the new value
//
// The returned function is a CLOSURE because it references 'count'
// which is defined in counter()'s scope, not in its own scope.
//
// Python equivalent:
//   def counter():
//       count = 0
//       def inner():
//           nonlocal count     # ← MUST declare nonlocal!
//           count += 1
//           return count
//       return inner
//
// KEY DIFFERENCE:
//   - Go: closures can MUTATE captured variables WITHOUT any special
//     keyword. It just works.
//   - Python: you MUST declare 'nonlocal' to modify outer variables.
//     Without it, count += 1 creates a NEW local 'count' (shadowing).
//     This is a common Python gotcha that Go avoids entirely.
//
//   Reading (not mutating) captured variables works the same in both:
//   - Go:     func() int { return count }  ← OK, reads count
//   - Python: def inner(): return count      ← OK, reads count
//   No 'nonlocal' needed for read-only access in either language.
// ============================================================

func counter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func main() {
	// ============================================================
	//  USING A CLOSURE
	// ============================================================
	// next holds the closure returned by counter().
	// Each call to next() accesses and updates the SAME 'count'.
	//
	// Python:
	//   next = counter()
	//   print(next())  # 1
	//   print(next())  # 2
	//   print(next())  # 3
	// ============================================================
	next := counter()

	fmt.Println(next()) // 1
	fmt.Println(next()) // 2
	fmt.Println(next()) // 3

	// ============================================================
	//  MULTIPLE INDEPENDENT CLOSURES
	// ============================================================
	// Each call to counter() creates a FRESH 'count' variable.
	// nextA and nextB have SEPARATE, independent counters.
	//
	// Python:
	//   next_a = counter()
	//   next_b = counter()
	//   print(next_a())  # 1
	//   print(next_a())  # 2
	//   print(next_b())  # 1  ← separate counter!
	// ============================================================
	nextA := counter()
	nextB := counter()

	fmt.Println("\nIndependent closures:")
	fmt.Println("nextA:", nextA()) // 1
	fmt.Println("nextA:", nextA()) // 2
	fmt.Println("nextB:", nextB()) // 1 (separate counter)
	fmt.Println("nextA:", nextA()) // 3
	fmt.Println("nextB:", nextB()) // 2

	// ============================================================
	//  CLOSURE CAPTURE GOTCHA (Go vs Python)
	// ============================================================
	// Go shares the SAME gotcha as Python with loop variables:
	//
	// WRONG:
	//   funcs := []func() int{}
	//   for i := 0; i < 3; i++ {
	//       funcs = append(funcs, func() int { return i })
	//   }
	//   // Each closure captures the SAME 'i' — all return 3!
	//
	// Python equivalent:
	//   funcs = []
	//   for i in range(3):
	//       funcs.append(lambda: i)
	//   # All lambdas return 2 (the final i value)!
	//
	// Fix in Go:   i := i  (shadow the loop variable)
	// Fix in Python: lambda i=i: i  (capture by value as default arg)
	// ============================================================
	fmt.Println("\nClosure gotcha (both languages have this!):")
	var funcs []func() int
	for i := 0; i < 3; i++ {
		i := i // ← this shadows i, creating a NEW variable per iteration
		funcs = append(funcs, func() int { return i })
	}
	for _, f := range funcs {
		fmt.Println(f()) // 0, 1, 2 — correct because of i := i
	}

	// ============================================================
	//  SUMMARY: Go Closures vs Python Closures
	// ============================================================
	//  ┌────────────────────┬──────────────────────┬──────────────────────────┐
	//  │ Feature            │ Go                   │ Python                    │
	//  ├────────────────────┼──────────────────────┼──────────────────────────┤
	//  │ Mutating captured  │ Works automatically  │ Requires 'nonlocal' kw   │
	//  │ variables          │                      │                          │
	//  │ Read captured vars │ Works automatically  │ Works automatically      │
	//  │ Syntax             │ func() { ... }       │ def inner(): ...         │
	//  │                    │                      │ or lambda: ...           │
	//  │ Multiple lines     │ Always block { }     │ def or lambda (1 expr)  │
	//  │ Loop variable      │ i := i fix needed    │ default=arg fix needed   │
	//  │ gotcha             │                      │                          │
	//  │ Garbage collection │ Captured vars freed  │ Captured vars freed      │
	//  │                    │ when closure is GC'd │ when closure is GC'd     │
	//  └────────────────────┴──────────────────────┴──────────────────────────┘
	//
	// Bottom line: Closures work almost identically in both languages,
	// with ONE exception — Go doesn't need 'nonlocal' for mutations,
	// while Python does. Both languages share the loop-variable gotcha.
	// ============================================================
}
