// ============================================================
//  TYPE CONVERSION IN GO — with Python Comparisons
// ============================================================
//  Go has NO implicit type conversion. You must convert EXPLICITLY
//  even between similar types like int and int32.
//
//  Python:  result = 1 + 2.0       # implicitly converts 1 → 1.0
//  Go:      result := 1 + int(2.0) # must convert 2.0 → 2 explicitly
//                               OR := float64(1) + 2.0
//
// ┌─────────────────────────┬─────────────────────────────────┐
// │          Go             │             Python              │
// ├─────────────────────────┼─────────────────────────────────┤
// │ float64(n)              │ float(n)                        │
// │ int(f)     [truncates]  │ int(f)      [truncates]         │
// │ uint(f)    [truncates]  │ int(f)      [truncates]         │
// │ strconv.Atoi(s)         │ int(s)      [raises]            │
// │ strconv.Itoa(n)         │ str(n)                          │
// │ strconv.ParseFloat(s,64)│ float(s)    [raises]            │
// │ []byte(s)               │ s.encode()                      │
// │ string(b)               │ b.decode()                      │
// └─────────────────────────┴─────────────────────────────────┘
// ============================================================

package main

import (
	"fmt"
	"math"
	"strconv"
)

// =============================================================================
// 1. NUMERIC TYPE CONVERSIONS
// =============================================================================
// Go has many numeric types: int, int8, int16, int32, int64,
// uint, uint8, uint16, uint32, uint64, float32, float64.
//
// ALL conversions between them must be EXPLICIT.
//
// Python has one int type and one float type — no size distinctions.

func demoNumericConversion() {
	fmt.Println("=== Numeric Type Conversion ===")

	// int → float64 (safe, no data loss)
	var age int = 30
	var f float64 = float64(age)
	fmt.Printf("int %d → float64 %.1f\n", age, f)

	// float64 → int (TRUNCATES toward zero — no rounding)
	var pi float64 = 3.14159
	var truncated int = int(pi)
	fmt.Printf("float64 %.5f → int %d (truncated!)\n", pi, truncated)

	// float64 → int with rounding
	rounded := int(math.Round(pi))
	fmt.Printf("float64 %.5f → int %d (rounded)\n", pi, rounded)

	// Between sized int types (int32 ↔ int64)
	var small int32 = 100
	var big int64 = int64(small)   // widening — safe
	var back int32 = int32(big)    // narrowing — safe only if value fits

	fmt.Printf("int32 %d → int64 %d → back to int32 %d\n", small, big, back)

	// 🚨 Narrowing can overflow silently!
	var large int64 = 1_000_000_000_000 // larger than int32 max
	var overflow int32 = int32(large)   // ⚠️ truncates to fit!
	fmt.Printf("int64 %d → int32 %d (OVERFLOW! data lost)\n", large, overflow)

	// uint ↔ int (signed ↔ unsigned)
	var pos uint = 42
	var signed int = int(pos)
	fmt.Printf("uint %d → int %d\n", pos, signed)

	// 🚨 uint → int can become negative if uint is large
	var bigUint uint = math.MaxInt64 + 1
	var negative int = int(bigUint)
	fmt.Printf("uint %d → int %d (wrong! became negative)\n", bigUint, negative)

	fmt.Println("\nPython equivalent:")
	fmt.Println("  There's no distinction between int8/int32/int64 in Python.")
	fmt.Println("  Go's explicit conversion prevents accidental overflow.")
}

// =============================================================================
// 2. STRING ↔ NUMBER CONVERSIONS (strconv package)
// =============================================================================
// Python:  int("123")  → 123    (raises ValueError on invalid)
//
// Go has MULTIPLE functions in the strconv package:
//   strconv.Atoi(s)          → (int, error)          ASCII to int
//   strconv.Itoa(n)          → string                int to ASCII
//   strconv.ParseInt(s, b, n)→ (int64, error)        parse with base + bit size
//   strconv.ParseFloat(s, b) → (float64, error)      parse float
//   strconv.FormatInt(n, b)  → string                format with base
//   strconv.FormatFloat(...) → string                formatted float

func demoStringToNumber() {
	fmt.Println("\n=== String ↔ Number Conversion ===")

	// === STRING → INT ===
	s := "123"
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("strconv.Atoi(%q) = %d (type: %T)\n", s, num, num)
	}

	// Invalid string → error returned (NOT a panic/exception)
	invalid := "abc"
	_, err = strconv.Atoi(invalid)
	if err != nil {
		fmt.Printf("strconv.Atoi(%q) → error: %v\n", invalid, err)
	}

	// === INT → STRING ===
	n := 255
	str := strconv.Itoa(n)
	fmt.Printf("strconv.Itoa(%d) = %q\n", n, str)

	// === STRING → FLOAT ===
	piStr := "3.14159"
	pi, err := strconv.ParseFloat(piStr, 64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("strconv.ParseFloat(%q, 64) = %f\n", piStr, pi)
	}

	// === FLOAT → STRING (formatted) ===
	formatted := strconv.FormatFloat(pi, 'f', 2, 64)
	fmt.Printf("strconv.FormatFloat(%f, 'f', 2, 64) = %s\n", pi, formatted)

	// === PARSE WITH BASE ===
	hexStr := "FF"
	val, _ := strconv.ParseInt(hexStr, 16, 64)
	fmt.Printf("strconv.ParseInt(%q, 16, 64) = %d\n", hexStr, val)

	binStr := "1010"
	val, _ = strconv.ParseInt(binStr, 2, 64)
	fmt.Printf("strconv.ParseInt(%q, 2, 64) = %d\n", binStr, val)

	// === FORMAT WITH BASE ===
	binary := strconv.FormatInt(42, 2)
	hex := strconv.FormatInt(255, 16)
	fmt.Printf("strconv.FormatInt(42, 2) = %s\n", binary)
	fmt.Printf("strconv.FormatInt(255, 16) = %s\n", hex)

	fmt.Println("\nPython vs Go:")
	fmt.Println("  Python:    int(s)  → raises ValueError or returns int")
	fmt.Println("  Go:        strconv.Atoi(s)  → returns (int, error)")
	fmt.Println("  Key: Go NEVER panics on bad input — it returns an error")
}

// =============================================================================
// 3. STRING ↔ BYTE SLICE CONVERSION
// =============================================================================
// Python:  "hello".encode()  → b'hello'
//          b'hello'.decode() → "hello"
//
// Go:      []byte("hello")   → []byte{104, 101, 108, 108, 111}
//          string([]byte{...}) → "hello"
//
// This is COMMON in Go for I/O operations.
// In Python, encode/decode are method calls.
// In Go, they're TYPE CONVERSIONS — fast and direct.

func demoStringBytes() {
	fmt.Println("\n=== String ↔ []byte Conversion ===")

	// string → []byte
	s := "Hello, Go!"
	bytes := []byte(s)
	fmt.Printf("string %q → []byte %v\n", s, bytes)

	// []byte → string
	back := string(bytes)
	fmt.Printf("[]byte %v → string %q\n", bytes, back)

	// Modify bytes independently (it's a COPY)
	bytes[0] = 'Y'
	fmt.Printf("After modifying bytes: bytes[0]='Y' → string still %q\n", back)

	// ⚠️ String → []byte converts each byte, not each rune
	// Multi-byte characters (like 世) become multiple bytes
	greeting := "Hello, 世界"
	gb := []byte(greeting)
	fmt.Printf("String %q has %d bytes: %v\n", greeting, len(gb), gb)
	fmt.Println("  (11 characters but 13 bytes — 世 and 界 are 3 bytes each)")

	fmt.Println("\nPython equivalent:")
	fmt.Println("  s.encode()     → bytes  (method call)")
	fmt.Println("  b.decode()     → string (method call)")
	fmt.Println("  Go: []byte(s)  → []byte (type conversion)")
	fmt.Println("  Go: string(b)  → string (type conversion)")
	fmt.Println("  Both are fast, but Go's syntax is more explicit.")
}

// =============================================================================
// 4. RUNE ↔ STRING CONVERSION
// =============================================================================
// A rune is a single Unicode code point (int32).
// Python doesn't have a separate rune type — a character is just a string.
//
// See strings-and-runes/ for the full treatment.

func demoRuneConversion() {
	fmt.Println("\n=== Rune ↔ String Conversion ===")

	// rune → string
	var r rune = 'A'
	s := string(r)
	fmt.Printf("rune %c (%U) → string %q\n", r, r, s)

	// int32 → string (same as rune — rune is an alias for int32)
	var code int32 = 9731 // ☃ snowman
	snowman := string(code)
	fmt.Printf("int32 %d → string %q\n", code, snowman)

	// string → []rune (get all Unicode code points)
	msg := "Hello"
	runes := []rune(msg)
	fmt.Printf("string %q → []rune %v\n", msg, runes)

	// ⚠️ string → []byte ≠ string → []rune for multi-byte strings
	world := "世界"
	byteCount := len([]byte(world))
	runeCount := len([]rune(world))
	fmt.Printf("%q: %d bytes, %d runes\n", world, byteCount, runeCount)
}

// =============================================================================
// 5. BOOL CONVERSION (strconv)
// =============================================================================
// Python:  bool("true")  → True    (truthy: any non-empty string)
//          bool("false") → True    (still truthy! non-empty)
//          bool("")      → False
//
// Go:      strconv.ParseBool("true")  → (true, nil)
//          strconv.ParseBool("false") → (false, nil)
//          strconv.ParseBool("yes")   → (false, error)
//
// Go is STRICT: only "true"/"1" and "false"/"0" are valid.
// Python treats ANY non-empty string as True.

func demoBoolConversion() {
	fmt.Println("\n=== Bool Conversion ===")

	// String → bool (strict)
	vals := []string{"true", "false", "1", "0", "yes", ""}
	for _, v := range vals {
		b, err := strconv.ParseBool(v)
		if err != nil {
			fmt.Printf("  strconv.ParseBool(%q) → error: %v\n", v, err)
		} else {
			fmt.Printf("  strconv.ParseBool(%q) → %t\n", v, b)
		}
	}

	fmt.Println("\nPython vs Go:")
	fmt.Println("  Python: bool('true') = True, bool('false') = True (truthy!)")
	fmt.Println("  Go:    strconv.ParseBool('true') = true, nil")
	fmt.Println("  Go:    strconv.ParseBool('false') = false, nil")
	fmt.Println("  Go:    strconv.ParseBool('yes') = error!")
	fmt.Println("  Go is strict — only 'true'/'1'/'false'/'0' are valid.")
}

// =============================================================================
// 6. TYPE ASSERTION — Converting Interface Values
// =============================================================================
// This is DIFFERENT from type conversion.
// Type ASSERTION extracts a concrete value from an interface.
// See interfaces.go for the full treatment.
//
// Python:  isinstance(x, int)  → True/False
//          int(x)              → convert (or raise)
//
// Go:      x.(int)             → (value, bool) — assertion, not conversion
//
// The difference:
//   Type CONVERSION:    int(f)    — changes float64 3.14 to int 3
//   Type ASSERTION:     x.(int)   — extracts an int from interface{}, no value change

func demoTypeAssertion() {
	fmt.Println("\n=== Type Assertion (not conversion) ===")

	var val any = 42 // any = interface{}

	// Type assertion extracts the concrete value
	extracted, ok := val.(int)
	if ok {
		fmt.Printf("Type assertion val.(int) = %d (type: %T)\n", extracted, extracted)
	}

	// Failed assertion (wrong type) — returns zero + false
	strVal, ok := val.(string)
	if !ok {
		fmt.Printf("val.(string) failed — ok=%v, value=%q (zero)\n", ok, strVal)
	}

	fmt.Println("\nType CONVERSION vs ASSERTION:")
	fmt.Println("  Conversion:  int(3.14)     → 3        (VALUE changes)")
	fmt.Println("  Assertion:   x.(int)       → 42       (value UNCHANGED, same 42)")
	fmt.Println("  Conversion:  changes representation")
	fmt.Println("  Assertion:   confirms type, keeps the value")
}

// =============================================================================
// 7. COMMON CONVERSION PITFALLS
// =============================================================================

func demoPitfalls() {
	fmt.Println("\n=== Common Pitfalls ===")

	// Pitfall 1: int division with float result
	a := 5
	b := 2
	// Both are ints → int division → 2 (not 2.5)
	result := a / b
	fmt.Printf("Pitfall 1: %d / %d = %d (int division!)\n", a, b, result)
	fmt.Println("  Fix: convert one operand first:")
	fmt.Printf("  float64(%d) / float64(%d) = %.1f\n", a, b, float64(a)/float64(b))

	// Pitfall 2: Converting string with non-ASCII
	s := "Hi 世界"
	bSlice := []byte(s)
	// Slicing bytes can split a multi-byte character
	broken := string(bSlice[:7]) // 7 bytes might cut a rune in half
	fmt.Printf("Pitfall 2: string(bSlice[:7]) = %q (may be invalid UTF-8!)\n", broken)

	// Pitfall 3: int → string is NOT strconv.Itoa
	// This compiles but does something WRONG:
	n := 65
	s3 := string(n) // ❌ not "65"! It's the Unicode character with code 65 = 'A'
	fmt.Printf("Pitfall 3: string(%d) = %q (not %q!)\n", n, s3, strconv.Itoa(n))
	fmt.Println("  string(int) converts the int as a UNICODE CODE POINT!")
	fmt.Println("  Use strconv.Itoa(n) for '65' → '65' ")
	fmt.Println("  Use string(n) for code point: 65 → 'A'")

	// Pitfall 4: rune → string is NOT the same as int → string
	// This is the SAME operation! rune is int32.
	fmt.Println("  string(rune(65)) =", string(rune(65))) // same as string(65)
}

// =============================================================================
// 8. COMPLETE COMPARISON TABLE
// =============================================================================
//
// ┌──────────────────────┬────────────────────────┬──────────────────────────┐
// │     Operation        │          Go            │         Python           │
// ├──────────────────────┼────────────────────────┼──────────────────────────┤
// │ int → float          │ float64(42)            │ float(42)                │
// ├──────────────────────┼────────────────────────┼──────────────────────────┤
// │ float → int          │ int(3.14)              │ int(3.14)                │
// │   (truncates)        │ → 3                    │ → 3                      │
// ├──────────────────────┼────────────────────────┼──────────────────────────┤
// │ float → int (round)  │ int(math.Round(3.14))  │ round(3.14)              │
// │                      │ → 3                    │ → 3                      │
// ├──────────────────────┼────────────────────────┼──────────────────────────┤
// │ string → int         │ strconv.Atoi("123")    │ int("123")               │
// │                      │ returns (int, error)   │ raises ValueError         │
// ├──────────────────────┼────────────────────────┼──────────────────────────┤
// │ int → string         │ strconv.Itoa(123)      │ str(123)                 │
// ├──────────────────────┼────────────────────────┼──────────────────────────┤
// │ string → float       │ strconv.ParseFloat     │ float("3.14")            │
// ├──────────────────────┼────────────────────────┼──────────────────────────┤
// │ float → string       │ strconv.FormatFloat    │ str(3.14), f"{x:.2f}"    │
// ├──────────────────────┼────────────────────────┼──────────────────────────┤
// │ int → other size     │ int32(n), int64(n)     │ n (Python handles it)    │
// ├──────────────────────┼────────────────────────┼──────────────────────────┤
// │ string → bytes       │ []byte("hi")           │ "hi".encode()            │
// ├──────────────────────┼────────────────────────┼──────────────────────────┤
// │ bytes → string       │ string([]byte{104,105})│ b"hi".decode()           │
// ├──────────────────────┼────────────────────────┼──────────────────────────┤
// │ rune → string        │ string('A')            │ 'A' (just a string)      │
// ├──────────────────────┼────────────────────────┼──────────────────────────┤
// │ string → bool        │ strconv.ParseBool      │ bool(s) (truthy/falsy)   │
// ├──────────────────────┼────────────────────────┼──────────────────────────┤
// │ interface → type     │ x.(int)  (assertion)   │ isinstance(x, int)       │
// ├──────────────────────┼────────────────────────┼──────────────────────────┤
// │ Implicit conversion  │ ❌ Never               │ ✅ int + float → float   │
// └──────────────────────┴────────────────────────┴──────────────────────────┘

// =============================================================================
// MAIN — Run All Demos
// =============================================================================

func main() {
	demoNumericConversion()
	demoStringToNumber()
	demoStringBytes()
	demoRuneConversion()
	demoBoolConversion()
	demoTypeAssertion()
	demoPitfalls()
}

// =============================================================================
// CHEAT SHEET
// =============================================================================
//
// // Numeric
// f := float64(i)          // int → float64 (safe)
// i := int(f)              // float64 → int (TRUNCATES!)
// i := int(math.Round(f))  // float64 → int (ROUNDS)
// i32 := int32(i64)        // int64 → int32 (may overflow!)
//
// // String ↔ Number
// n, err := strconv.Atoi(s)     // string → int
// s := strconv.Itoa(n)          // int → string
// f, err := strconv.ParseFloat(s, 64)  // string → float64
// s := strconv.FormatFloat(f, 'f', 2, 64)  // float64 → string (2 decimal)
// n, _ := strconv.ParseInt("FF", 16, 64)   // hex string → int64
// s := strconv.FormatInt(255, 16)          // int64 → hex string
//
// // String ↔ Bytes
// b := []byte(s)             // string → []byte
// s := string(b)             // []byte → string
//
// // Rune
// s := string(r)             // rune → string (Unicode code point)
// r := []rune(s)             // string → []rune
//
// // Bool
// b, err := strconv.ParseBool("true")  // string → bool (strict)
// s := strconv.FormatBool(true)        // bool → string
