// ============================================================
//  PACKAGE CALCULATOR — demonstrates exported vs unexported
// ============================================================
//  Every Go file starts with `package <name>`. This declares
//  which package the file belongs to.
//
//  All files in the SAME folder MUST have the SAME package name.
//
//  ┌────────────────────┬────────────────────────────────────┐
//  │       Go           │              Python                │
//  ├────────────────────┼────────────────────────────────────┤
//  │ package calculator │ calculator/  (directory)           │
//  │                    │ __init__.py  (optional)            │
//  ├────────────────────┼────────────────────────────────────┤
//  │ Exported: Add      │ Public: def add()                  │
//  │ Unexported: add    │ Private: def _add()                │
//  ├────────────────────┼────────────────────────────────────┤
//  │ Folder = Package   │ Directory = package with __init__.py│
//  │ No __init__.py     │ __init__.py runs on import         │
//  └────────────────────┴────────────────────────────────────┘

package calculator

import "fmt"

// =============================================================================
// EXPORTED vs UNEXPORTED NAMES
// =============================================================================
// Go has NO `public` / `private` keywords.
// Instead, visibility is determined by the FIRST LETTER of the name:
//
//   UPPERCASE = EXPORTED (public)   — can be used by other packages
//   lowercase = UNEXPORTED (private) — only visible within this package
//
// Python uses underscore convention: _private
// Go ENFORCES it at compile time — unexported names CANNOT be accessed
// from other packages.

// Exported (public) — uppercase first letter
// Python equivalent: def add(a, b): (public by default)
func Add(a, b int) int {
	return a + b
}

// Exported (public) — uppercase first letter
func Subtract(a, b int) int {
	return a - b
}

// Exported — multiple return values with error
// Python: def divide(a, b): ... raise ValueError if b == 0
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// =============================================================================
// UNEXPORTED (private) FUNCTIONS
// =============================================================================
// These start with lowercase — ONLY accessible within this package.
//
// Python equivalent:
//   def _calculate_discount(price):
//       return price * 0.1
//   (just a convention — Python doesn't enforce it)

// unexported — can only be called from inside the calculator package
func calculateDiscount(price float64) float64 {
	return price * 0.1
}

// Exported function that uses an unexported helper internally
func ApplyDiscount(price float64) float64 {
	return price - calculateDiscount(price)
}

// =============================================================================
// PACKAGE-LEVEL VARIABLES (also follow the uppercase rule)
// =============================================================================

// Exported — accessible to other packages
var Version = "1.0.0"

// Unexported — only visible inside calculator package
var author = "GoLan Learner"

// Exported function accessing unexported variable internally
func GetAuthor() string {
	return author
}

// =============================================================================
// init() FUNCTION — package initialization
// =============================================================================
// Go has something similar to Python's module-level code in __init__.py.
// The init() function runs AUTOMATICALLY when the package is first imported.
//
// Python equivalent:
//   # calculator/__init__.py
//   print("calculator package initialized")
//   __all__ = ['add', 'subtract']
//
// Key properties of init():
//   1. No parameters, no return value
//   2. Runs automatically on first import
//   3. Runs BEFORE main() in the main package
//   4. Multiple init() in the same file — runs in order
//   5. init() across different files in same package — order depends on file names

func init() {
	fmt.Println("[calculator] package initialized (version:", Version, ")")
	fmt.Println("[calculator] author:", author)
}

// =============================================================================
// NOTE: Python __init__.py vs Go init()
// =============================================================================
// Python __init__.py:
//   - Requires a file for the directory to be a package
//   - Runs automatically on import
//   - Can define __all__ to control what's imported with `from pkg import *`
//   - Can run arbitrary setup code
//
// Go init():
//   - NO __init__.py equivalent — any .go file with `package name` is a package
//   - init() runs automatically on first import
//   - No __all__ equivalent — uppercase = exported, lowercase = unexported
//   - Can run arbitrary setup code (DB connection, config loading)
//   - init() in dependency packages runs BEFORE init() in importing packages
