// ============================================================
//  IF-ELSE DECISIONS IN GO — with Python Comparisons
//  Go has NO ternary (? :) operator — all branching uses if/else
// ============================================================
//
// ┌──────────────────────┬────────────────────────────────────────────┐
// │        Go            │                 Python                     │
// ├──────────────────────┼────────────────────────────────────────────┤
// │ if condition {       │ if condition:                              │
// │     ...              │     ...                                    │
// │ } else {             │ else:                                      │
// │     ...              │     ...                                    │
// │ }                    │                                            │
// │                      │                                            │
// │ BRACES REQUIRED      │ COLON + INDENT required                    │
// │ Condition NO parens  │ Condition NO parens (both agree!)          │
// │                      │                                            │
// │ No ternary operator  │ value_if_true if cond else value_if_false  │
// │ (?:)                 │ Python HAS ternary (Go doesn't)            │
// │                      │                                            │
// │ if s; cond { }       │ No short-statement form (no equivalent)    │
// │ s scoped to if-block │ Must declare var before if                 │
// └──────────────────────┴────────────────────────────────────────────┘
// ============================================================

package main

import ("fmt")

func main() {

	// ============================================================
	//  EXAMPLE 1: if / else
	//  Condition: 7 % 2 == 0  (is 7 even?)
	// ============================================================
	// The modulo operator (%) returns the remainder of division.
	// 7 % 2 = 1 (not 0), so the condition is false.
	// The 'else' branch runs instead of the 'if' branch.
	//
	// Python:
	//   if 7 % 2 == 0:
	//       print("7 is even")
	//   else:
	//       print("7 is odd")
	//
	// Differences:
	//   - Go: BRACES { } around each branch (required)
	//   - Python: COLON (:) + indentation (required)
	//   - Go: else on SAME line as closing brace }
	//   - Python: else at SAME indentation as if
	// ============================================================
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	// Output:
	//   7 is odd

	// ============================================================
	//  EXAMPLE 2: if only (no else)
	//  Condition: 8 % 4 == 0  (is 8 divisible by 4?)
	// ============================================================
	// When you only need to handle the true case, you can
	// omit the 'else' block entirely.
	//
	// Python:
	//   if 8 % 4 == 0:
	//       print("8 is divisible by 4")
	// ============================================================
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}
	// Output:
	//   8 is divisible by 4

	// ============================================================
	//  EXAMPLE 3: Logical OR (||) operator
	//  Condition: 8 % 2 == 0  ||  7 % 2 == 0
	// ============================================================
	// The || operator returns true if EITHER side is true.
	// 8 % 2 == 0  -> true  (8 is even)
	// 7 % 2 == 0  -> false (7 is odd)
	// Since one side is true, the whole condition is true.
	//
	// Python:
	//   if 8 % 2 == 0 or 7 % 2 == 0:       # "or" not "||"
	//       print("either 8 or 7 are even")
	//
	// Operator differences:
	//   Go:      &&    ||    !
	//   Python:  and   or    not
	//   Go uses C-style operators; Python uses English words.
	// ============================================================
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}
	// Output:
	//   either 8 or 7 are even

	// ============================================================
	//  EXAMPLE 4: if with short statement + else if chain
	//  Syntax: if statement; condition { ... }
	// ============================================================
	// Go allows a short statement (num := 9) before the condition.
	// The variable 'num' is scoped to the entire if/else chain.
	//
	// Flow:
	//   1. num := 9  (runs first)
	//   2. num < 0?  -> false (9 is not negative), skip to else if
	//   3. num < 10? -> true  (9 < 10), execute this block
	//   4. else block is skipped
	//
	// This pattern is cleaner than declaring num outside the if.
	//
	// Python has NO equivalent of this short-statement form.
	// You must declare the variable before the if:
	//   num = 9
	//   if num < 0:
	//       print(num, "is negative")
	//   elif num < 10:
	//       print(num, "has 1 digit")
	//   else:
	//       print(num, "has multiple digits")
	//
	// Key difference: num is scoped to the if-block in Go — it
	// doesn't leak to the outer scope. In Python, num is always
	// in the function scope.
	//
	// Go uses "else if" (two words), Python uses "elif" (one word).
	// ============================================================
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	// Output:
	//   9 has 1 digit
	// Note: num is NOT accessible after this line!
	// In Python, num would still be in scope.

	// ============================================================
	//  COMPARISON: No Ternary Operator
	// ============================================================
	// Go has NO ternary (cond ? a : b) operator.
	// Python DOES: a if cond else b
	//
	// Python:
	//   status = "even" if x % 2 == 0 else "odd"
	//
	// Go:
	//   var status string
	//   if x % 2 == 0 {
	//       status = "even"
	//   } else {
	//       status = "odd"
	//   }
	//
	// This is INTENTIONAL — the Go authors felt ternaries lead
	// to unreadable code. In practice, you get used to writing
	// full if/else for conditional assignments.
	// ============================================================

	// ============================================================
	//  COMPARISON: Boolean values
	// ============================================================
	// Go:  true / false  (lowercase)
	// Python: True / False (capitalized)
	//
	// Go:  if b { ... }          — b must be a BOOL
	// Python: if b: ...          — b can be ANY truthy/falsy value
	//
	// Go does NOT have truthy/falsy coercion outside of bools.
	// Python treats 0, "", [], None, etc. as false in conditions.
	// In Go, if 0 { ... } is a COMPILE ERROR (0 is int, not bool).
	// You must write: if x != 0 { ... }
	// ============================================================

	// ============================================================
	//  SUMMARY: if/else Decision Table (with Python comparison)
	// ============================================================
	//  ┌──────────────────────┬──────────────────────┬──────────────────────────┐
	//  │ Form                 │ Go                   │ Python                    │
	//  ├──────────────────────┼──────────────────────┼──────────────────────────┤
	//  │ if only              │ if c { }             │ if c:                    │
	//  │ if/else              │ if c { } else { }    │ if c: ... else:           │
	//  │ Multi-branch         │ if c { } else if { } │ if c: ... elif:           │
	//  │ Short statement      │ if s; c { }          │ Not supported            │
	//  │ Ternary              │ Not supported        │ x if c else y            │
	//  │ Scoped variable      │ if-scoped            │ Function-scoped          │
	//  │ Boolean in cond      │ Only bool            │ Any truthy/falsy value   │
	//  │ Logical operators    │ &&  ||  !            │ and  or  not             │
	//  └──────────────────────┴──────────────────────┴──────────────────────────┘
	//
	// Bottom line: Go's if/else is more RESTRICTIVE than Python's:
	//   - No ternary operator
	//   - Conditions must be boolean (no truthy/falsy coercion)
	//   - But Go adds: short statements with scoped variables
	// ============================================================
}
