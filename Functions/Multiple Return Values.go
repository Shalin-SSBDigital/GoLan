// ============================================================
//  MULTIPLE RETURN VALUES IN GO — with Python Comparisons
// ============================================================
//  Go functions can return MULTIPLE values — a signature feature.
//
// ┌──────────────────────┬──────────────────────────────────────────────┐
// │        Go            │                  Python                      │
// ├──────────────────────┼──────────────────────────────────────────────┤
// │ func values()        │ def values():                                │
// │     (int, int) {     │    return 3, 7                               │
// │     return 3, 7      │                                              │
// │ }                    │                                              │
// │                      │                                              │
// │ a, b := values()     │ a, b = values()  — same destructuring!       │
// │                      │                                              │
// │ Go: MULTIPLE RETURN  │ Python: returns a SINGLE TUPLE               │
// │ values are a FIRST-  │ (3, 7) that gets UNPACKED into a, b          │
// │ CLASS feature, not   │ In Python: type(values()) -> tuple           │
// │ a tuple!             │ In Go:     type(values()) -> (int, int)      │
// │                      │                                              │
// │ Most common use:     │ Python typically uses exceptions or          │
// │ returning (result,   │ returning None on failure. Go uses           │
// │ error) pairs         │ (result, error) — no exceptions.            │
// └──────────────────────┴──────────────────────────────────────────────┘
// ============================================================

package main

import "fmt"

// ============================================================
//  FUNCTION WITH MULTIPLE RETURN VALUES
//  Syntax: func name(params) (returnType1, returnType2) { ... }
// ============================================================
// This function returns TWO integers. The caller receives both.
// The return types are in parentheses, separated by commas.
//
// Python equivalent:
//   def values() -> tuple[int, int]:
//       return 3, 7
//
// IMPORTANT DIFFERENCE:
//   Go:   (int, int) — these are TWO distinct return values.
//         The function truly returns two separate ints.
//   Python: return 3, 7 — this returns ONE tuple (3, 7).
//         Python's "multiple return" is actually a single tuple
//         being unpacked by the caller.
//
// Proof:
//   Go:   a, b := values()   ← two separate variables
//         x := values()       ← COMPILE ERROR: can't assign
//                               (int, int) to a single variable
//   Python: a, b = values()   ← tuple unpacking
//           x = values()      ← x is the tuple (3, 7) — WORKS!
// ============================================================
func values() (int, int) {
	return 3, 7
}

func main() {
	// ============================================================
	//  RECEIVING MULTIPLE RETURN VALUES
	// ============================================================
	// a gets 3, b gets 7. Assignment is positional.
	//
	// Python:
	//   a, b = values()
	//   # Same syntax! But Python is unpacking a tuple.
	//   # Go is directly assigning two return values.
	// ============================================================
	a, b := values()
	fmt.Println(a)
	fmt.Println(b)
	// Output:
	//   3
	//   7

	// ============================================================
	//  DISCARDING VALUES WITH _
	// ============================================================
	// The blank identifier _ discards a return value.
	// This is like Python's convention of using _ for unused values.
	//
	// Python:
	//   a, _ = values()     # same! _ is a valid variable name
	//   # In Python, _ is just a convention (like any other var).
	//   # In Go, _ is a BUILT-IN blank identifier — you CANNOT
	//   # read from it. It's a true "throwaway."
	// ============================================================
	onlyFirst, _ := values()
	fmt.Println("Only first:", onlyFirst)
	// Output: Only first: 3

	// ============================================================
	//  THE ERROR PATTERN (idiomatic Go)
	// ============================================================
	// Go has NO exceptions (like Python's try/except).
	// Instead, functions that can fail return a (result, error) pair.
	//
	// Python:
	//   def divide(a, b):
	//       if b == 0:
	//           raise ValueError("division by zero")
	//       return a / b
	//   try:
	//       result = divide(10, 0)
	//   except ValueError as e:
	//       print(e)
	//
	// Go:
	//   func divide(a, b float64) (float64, error) {
	//       if b == 0 {
	//           return 0, errors.New("division by zero")
	//       }
	//       return a / b, nil
	//   }
	//   result, err := divide(10, 0)
	//   if err != nil {
	//       fmt.Println(err)
	//   }
	//
	// This (result, error) pattern is EVERYWHERE in Go.
	// Unlike Python's try/except, it's explicit and visible in
	// the function signature. You CAN'T hide an error in Go.
	// ============================================================

	// ============================================================
	//  SUMMARY: Multiple Return Values
	// ============================================================
	//  ┌────────────────────┬──────────────────────┬──────────────────────────┐
	//  │ Feature            │ Go                   │ Python                    │
	//  ├────────────────────┼──────────────────────┼──────────────────────────┤
	//  │ Mechanism          │ True multiple returns│ Single tuple + unpacking │
	//  │ Syntax             │ func() (A, B)        │ def f(): return a, b     │
	//  │ Destructuring      │ a, b := f()          │ a, b = f()               │
	//  │ Single var capture │ COMPILE ERROR        │ Works — gets the tuple   │
	//  │ Error handling     │ result, err := f()   │ try/except + raise       │
	//  │ Named returns      │ Supported (bare ret) │ Not supported            │
	//  └────────────────────┴──────────────────────┴──────────────────────────┘
	//
	// Bottom line: Go's multiple return values are NOT tuples — they are
	// true multiple values baked into the language. This is most famously
	// used for Go's error-handling pattern (result, error). Python returns
	// a single tuple that gets unpacked — similar syntax, different semantics.
	// ============================================================
}
