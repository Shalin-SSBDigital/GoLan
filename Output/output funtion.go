// ============================================================
//  OUTPUT FUNCTIONS IN GO — with Python Comparisons
//  Printing to console using the fmt package
// ============================================================
//
// ┌──────────────────────────┬────────────────────────────────────────────┐
// │        Go (fmt)          │                 Python                     │
// ├──────────────────────────┼────────────────────────────────────────────┤
// │ fmt.Println(a, b)        │ print(a, b)     — with spaces + newline    │
// │ fmt.Print(a, b)          │ print(a, b, end="")  — no newline          │
// │ fmt.Printf("%d", x)      │ print(f"{x}")  or print("%d" % x)         │
// │                          │                                            │
// │ fmt.Sprintf("%s", s)     │ f"{s}"  or  "%s" % s  or  s.format()      │
// │                          │                                            │
// │ %d  = int                │ %d or {0} — same as C-style                │
// │ %s  = string             │ %s  — same                                 │
// │ %f  = float64            │ %f  — same                                 │
// │ %v  = any (Go-specific)  │ %r  = repr (similar, not identical)        │
// │ %T  = type name          │ type(x).__name__                          │
// │ %+v = struct with fields │ No direct equivalent                       │
// └──────────────────────────┴────────────────────────────────────────────┘
// ============================================================

package main

import ("fmt")

func main() {
	// ============================================================
	//  Variable Declaration with :=
	// ============================================================
	// Short variable declaration (:=) infers the type from the
	// value. Here both a and b are inferred as int.
	// Equivalent to: var a int = 10; var b int = 20
	//
	// Python:
	//   a, b = 10, 20      # same syntax!
	//   # BUT: Python variables have no type enforcement.
	//   # a = "hello" would work in Python but fail in Go.
	// ============================================================
	a, b := 10, 20

	// ============================================================
	//  fmt.Println — Print values with spaces between arguments
	// ============================================================
	// Println accepts multiple arguments separated by commas.
	// It inserts spaces between arguments and adds a newline at
	// the end. The expression a + b is evaluated first, then the
	// result (30) is printed.
	//
	// Python:
	//   print("The sum of a and b is:", a + b)
	//   # Same! Python's print() also inserts spaces and adds newline.
	//   # print() is a BUILT-IN (no import needed).
	//   # fmt.Println requires import "fmt".
	// ============================================================
	fmt.Println("The sum of a and b is: ", a + b)
	// Output:
	//   The sum of a and b is:  30

	fmt.Println("Hello World!")
	// Output:
	//   Hello World!

	// ============================================================
	//  fmt.Print — Print WITHOUT a trailing newline
	// ============================================================
	// Like Println but does NOT append "\n". Consecutive Print
	// calls appear on the same line.
	//
	// Python:
	//   print("Hello ", end="")
	//   print("World")     # → "Hello World" on one line
	// ============================================================
	fmt.Print("Hello ")
	fmt.Println("World!")
	// Output: Hello World!

	// ============================================================
	//  fmt.Printf — Formatted printing
	// ============================================================
	// Printf uses FORMAT VERBS (like C's printf).
	// The verb determines how the argument is displayed.
	//
	// Common verbs:
	//   %d   decimal integer
	//   %s   string
	//   %f   floating point
	//   %.2f float with 2 decimal places
	//   %v   value in default format
	//   %T   type
	//   %t   boolean
	//   %x   hex
	//
	// Python equivalent:
	//   name = "Alice"
	//   age = 30
	//   print(f"Name: {name}, Age: {age}")     # f-string (modern)
	//   print("Name: %s, Age: %d" % (name, age))  # %-formatting (old)
	//   print("Name: {}, Age: {}".format(name, age))  # .format()
	//
	// Key differences:
	//   - Python's f-strings are MORE readable than Printf verbs
	//   - Go has NO f-strings — you must use Printf or Sprintf
	//   - Go's %v verb is powerful — it prints ANY value in a
	//     reasonable default format (Python's %r is closest)
	// ============================================================
	name := "Alice"
	age := 30
	fmt.Printf("Name: %s, Age: %d\n", name, age)
	// Output: Name: Alice, Age: 30

	// ============================================================
	//  %v — the "value" verb (automatic formatting)
	// ============================================================
	// %v works for ANY type. It's useful for debugging because
	// you don't need to know the exact type.
	//
	// Python:  print(repr(x))  # %r in %-formatting
	// ============================================================
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Bob", Age: 25}
	fmt.Printf("Struct: %v\n", p)   // without field names
	fmt.Printf("Struct: %+v\n", p)  // WITH field names (Python-like __dict__)
	fmt.Printf("Type: %T\n", p)     // type name
	// Output:
	//   Struct: {Bob 25}
	//   Struct: {Name:Bob Age:25}
	//   Type: main.Person

	// ============================================================
	//  fmt.Sprintf — Format to string (like Python's f-strings)
	// ============================================================
	// Sprintf returns the formatted string instead of printing it.
	// Useful for building strings without printing.
	//
	// Python:
	//   greeting = f"My name is {name} and I am {age} years old."
	//   # Python's f-strings are much more concise and readable.
	// ============================================================
	greeting := fmt.Sprintf("My name is %s and I am %d years old.", name, age)
	fmt.Println(greeting)
	// Output: My name is Alice and I am 30 years old.

	// ============================================================
	//  SUMMARY: Output Functions (Python comparison)
	// ============================================================
	//  ┌───────────────────┬──────────────────────┬────────────────────────────┐
	//  │ Go                │ Python               │ Notes                      │
	//  ├───────────────────┼──────────────────────┼────────────────────────────┤
	//  │ Println(...)      │ print(...)           │ Both add newline + spaces  │
	//  │ Print(...)         │ print(..., end="")   │ Go: no newline            │
	//  │ Printf(fmt, args) │ print(f"{args}")     │ Go: C-style verbs          │
	//  │                   │ print("%s" % args)   │ Python: f-strings (clean) │
	//  │ Sprintf(f, args)  │ f"...{args}..."      │ Go: returns string         │
	//  │                   │ "...{}".format(args) │ Python: f-string or .fmt   │
	//  │ %v / %+v          │ repr() / %r          │ Go: default/verbose format │
	//  │ %T                │ type(x).__name__     │ Go: compile-time type name │
	//  └───────────────────┴──────────────────────┴────────────────────────────┘
	//
	// Bottom line: Go's output functions use C-style format verbs
	// and require importing "fmt". Python's print() is built-in and
	// Python's f-strings are more readable. But Go's %v and %+v
	// provide powerful default formatting for any value.
	// ============================================================
}

// ============================================================
//  USER-DEFINED FUNCTION: add
//  Parameters: a int, b int   (both integers)
//  Return:     int            (sum of a and b)
// ============================================================
// This function is defined but NOT called in main() above.
// To use it, you would call: result := add(10, 20)
//
// Key Go syntax rules:
//   - Type follows the variable name (unlike C/Java)
//   - If two consecutive params share the same type, you can
//     write: add(a, b int) int
//
// Python:
//   def add(a: int, b: int) -> int:
//       return a + b
//   # In Go, types are REQUIRED and ENFORCED.
//   # In Python, type hints are OPTIONAL and NOT enforced.
// ============================================================
func add(a int, b int) int {
	return a + b
}
