// ============================================================
//  OPERATORS IN GO — with Python Comparisons
// ============================================================
//  Go supports the same basic operator categories as Python:
//  Arithmetic, Comparison, Logical, Assignment — plus Bitwise.
//
//  ┌──────────────────────┬──────────────────────────────────┐
//  │         Go           │             Python               │
//  ├──────────────────────┼──────────────────────────────────┤
//  │ +, -, *, /, %        │ +, -, *, /, //, %               │
//  │ ++, -- (statements)  │ ++, -- (expressions)             │
//  │ &&, ||, !            │ and, or, not                     │
//  │ ==, !=, <, >, <=, >= │ ==, !=, <, >, <=, >=            │
//  │ &, |, ^, &^, <<, >> │ &, |, ^, <<, >>                 │
//  │ + (string concat)    │ + (string concat)                │
//  └──────────────────────┴──────────────────────────────────┘
// ============================================================

package main

import "fmt"

// =============================================================================
// 1. ARITHMETIC OPERATORS
// =============================================================================

func demoArithmetic() {
	fmt.Println("=== Arithmetic Operators ===")

	a, b := 10, 3

	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Printf("a + b = %d  (addition)\n", a+b)
	fmt.Printf("a - b = %d  (subtraction)\n", a-b)
	fmt.Printf("a * b = %d  (multiplication)\n", a*b)
	fmt.Printf("a / b = %d  (int division! truncates toward zero)\n", a/b)
	fmt.Printf("a %% b = %d  (modulo/remainder)\n", a%b)

	// Float division — at least one operand must be float
	fmt.Printf("float64(a) / float64(b) = %.2f (float division)\n",
		float64(a)/float64(b))

	// Unary: + and -
	fmt.Printf("-a = %d  (negation)\n", -a)
	fmt.Printf("+a = %d  (no effect)\n", +a)

	// ⚠️ Python differences:
	fmt.Println("\nPython vs Go:")
	fmt.Println("  Go:  10 / 3  = 3  (int division, truncates)")
	fmt.Println("  Py:  10 / 3  = 3.333... (true division)")
	fmt.Println("  Py:  10 // 3 = 3  (floor division)")
	fmt.Println("  Go has NO // operator")
	fmt.Println("  Go has NO ** operator (use math.Pow)")
}

// =============================================================================
// 2. INCREMENT / DECREMENT
// =============================================================================
// Python:  i += 1  (statement)
//          i++     (statement)
//
// Go:      i++     (statement ONLY — CANNOT be used in expressions)
//          i--     (statement ONLY)
//
// Critical difference:
//   Python: x = i++  ❌ not allowed (Python also doesn't have ++ in expressions)
//   Go:     x = i++  ❌ COMPILE ERROR — i++ is a statement, not an expression

func demoIncDec() {
	fmt.Println("\n=== Increment / Decrement ===")

	i := 5
	fmt.Printf("i = %d\n", i)

	i++
	fmt.Printf("i++ = %d  (post-increment)\n", i)

	i--
	fmt.Printf("i-- = %d  (post-decrement)\n", i)

	i += 3
	fmt.Printf("i += 3 = %d\n", i)

	i *= 2
	fmt.Printf("i *= 2 = %d\n", i)

	// The following would NOT compile:
	// x := i++       // ❌ i++ is a statement, not an expression
	// fmt.Println(i++)  // ❌ same reason
	// for j := 0; j < 5; i++ {  // ✅ this is fine (statement context)
}

// =============================================================================
// 3. COMPARISON (RELATIONAL) OPERATORS
// =============================================================================
// Same as Python: ==, !=, <, >, <=, >=
//
// But Go can also compare structs (field-by-field) and arrays —
// Python requires __eq__ for custom types.

func demoComparison() {
	fmt.Println("\n=== Comparison Operators ===")

	a, b := 10, 3
	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Printf("a == b = %t\n", a == b)
	fmt.Printf("a != b = %t\n", a != b)
	fmt.Printf("a <  b = %t\n", a < b)
	fmt.Printf("a >  b = %t\n", a > b)
	fmt.Printf("a <= b = %t\n", a <= b)
	fmt.Printf("a >= b = %t\n", a >= b)

	// Struct comparison (all comparable fields)
	type Point struct {
		X, Y int
	}
	p1 := Point{1, 2}
	p2 := Point{1, 2}
	p3 := Point{1, 3}
	fmt.Printf("\nPoint{1,2} == Point{1,2}: %t  (field-by-field)\n", p1 == p2)
	fmt.Printf("Point{1,2} == Point{1,3}: %t\n", p1 == p3)

	// Chaining ⚠️
	fmt.Println("\nGo vs Python chaining:")
	fmt.Println("  Python:  1 < a < 10  (chained comparison)")
	fmt.Println("  Go:      1 < a && a < 10  (no chaining, use &&)")
}

// =============================================================================
// 4. LOGICAL OPERATORS
// =============================================================================

func demoLogical() {
	fmt.Println("\n=== Logical Operators ===")

	t, f := true, false
	fmt.Printf("t = %t, f = %t\n", t, f)
	fmt.Printf("t && f = %t  (AND — Go uses &&, Python uses 'and')\n", t && f)
	fmt.Printf("t || f = %t  (OR  — Go uses ||, Python uses 'or')\n", t || f)
	fmt.Printf("!t     = %t  (NOT — Go uses !, Python uses 'not')\n", !t)

	// Short-circuit evaluation (same as Python)
	fmt.Println("\nShort-circuit: both Go and Python stop early")
	fmt.Println("  false && expensive()  → expensive() NEVER called")
	fmt.Println("  true  || expensive()  → expensive() NEVER called")

	// ⚠️ Go requires BOOL operands — no truthiness
	fmt.Println("\nGo vs Python — Truthiness:")
	fmt.Println("  Python: if 1: ...          // 1 is truthy")
	fmt.Println("  Go:     if 1 { ... }       // ❌ COMPILE ERROR!")
	fmt.Println("  Go:     if x == 1 { ... }  // ✅ must be explicit")
	fmt.Println("  Go has NO truthy/falsy — conditions MUST be bool")
}

// =============================================================================
// 5. BITWISE OPERATORS
// =============================================================================

func demoBitwise() {
	fmt.Println("\n=== Bitwise Operators ===")

	a, b := 12, 5 // 12=1100, 5=0101
	fmt.Printf("a = %d (%08b), b = %d (%08b)\n", a, a, b, b)
	fmt.Printf("a &  b = %3d (%08b)  AND\n", a&b, a&b)
	fmt.Printf("a |  b = %3d (%08b)  OR\n", a|b, a|b)
	fmt.Printf("a ^  b = %3d (%08b)  XOR (Go: ^, Python: ^)\n", a^b, a^b)
	fmt.Printf("a &^ b = %3d (%08b)  AND NOT (Go ONLY! No Python equivalent)\n", a&^b, a&^b)
	fmt.Printf("  ^a   = %3d (%08b)  Bitwise NOT (unary)\n", ^a, ^a)

	// Shifts
	c := 1
	fmt.Printf("\n1 << 0 = %d (%08b)  Shift left\n", c<<0, c<<0)
	fmt.Printf("1 << 1 = %d (%08b)\n", c<<1, c<<1)
	fmt.Printf("1 << 2 = %d (%08b)\n", c<<2, c<<2)
	fmt.Printf("1 << 3 = %d (%08b)\n", c<<3, c<<3)

	d := 16
	fmt.Printf("\n16 >> 0 = %d (%08b)  Shift right\n", d>>0, d>>0)
	fmt.Printf("16 >> 1 = %d (%08b)\n", d>>1, d>>1)
	fmt.Printf("16 >> 2 = %d (%08b)\n", d>>2, d>>2)

	fmt.Println("\nGo vs Python — Bitwise:")
	fmt.Println("  &^ (AND NOT) is unique to Go — clears bits from first operand")
	fmt.Println("  Example: 12 &^ 5 = 8 (1100 &^ 0101 = 1000)")
	fmt.Println("  Python equivalent of &^: a & ~b")
	fmt.Println("  Shift operators work the same in both")
}

// =============================================================================
// 6. ASSIGNMENT OPERATORS
// =============================================================================

func demoAssignment() {
	fmt.Println("\n=== Assignment Operators ===")

	x := 10 // basic assignment (:= declare + assign)

	x = 10       // simple assignment
	x += 5       // x = x + 5
	x -= 3       // x = x - 3
	x *= 2       // x = x * 2
	x /= 4       // x = x / 4
	x %= 3       // x = x % 3

	x = 12 // reset for bitwise
	x &= 5       // x = x & 5
	x |= 3       // x = x | 3
	x ^= 6       // x = x ^ 6
	x &^= 2      // x = x &^ 2 (AND NOT — unique to Go)
	x <<= 1      // x = x << 1
	x >>= 2      // x = x >> 2

	fmt.Println("  Same as Python except &^= which is Go-only")
	fmt.Println("  += also works for string concatenation:")
	greeting := "Hello"
	greeting += ", World!"
	fmt.Printf("    %q\n", greeting)
}

// =============================================================================
// 7. STRING CONCATENATION WITH +
// =============================================================================

func demoStringConcat() {
	fmt.Println("\n=== String Concatenation ===")

	// Same as Python
	s := "Go" + " is " + "awesome"
	fmt.Println("s := \"Go\" + \" is \" + \"awesome\"")
	fmt.Println("s =", s)

	// Mixed types need conversion
	// "Age: " + 30  // ❌ COMPILE ERROR
	age := "Age: " + fmt.Sprint(30)
	fmt.Println("With conversion:", age)

	// ⚠️ + in a loop is SLOW (see Strings-Builder/)
	fmt.Println("  + in a loop is O(n²) — use strings.Builder or strings.Join")
}

// =============================================================================
// 8. COMPLETE COMPARISON TABLE
// =============================================================================
//
// ┌────────────────────┬───────────────────┬───────────────────────┐
// │   Category         │      Go           │       Python          │
// ├────────────────────┼───────────────────┼───────────────────────┤
// │ Addition           │ +                 │ +                     │
// │ Subtraction        │ -                 │ -                     │
// │ Multiplication     │ *                 │ *                     │
// │ Division (int)     │ / (truncates)     │ / (→float), // (floor)│
// │ Modulo             │ %                 │ %                     │
// │ Exponent           │ math.Pow(a, b)    │ **                    │
// │ Floor division     │ ❌ No equivalent  │ //                    │
// ├────────────────────┼───────────────────┼───────────────────────┤
// │ Increment          │ i++ (statement)   │ i += 1                │
// │ Decrement          │ i-- (statement)   │ i -= 1                │
// ├────────────────────┼───────────────────┼───────────────────────┤
// │ AND                │ &&                │ and                   │
// │ OR                 │ ||                │ or                    │
// │ NOT                │ !                 │ not                   │
// │ Truthy values      │ ❌ Not allowed    │ ✅ Any value works    │
// ├────────────────────┼───────────────────┼───────────────────────┤
// │ Equal              │ ==                │ ==                    │
// │ Not equal          │ !=                │ !=                    │
// │ Chained compare    │ ❌ No             │ ✅ 1 < x < 10        │
// ├────────────────────┼───────────────────┼───────────────────────┤
// │ Bitwise AND        │ &                 │ &                     │
// │ Bitwise OR         │ |                 │ |                     │
// │ Bitwise XOR        │ ^                 │ ^                     │
// │ Bitwise NOT        │ ^                 │ ~                     │
// │ AND NOT (clear)    │ &^                │ ❌ (use a & ~b)       │
// │ Left shift         │ <<                │ <<                    │
// │ Right shift        │ >>                │ >>                    │
// ├────────────────────┼───────────────────┼───────────────────────┤
// │ String concat      │ +                 │ +                     │
// │ String repeat      │ strings.Repeat    │ * (str * int)         │
// └────────────────────┴───────────────────┴───────────────────────┘

// =============================================================================
// MAIN
// =============================================================================

func main() {
	demoArithmetic()
	demoIncDec()
	demoComparison()
	demoLogical()
	demoBitwise()
	demoAssignment()
	demoStringConcat()
}
