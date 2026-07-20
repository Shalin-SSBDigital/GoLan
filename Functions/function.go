// ============================================================
//  FUNCTIONS IN GO — with Python Comparisons
// ============================================================
//  Functions are first-class citizens in Go.
//
// ┌──────────────────────┬────────────────────────────────────────────┐
// │        Go            │                 Python                     │
// ├──────────────────────┼────────────────────────────────────────────┤
// │ func add(a, b int)   │ def add(a, b):                             │
// │     int {            │     return a + b                           │
// │     return a + b     │                                            │
// │ }                    │                                            │
// │                      │                                            │
// │ Type AFTER name      │ Type annotations BEFORE name (optional)    │
// │ Types REQUIRED       │ Types OPTIONAL (dynamic typing)            │
// │ Return type is       │ return type is NOT declared                │
// │ REQUIRED after params│ (annotations optional)                     │
// │                      │                                            │
// │ func(var, var) ret   │ def(var, var) -> ret:  (type hints)        │
// │ RET-type must be     │ Type hints are NOT enforced                │
// │ ENFORCED at compile  │ at runtime                                 │
// └──────────────────────┴────────────────────────────────────────────┘
// ============================================================

package main

import "fmt"

// ============================================================
//  BASIC FUNCTION — with typed parameters
//  Syntax: func name(param type, param type) returnType { ... }
// ============================================================
// This function takes two integers and returns their sum.
//
// Python equivalent:
//   def add(a: int, b: int) -> int:
//       return a + b
//
// Key differences:
//   - Go: type comes AFTER the variable name
//   - Python: type comes AFTER colon (annotation)
//   - Go: type ERRORS are caught at COMPILE time
//   - Python: type hints are OPTIONAL and NOT enforced at runtime
//   - Go: return type is MANDATORY (after param list)
//   - Python: return type annotation is OPTIONAL
//
// Consecutive params of the SAME type can be shortened:
//   func add(a, b int) int    ← "a int, b int" shortened
// ============================================================
func add(a int, b int) int {
	return a + b
}

// Python lets you return ANY type from a function regardless of what the
// type hint says. Go ENFORCES the return type at compile time.

func main4() {
	// ============================================================
	//  CALLING A FUNCTION
	// ============================================================
	// Go uses the same call syntax as Python: name(arg1, arg2)
	//
	// Python:  result = add(3, 4)
	// ============================================================
	result := add(3, 4)
	fmt.Println("Result:", result)
	// Output: Result: 7

	// ============================================================
	//  KEY DIFFERENCES FROM PYTHON
	// ============================================================
	// ╔═══════════════════════════════════╤══════════════════════════════════╗
	// ║ Go                                │ Python                          ║
	// ╠═══════════════════════════════════╪══════════════════════════════════╣
	// ║ func add(a, b int) int { ... }    │ def add(a, b): ...              ║
	// ║ Types REQUIRED for ALL params     │ Types OPTIONAL (annotation)     ║
	// ║ Return type REQUIRED              │ Return type annotation OPTIONAL ║
	// ║ Type checking at COMPILE TIME     │ Type checking at RUNTIME (if    ║
	// ║                                   │ using type hints + mypy)        ║
	// ║ func name(param type) ret         │ def name(param: type) -> type:  ║
	// ║ Type AFTER name                   │ Type AFTER colon                ║
	// ║ No default param values           │ Default param values supported  ║
	// ║ No keyword arguments              │ Keyword arguments supported     ║
	// ║ No *args or **kwargs              │ *args, **kwargs supported       ║
	// ╚═══════════════════════════════════╧══════════════════════════════════╝
	//
	// Bottom line: Go functions are STATICALLY typed — every parameter
	// and the return value MUST have a declared type. Python functions
	// are DYNAMICALLY typed — types are optional annotations only.
	// Go catches type mismatches at compile time; Python catches them
	// at runtime (or not at all).
	// ============================================================
}

// ============================================================
//  EXTRA: Named Return Values (Go-only feature)
// ============================================================
// Go allows you to NAME the return values in the function signature.
// This documents what's being returned and creates zero-valued
// variables you can use with a bare return.
//
// Python has no direct equivalent.
// Python:  def divide(a, b): return a // b, a % b
// ============================================================
func divide(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return // bare return — returns quotient and remainder
	// This is equivalent to:  return quotient, remainder
}
