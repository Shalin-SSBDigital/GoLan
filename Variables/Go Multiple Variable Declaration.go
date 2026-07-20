// ============================================================
//  MULTIPLE VARIABLE DECLARATION IN GO — with Python Comparisons
//  Declare several variables of the SAME type in one line
// ============================================================
//
// ┌───────────────────────────────┬────────────────────────────────────────────┐
// │        Go                     │                 Python                     │
// ├───────────────────────────────┼────────────────────────────────────────────┤
// │ var a, b, c int = 1, 2, 3    │ a, b, c = 1, 2, 3                          │
// │                               │                                            │
// │ ALL variables must be the     │ Variables can be ANY type (dynamic)         │
// │ SAME type when using this     │ a, b, c = 1, "hello", True  ← works        │
// │ parallel syntax               │                                            │
// │ a, b, c = 1, "s", true        │                                            │
// │   → COMPILE ERROR (different  │                                            │
// │     types with declared var)  │                                            │
// ├───────────────────────────────┼────────────────────────────────────────────┤
// │ var (                         │ # No block declaration in Python           │
// │     x int = 1                 │ x = 1; y = "hi"; z = True                  │
// │     y string = "hi"           │                                            │
// │     z bool = true             │                                            │
// │ )                             │                                            │
// │ Different types OK in blocks  │                                            │
// └───────────────────────────────┴────────────────────────────────────────────┘
// ============================================================

package variables

import ("fmt")

func main() {
	// ============================================================
	//  Parallel Variable Declaration
	//  Syntax: var name1, name2, ... type = val1, val2, ...
	// ============================================================
	// This declares four int variables (a, b, c, d) in one line.
	// The values are assigned positionally:
	//   a = 1, b = 3, c = 5, d = 7
	//
	// This is equivalent to writing:
	//   var a int = 1
	//   var b int = 3
	//   var c int = 5
	//   var d int = 7
	//
	// All variables MUST share the same type.
	// The number of names on the left must equal the number of
	// values on the right.
	//
	// Python equivalent:
	//   a, b, c, d = 1, 3, 5, 7
	//
	// KEY DIFFERENCES:
	//   - Go: type is declared ONCE for all (int)
	//   - Python: no type declaration needed (dynamic)
	//   - Go: variables can be ANY type within the SAME type
	//   - Python: each variable can be a DIFFERENT type
	//     a, b = 1, "hello"  → OK in Python
	//     var a, b int = 1, "hello"  → COMPILE ERROR in Go
	// ============================================================
	var a, b, c, d int = 1, 3, 5, 7

	// ============================================================
	//  Print each variable on its own line
	// ============================================================
	// Each call to fmt.Println prints one value followed by a
	// newline. There is no formatting — just the raw value.
	//
	// Python:
	//   print(a)
	//   print(b)
	//   print(c)
	//   print(d)
	// ============================================================
	fmt.Println(a)   // Output: 1
	fmt.Println(b)   // Output: 3
	fmt.Println(c)   // Output: 5
	fmt.Println(d)   // Output: 7

	// ============================================================
	//  ALTERNATIVE PATTERNS (Python comparison)
	// ============================================================
	// Go — var block (mixed types):
	//   var (
	//       name   string = "Go"
	//       year   int    = 2024
	//       active bool   = true
	//   )
	//
	// Python — no block declaration; use separate lines:
	//   name = "Go"
	//   year = 2024
	//   active = True
	//
	// Go — short declaration with := (mixed types):
	//   x, y, z := "a", 42, true
	//
	// Python — same syntax works:
	//   x, y, z = "a", 42, True
	//
	// DIFFERENCE: Go's := requires type consistency only within
	// inferred types (the compiler infers each variable's type
	// separately from its value). Python just assigns.

	// ============================================================
	//  SWAP — Go has built-in swap
	// ============================================================
	// Go supports direct variable swap without a temporary variable.
	//
	// Python:
	//   a, b = b, a    # same syntax!
	//
	// This is one area where Go and Python are IDENTICAL.
	// Both use tuple-style assignment for swapping.
	// ============================================================
	x, y := 10, 20
	fmt.Printf("Before swap: x=%d, y=%d\n", x, y) // x=10, y=20
	x, y = y, x
	fmt.Printf("After swap: x=%d, y=%d\n", x, y)  // x=20, y=10

	// ============================================================
	//  SUMMARY: Multiple Variable Declaration
	// ============================================================
	//  ┌────────────────────┬──────────────────────┬──────────────────────────┐
	//  │ Feature            │ Go                   │ Python                    │
	//  ├────────────────────┼──────────────────────┼──────────────────────────┤
	//  │ Parallel decl      │ var a, b int = 1, 2  │ a, b = 1, 2              │
	//  │ Mixed types        │ var block or :=      │ Allowed always           │
	//  │ Type enforcement   │ Same type required   │ No enforcement           │
	//  │                   │ for parallel var     │ (dynamic typing)         │
	//  │ Block declaration  │ var ( ... )          │ Not supported            │
	//  │ Swap               │ a, b = b, a          │ a, b = b, a              │
	//  │ := vs =            │ := declares new      │ = always (declare or     │
	//  │                   │ = reassigns existing │ reassign)                │
	//  └────────────────────┴──────────────────────┴──────────────────────────┘
	//
	// Bottom line: Go's parallel declaration is MORE RESTRICTIVE than
	// Python's (same type required), but Go's swap is IDENTICAL.
	// Go's var block is useful for grouping related declarations,
	// which Python lacks.
	// ============================================================
}
