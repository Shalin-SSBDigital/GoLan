// ============================================================
//  VARIABLE DECLARATION IN GO — with Python Comparisons
//  Go is statically typed — variable types are checked at
//  compile time. There are multiple ways to declare variables.
// ============================================================
//
// ┌──────────────────────┬────────────────────────────────────────────────┐
// │        Go            │                 Python                         │
// ├──────────────────────┼────────────────────────────────────────────────┤
// │ var x int = 5        │ x: int = 5  # type annotation (OPTIONAL)      │
// │ x := 5               │ x = 5       # no annotation needed             │
// │                      │                                                │
// │ TYPE AFTER name      │ TYPE AFTER colon (annotation)                  │
// │ Types are REQUIRED   │ Types are OPTIONAL (dynamic typing)            │
// │ unless inferred      │                                                │
// │                      │                                                │
// │ var x int            │ x = 0  # but x must already exist              │
// │ (zero value: 0)      │ Python has NO zero-value initialization        │
// │                      │ for undeclared variables                       │
// │                      │                                                │
// │ := only inside funcs │ = anywhere (no equivalent of := scoping)      │
// │ var at package/func  │                                                │
// └──────────────────────┴────────────────────────────────────────────────┘
// ============================================================

package variables

// ============================================================
//  GO DATA TYPES (Overview)
// ============================================================
//  ┌──────────┬──────────────────────────────────────────────┬────────────────────────────┐
//  │ Type     │ Description                                  │ Python equivalent           │
//  ├──────────┼──────────────────────────────────────────────┼────────────────────────────┤
//  │ int      │ Whole numbers (positive/zero/negative)       │ int (arbitrary precision)  │
//  │ float32  │ Decimal numbers (single precision)           │ float (always 64-bit)      │
//  │ float64  │ Decimal numbers (double precision)           │ float (same as float64)    │
//  │ string   │ Text — always wrapped in double quotes ""    │ str (single/double/triple) │
//  │ bool     │ Logical — only true or false                 │ bool (True/False, caps!)   │
//  │ byte     │ uint8 (0-255)                               │ int (no separate byte type)│
//  │ rune     │ Unicode code point (int32)                   │ ord() returns int          │
//  └──────────┴──────────────────────────────────────────────┴────────────────────────────┘
//
// KEY DIFFERENCES:
//   - Go int is FIXED precision (64-bit on 64-bit systems).
//     Python int is ARBITRARY precision — no overflow.
//   - Go bool: true/false (lowercase). Python: True/False (capitalized).
//   - Go strings: double quotes ONLY. Python: single, double, triple.
//   - Go has byte and rune types. Python has no separate byte/rune types.
// ============================================================

// ============================================================
//  METHOD 1: var keyword with explicit type
//  Syntax:   var variablename type = value
// ============================================================
// The 'var' keyword is followed by the variable name, then
// the type, then = and the value. You can also omit the value
// (it defaults to the zero-value: 0 for int, "" for string,
// false for bool, 0.0 for float).
//
// Usage: when you want to be explicit about the type, or when
// you're declaring a package-level variable (:= is not allowed
// outside functions).
//
// Python equivalent:
//   shalin: int = 25   # type annotation (optional, not enforced)
//   shalin = 25         # no annotation (dynamic typing)
//
// CRITICAL DIFFERENCE:
//   - Go:  var shalin int    → shalin is 0 (zero value, always safe)
//   - Python: shalin            → NameError! Python has no zero value.
//          shalin = None       → you must explicitly initialize
//          shalin: int         → SyntaxError! annotations need a value
// ============================================================
var shalin int = 25   // Explicit: var name type = value

// ============================================================
//  METHOD 2: Short declaration with :=
//  Syntax:   variablename := value
// ============================================================
// The := operator declares AND assigns in one step. Go infers
// the type from the value on the right side.
//
// Rules:
//   - CANNOT be used outside of functions (package-level only
//     accepts var or const)
//   - At least ONE variable on the left must be new (no pure
//     reassignment with :=)
//   - Preferred style inside functions for brevity
//
// Example (inside a function):
//   name := "Shalin"       // inferred as string
//   age  := 25             // inferred as int
//   pi   := 3.14           // inferred as float64
//   ok   := true           // inferred as bool
//
// Python equivalent:
//   name = "Shalin"        # same simplicity
//   age = 25
//
// DIFFERENCE:
//   Go:  := always DECLARES a new variable
//   Python: = either declares (if new) or reassigns (if existing)
//   Go distinguishes between declaration and assignment.
//   Python does NOT — the same = operator does both.
// ============================================================
// (see accompanying files for function-scoped examples)

// ============================================================
//  COMPARISON: var vs := vs Python
// ============================================================
//  ┌──────────────┬──────────────────────┬──────────────────────┬─────────────────────────┐
//  │              │ var                  │ :=                   │ Python                  │
//  ├──────────────┼──────────────────────┼──────────────────────┼─────────────────────────┤
//  │ Scope        │ package or function  │ function only        │ function or global       │
//  │ Type         │ explicit or inferred │ always inferred      │ always inferred          │
//  │ Zero value   │ supported            │ not supported        │ NOT supported (NameError)│
//  │ Reassign     │ no (new variable)    │ no (new variable)    │ yes (same = operator)   │
//  │ Must init?   │ no (zero value used) │ YES (value required) │ YES (or NameError)      │
//  │ Outside func │ YES                  │ NO                   │ YES (but `global` kw)   │
//  └──────────────┴──────────────────────┴──────────────────────┴─────────────────────────┘
//
// Bottom line: Go has TWO declaration forms (var and :=) with
// DIFFERENT rules. Python has ONE (=) that does everything.
// Go's zero-value initialization means you can safely declare
// a variable without initializing it. Python would raise NameError.
// ============================================================
