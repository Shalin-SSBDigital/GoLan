// ============================================================
//  PACKAGE OVERVIEW — with Python Comparisons
// ============================================================
// package main  -> special package that tells Go to compile
//                  this as an executable (not a reusable library)
// import "fmt"  -> fmt = format — Go's standard I/O package for
//                  printing, formatting, and reading data
// func main()   -> mandatory entry point; execution starts here
//
// ┌──────────────────────┬────────────────────────────────────────────┐
// │        Go            │                 Python                     │
// ├──────────────────────┼────────────────────────────────────────────┤
// │ package main         │ No direct equivalent. Python scripts       │
// │                      │ just start executing from top to bottom.  │
// ├──────────────────────┼────────────────────────────────────────────┤
// │ import "fmt"         │ import builtins or print() auto-available  │
// │                      │ Python's print() is a BUILT-IN — no import │
// │                      │ needed. Go's fmt must be explicitly        │
// │                      │ imported.                                 │
// ├──────────────────────┼────────────────────────────────────────────┤
// │ func main() { ... }  │ if __name__ == "__main__":                 │
// │                      │     ...                                    │
// │                      │ Go: execution ALWAYS starts at func main() │
// │                      │ Python: executes top-to-bottom; the guard  │
// │                      │ idiom only runs when invoked directly      │
// │                      │ (not imported as a module)                │
// ├──────────────────────┼────────────────────────────────────────────┤
// │ fmt.Println("Hello") │ print("Hello")                             │
// │                      │                                            │
// │ Go: function MUST be │ Python: print is a built-in function       │
// │ imported (from "fmt")│ always available, no import needed.        │
// └──────────────────────┴────────────────────────────────────────────┘
// ============================================================

package main

import ("fmt")

func main() {
	// ============================================================
	//  fmt.Println — Print Line
	// ============================================================
	// Prints the given text to the console, followed by a newline.
	// The string "Hello World!" is passed as an argument and
	// displayed in the terminal when the program runs.
	//
	// Python equivalent:
	//   print("Hello World!")
	//
	// Key differences:
	//   - fmt.Println ALWAYS adds a newline at the end.
	//     Python's print() also adds a newline by default (end="\n").
	//   - fmt.Println takes multiple arguments separated by commas
	//     and inserts spaces between them: fmt.Println("a", "b") → "a b"
	//     Python does the same: print("a", "b") → "a b"
	//   - fmt.Print() prints without newline (like Python's print(..., end=""))
	//   - fmt.Printf() formats with placeholders (like Python's f-strings
	//     or "".format())
	// ============================================================
	fmt.Println("Hello World!")
	// Output:
	//   Hello World!

	// ============================================================
	//  STRUCTURE COMPARISON: Python vs Go "Hello World"
	// ============================================================
	// ╔═══════════════════════════════════╤══════════════════════════════════╗
	// ║  Go                               │  Python                         ║
	// ╠═══════════════════════════════════╪══════════════════════════════════╣
	// ║  package main                     │  # no package declaration        ║
	// ║  import "fmt"                     │  # print is built-in             ║
	// ║  func main() {                    │  if __name__ == "__main__":      ║
	// ║      fmt.Println("Hello World!")  │      print("Hello World!")       ║
	// ║  }                                │                                  ║
	// ╚═══════════════════════════════════╧══════════════════════════════════╝
	//
	// Bottom line: Go requires more STRUCTURE (package, func main, imports)
	// than Python. Python lets you jump straight to print(). But Go's
	// structure makes dependencies and entry points EXPLICIT — you always
	// know where a program starts, unlike Python where any line of code
	// at module level runs on import.
	// ============================================================
}
