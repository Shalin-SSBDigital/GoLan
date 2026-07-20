// ============================================================
//  SWITCH STATEMENTS IN GO — with Python Comparisons
// ============================================================
//
// ┌──────────────────────┬──────────────────────────────────────────────┐
// │        Go            │                 Python                       │
// ├──────────────────────┼──────────────────────────────────────────────┤
// │ switch num {         │ match num:            # Python 3.10+         │
// │ case 1:              │     case 1:                                   │
// │     fmt.Println(..)  │         print(...)                            │
// │ }                    │                                               │
// │                      │                                               │
// │ NO break needed      │ No break needed (match-case)                 │
// │ Cases DON'T fall     │ Cases DON'T fall through                     │
// │ through by default   │ by default                                    │
// │                      │                                               │
// │ fallthrough keyword  │ No explicit fallthrough (not supported)      │
// │ for explicit fall-   │                                               │
// │ through              │                                               │
// │                      │                                               │
// │ switch no expr =     │ if/elif chain (no match equivalent)          │
// │ if/else chain        │                                               │
// │ (case uses bools)    │                                               │
// │                      │                                               │
// │ Older than match-case│ match-case was added in Python 3.10 (2021).  │
// │ Go has had switch    │ Before that, Python used if/elif chains.     │
// │ since Go 1.0 (2012)  │                                               │
// └──────────────────────┴──────────────────────────────────────────────┘
// ============================================================

package main

import ("fmt")

func main() {
	// ============================================================
	//  SWITCH STATEMENT
	//  Syntax: switch expression { case value: ... }
	// ============================================================
	// A switch statement evaluates an expression and compares it
	// against multiple cases. The first matching case executes.
	// If no case matches, the optional 'default' case runs.
	//
	// Python (3.10+):
	//   match num:
	//       case 1:
	//           print("num is one")
	//       case 2:
	//           print("num is two")
	//       case 3:
	//           print("num is three")
	//       case _:           # default in Python
	//           print("num is something else")
	//
	// KEY DIFFERENCE — NO BREAK NEEDED:
	//   In Go, cases do NOT fall through by default.
	//   Once a case matches, the switch exits.
	//   In C/Java, you need "break" to prevent fallthrough.
	//   Go removes this footgun entirely.
	//   Python match-case also doesn't fall through.
	// ============================================================
	num := 2

	switch num {
	case 1:
		fmt.Println("num is one")
	case 2:
		fmt.Println("num is two")
	case 3:
		fmt.Println("num is three")
	default:
		fmt.Println("num is something else")
	}
	// Output:
	//   num is two

	// ============================================================
	//  SWITCH WITH MULTIPLE CASES (comma-separated)
	// ============================================================
	// You can group multiple values in a single case using commas.
	// The first matching case executes, and the rest are skipped.
	//
	// Python (3.10+):
	//   match day:
	//       case "Saturday" | "Sunday":    # | for OR in patterns
	//           print("It's the weekend!")
	//       case _:
	//           print("It's a weekday.")
	//
	// Note: Python uses | between patterns, Go uses commas.
	// ============================================================
	day := "Saturday"

	switch day {
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}
	// Output:
	//   It's the weekend!

	// ============================================================
	//  SWITCH WITH NO EXPRESSION — if/else chain
	//  Syntax: switch { case condition: ... }
	// ============================================================
	// When you omit the expression after "switch", each case
	// must be a BOOLEAN expression. The first true case runs.
	// This is a cleaner alternative to a long if/else if chain.
	//
	// Python:
	//   score = 85
	//   if score >= 90:
	//       grade = "A"
	//   elif score >= 80:
	//       grade = "B"
	//   elif score >= 70:
	//       grade = "C"
	//   else:
	//       grade = "D"
	//
	// This is purely a STYLISTIC choice in Go — no performance
	// difference from if/else. Use whichever reads better.
	// ============================================================
	score := 85
	var grade string
	switch {
	case score >= 90:
		grade = "A"
	case score >= 80:
		grade = "B"
	case score >= 70:
		grade = "C"
	default:
		grade = "D"
	}
	fmt.Printf("Score %d = Grade %s\n", score, grade)
	// Output: Score 85 = Grade B

	// ============================================================
	//  SWITCH WITH SHORT STATEMENT (like if)
	//  Syntax: switch s := expr; s { ... }
	// ============================================================
	// Like 'if', 'switch' can have a short statement before the
	// expression. The declared variable is scoped to the switch.
	//
	// Python has no equivalent — you'd declare the variable first.
	// ============================================================
	switch day := "Monday"; day {
	case "Saturday", "Sunday":
		fmt.Println("Weekend!")
	default:
		fmt.Println("Weekday:", day)
	}
	// Output: Weekday: Monday
	// day is NOT accessible after this point.

	// ============================================================
	//  fallthrough — explicit case chaining (rarely used)
	// ============================================================
	// Go cases DON'T fall through, but you can FORCE it with
	// the 'fallthrough' keyword. This is rarely used but exists
	// for compatibility with C-style state machines.
	//
	// Python's match-case does NOT support fallthrough at all.
	// ============================================================
	switch n := 2; n {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
		fallthrough // ← continues to the next case regardless!
	case 3:
		fmt.Println("three (also prints because of fallthrough)")
	default:
		fmt.Println("other")
	}
	// Output:
	//   two
	//   three (also prints because of fallthrough)

	// ============================================================
	//  COMPARISON: Type Switch (Go-only)
	// ============================================================
	// Go's switch can also MATCH ON TYPES using a type assertion.
	// Python has no direct equivalent (isinstance() + if/elif).
	//
	// var x interface{} = "hello"
	// switch v := x.(type) {
	// case int:
	//     fmt.Println("int:", v)
	// case string:
	//     fmt.Println("string:", v)
	// }
	// ============================================================

	// ============================================================
	//  SUMMARY: Switch (Python comparison)
	// ============================================================
	//  ┌──────────────────────┬──────────────────────┬────────────────────────────┐
	//  │ Feature              │ Go                   │ Python                     │
	//  ├──────────────────────┼──────────────────────┼────────────────────────────┤
	//  │ Syntax               │ switch x { case y: }│ match x: case y:           │
	//  │ Added                │ Go 1.0 (2012)        │ Python 3.10 (2021)        │
	//  │ Fallthrough          │ No by default        │ No by default              │
	//  │ Explicit fallthrough │ fallthrough keyword  │ Not supported              │
	//  │ Multi-value case     │ case 1, 2, 3:        │ case 1 | 2 | 3:           │
	//  │ No-expression switch │ switch { case c: }   │ if/elif chain              │
	//  │ Short statement      │ switch s:=x; s { }   │ Not supported              │
	//  │ Type switch          │ x.(type)             │ isinstance() + if/elif     │
	//  │ Default              │ default:             │ case _:                    │
	//  └──────────────────────┴──────────────────────┴────────────────────────────┘
	//
	// Bottom line: Go has had switch since day one; Python got
	// match-case in 3.10 (2021). Both prevent accidental fallthrough.
	// Go is unique for: no-expression switch (boolean cases),
	// short statement, explicit fallthrough, and type switches.
	// ============================================================
}
